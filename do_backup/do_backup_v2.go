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
	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/config"
	"github.com/sikalabs/tergum/telemetry"
	"github.com/sikalabs/tergum/version"
)

func DoBackupV2(
	configPath string,
	expandEnv bool,
	telemetryDisabled bool,
	extraName string,
	jsonLogs bool,
	debugLogs bool,
) {
	if !jsonLogs {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debugLogs {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// Seed random library
	rand.Seed(time.Now().UTC().UnixNano())

	log.Info().Str("version", version.Version).Msg("Tergum Backup")
	log.Info().Str("do_backup", "v2").Msg("Runs new (experimental) DoBackupV2 implementation")

	if extraName != "" {
		log.Info().Str("extra_name", extraName).Msg("extra name: " + extraName)
	}

	// Load config from file
	var config config.TergumConfig
	config.Load(configPath, expandEnv)

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
		var bo backup_output.BackupOutput
		var data io.ReadSeeker

		sleep(b, tel, i)

		// Backup source
		pb := DoBackupProcess{
			Telemetry: tel,
			BackupLog: &bl,
			Backup:    b,
		}

		pb.BackupStart()
		if b.RemoteExec == nil {
			// Standart local backup
			bo, pb.BackupError = b.Source.Backup()
			data = bo.Data
			pb.BackupStdError = bo.Stderr
		} else {
			// Remote backup using tergum server
			data, pb.BackupError = remoteBackup(b)
		}
		pb.BackupFinish()

		if pb.BackupError != nil {
			pb.BackupErr()
			continue
		}
		pb.BackupOK()

		// Process Backup's Middlewares
		pb.BackupMiddlewareStart()
		for _, m := range b.Middlewares {
			pb.Middleware = m
			pb.BackupMiddlewareStart()
			data, pb.BackupMiddlewareError = m.Process(data)
			pb.BackupMiddlewareFinish()
			if pb.BackupMiddlewareError != nil {
				pb.BackupMiddlewareErr()
				continue
			}
		}

		if pb.BackupMiddlewareError != nil {
			continue
		}

		for _, t := range b.Targets {
			pb.Target = t
			targetData := data
			targetData.Seek(0, 0)

			// Process Targets's Middlewares
			for _, m := range t.Middlewares {
				pb.Middleware = m
				pb.TargetMiddlewareStart()
				targetData, pb.TargetMiddlewareError = m.Process(targetData)
				pb.TargetMiddlewareFinish()

				if pb.TargetMiddlewareError != nil {
					pb.TargetMiddlewareErr()
					continue
				}
				pb.TargetMiddlewareOK()
			}
			if pb.TargetMiddlewareError != nil {
				continue
			}

			// Save backup to target
			pb.TargetStart()
			pb.TargetSize, pb.TargetError = t.Save(targetData)
			pb.TargetFinish()
			if pb.TargetError != nil {
				pb.SaveErr()
				continue
			}
			pb.SaveOK()
		}
	}

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

	output.BackupLogToOutput(bl)
	output.BackupErrorLogToOutput(bl)
}
