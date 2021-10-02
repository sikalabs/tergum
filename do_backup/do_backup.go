package do_backup

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sikalabs/tergum/backup_log"
	"github.com/sikalabs/tergum/backup_log/backup_log/output"
	"github.com/sikalabs/tergum/config"
	"github.com/sikalabs/tergum/version"
)

func DoBackup(configPath, extraName string) {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	})

	// Seed random library
	rand.Seed(time.Now().UTC().UnixNano())

	log.Info().Str("version", version.Version).Msg("Tergum Backup")

	if extraName != "" {
		log.Info().Str("extra_name", extraName).Msg("extra name: " + extraName)
	}

	// Load config from file
	var config config.TergumConfig
	config.Load(configPath)

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

	for _, b := range config.Backups {
		// Backup source
		logBackupStart(b)
		data, err := b.Source.Backup()
		if err != nil {
			bl.SaveEvent(b.ID, "---", err)
			logBackupFailed(b, err)
			continue
		}
		logBackupDone(b)

		// Process Backup's Middlewares
		var errBackupMiddleware error = nil
		for _, m := range b.Middlewares {
			logBackupMiddlewareStart(b, m)
			data, errBackupMiddleware = m.Process(data)
			if errBackupMiddleware != nil {
				bl.SaveEvent(b.ID, "---", errBackupMiddleware)
				logBackupMiddlewareFailed(b, m, err)
				continue
			}
			logBackupMiddlewareDone(b, m)
		}

		if errBackupMiddleware != nil {
			continue
		}

		for _, t := range b.Targets {
			targetData := data
			targetData.Seek(0, 0)

			// Process Targets's Middlewares
			var errTargetMiddleware error = nil
			for _, m := range t.Middlewares {
				logTargetMiddlewareStart(b, t, m)
				targetData, errTargetMiddleware = m.Process(targetData)
				if errTargetMiddleware != nil {
					bl.SaveEvent(b.ID, t.ID, errTargetMiddleware)
					logTargetMiddlewareFailed(b, t, m, errTargetMiddleware)
					continue
				}
				logTargetMiddlewareDone(b, t, m)
			}
			if errTargetMiddleware != nil {
				continue
			}

			// Save backup to target
			logTargetStart(b, t)
			err = t.Save(targetData)
			if err != nil {
				bl.SaveEvent(b.ID, t.ID, err)
				logTargetFailed(b, t, err)
				continue
			}
			bl.SaveEvent(b.ID, t.ID, err)
			logTargetDone(b, t)
		}
	}

	output.BackupLogToOutput(bl)

	// Send Notifications
	if config.Notification != nil {
		config.Notification.SendNotification(bl)
	}
}
