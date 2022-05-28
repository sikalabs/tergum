package do_backup

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sikalabs/tergum/backup_log"
	"github.com/sikalabs/tergum/backup_log/backup_log/output"
	"github.com/sikalabs/tergum/config"
	"github.com/sikalabs/tergum/telemetry"
	"github.com/sikalabs/tergum/version"
)

func DoBackup(
	configPath string,
	telemetryDisabled bool,
	extraName string,
	jsonLogs bool,
) {
	if !jsonLogs {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	}

	// Seed random library
	rand.Seed(time.Now().UTC().UnixNano())

	log.Info().Str("version", version.Version).Msg("Tergum Backup")

	if extraName != "" {
		log.Info().Str("extra_name", extraName).Msg("extra name: " + extraName)
	}

	// Load config from file
	var config config.TergumConfig
	config.Load(configPath)

	var cloudEmail string
	if config.Cloud != nil {
		cloudEmail = config.Cloud.Email
	}

	// Init Telemetry
	tel := telemetry.NewTelemetry(
		config.Telemetry,
		telemetryDisabled,
		extraName,
		cloudEmail,
	)

	tel.SendEventInit()

	// Create Backup Log
	bl := backup_log.BackupLog{
		ExtraName: extraName,
	}

	// Validate config
	err := config.Validate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, b := range config.Backups {
		var data io.ReadSeeker
		var stdErr string

		if b.SleepBefore != 0 && i != 0 {
			logSleepStart(tel, b)
			time.Sleep(time.Duration(b.SleepBefore) * time.Second)
			logSleepDone(tel, b)
		}

		// Backup source
		logBackupStart(tel, b)
		backupStart := time.Now()
		if b.RemoteExec == nil {
			// Standart local backup
			data, stdErr, err = b.Source.Backup()
		} else {
			// Remote backup using tergum server
			data, err = remoteBackup(b)
		}
		backupDuration := time.Since(backupStart)

		if err != nil {
			bl.SaveEvent(
				b.Source.Name(), b.ID, "---", "---",
				int(backupDuration.Seconds()), 0, 0, 0, 0,
				err, stdErr)
			logBackupFailed(tel, b, int(backupDuration.Seconds()), err)
			continue
		}
		logBackupDone(tel, b, int(backupDuration.Seconds()))

		// Process Backup's Middlewares
		var errBackupMiddleware error = nil
		backupMiddlewareStart := time.Now()
		var backupMiddlewareDuration time.Duration
		for _, m := range b.Middlewares {
			logBackupMiddlewareStart(tel, b, m)
			data, errBackupMiddleware = m.Process(data)
			if errBackupMiddleware != nil {
				backupMiddlewareDuration = time.Since(backupMiddlewareStart)
				bl.SaveEvent(b.Source.Name(), b.ID, "---", "---",
					int(backupDuration.Seconds()),
					int(backupMiddlewareDuration.Seconds()),
					0, 0, 0,
					errBackupMiddleware, "")
				logBackupMiddlewareFailed(tel, b, m, int(backupMiddlewareDuration.Seconds()), err)
				continue
			}
			backupMiddlewareDuration = time.Since(backupMiddlewareStart)
			logBackupMiddlewareDone(tel, b, m, int(backupMiddlewareDuration.Seconds()))
		}

		if errBackupMiddleware != nil {
			continue
		}

		for _, t := range b.Targets {
			targetData := data
			targetData.Seek(0, 0)

			// Process Targets's Middlewares
			var errTargetMiddleware error = nil
			targetMiddlewareStart := time.Now()
			var targetMiddlewareDuration time.Duration
			for _, m := range t.Middlewares {
				logTargetMiddlewareStart(tel, b, t, m)
				targetData, errTargetMiddleware = m.Process(targetData)
				if errTargetMiddleware != nil {
					targetMiddlewareDuration = time.Since(targetMiddlewareStart)
					bl.SaveEvent(b.Source.Name(), b.ID, t.Name(), t.ID,
						int(backupDuration.Seconds()),
						int(backupMiddlewareDuration.Seconds()),
						0,
						int(targetMiddlewareDuration.Seconds()),
						0,
						errTargetMiddleware, "")
					logTargetMiddlewareFailed(tel, b, t, m,
						int(targetMiddlewareDuration.Seconds()),
						errTargetMiddleware)
					continue
				}
				targetMiddlewareDuration = time.Since(targetMiddlewareStart)
				logTargetMiddlewareDone(tel, b, t, m, int(targetMiddlewareDuration.Seconds()))
			}
			if errTargetMiddleware != nil {
				continue
			}

			// Save backup to target
			logTargetStart(tel, b, t)
			targetStart := time.Now()
			size, err := t.Save(targetData)
			targetDuration := time.Since(targetStart)
			if err != nil {
				bl.SaveEvent(
					b.Source.Name(), b.ID, t.Name(), t.ID,
					int(backupDuration.Seconds()),
					int(backupMiddlewareDuration.Seconds()),
					int(targetDuration.Seconds()),
					int(targetMiddlewareDuration.Seconds()),
					size,
					err, "")
				logTargetFailed(tel, b, t, int(targetDuration.Seconds()), err)
				continue
			}
			bl.SaveEvent(
				b.Source.Name(), b.ID, t.Name(), t.ID,
				int(backupDuration.Seconds()),
				int(backupMiddlewareDuration.Seconds()),
				int(targetDuration.Seconds()),
				int(targetMiddlewareDuration.Seconds()),
				size,
				err, "")
			logTargetDone(tel, b, t, int(targetDuration.Seconds()))
		}
	}

	output.BackupLogToOutput(bl)
	output.BackupErrorLogToOutput(bl)

	// Log BackupLog to STDOUT in JSON
	out, _ := json.Marshal(bl)
	log.Info().
		Str("log-id", "backup-log-dump").
		Msg(string(out))

	// Send BackupLog to telemetry server
	if config.Telemetry != nil &&
		config.Telemetry.CollectBackupLog {
		tel.SendEventBackupLog(bl)
	}

	// Send Notifications
	if config.Notification != nil {
		config.Notification.SendNotification(bl)
	}
}
