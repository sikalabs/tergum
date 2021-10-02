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
		log.Info().
			Str("phase", "backup-start").
			Str("id", b.ID).
			Msg("Start backing up " + b.ID + " ...")
		data, err := b.Source.Backup()
		if err != nil {
			bl.SaveEvent(b.ID, "---", err)
			log.Warn().
				Str("phase", "backup-finish-err").
				Str("id", b.ID).
				Msg("Backup failed " + b.ID + ": " + err.Error())
			continue
		}
		log.Info().
			Str("phase", "backup-finish-ok").
			Str("id", b.ID).
			Msg("Finish backing up " + b.ID)

		// Process Backup's Middlewares
		var errBackupMiddleware error = nil
		for _, m := range b.Middlewares {
			log.Info().
				Str("phase", "backup-middleware-start").
				Str("id", b.ID).
				Msg("Start backup middleware")
			data, errBackupMiddleware = m.Process(data)
			if errBackupMiddleware != nil {
				bl.SaveEvent(b.ID, "---", errBackupMiddleware)
				log.Warn().
					Str("phase", "backup-middleware-err").
					Str("id", b.ID).
					Msg("Backup middleware failed: " + errBackupMiddleware.Error())
				continue
			}
			log.Info().
				Str("phase", "backup-middleware-ok").
				Str("id", b.ID).
				Msg("Finish backup middleware " + b.ID)
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
				log.Info().
					Str("phase", "target-middleware-start").
					Str("id", t.ID).
					Msg("Start target middleware")
				targetData, errTargetMiddleware = m.Process(targetData)
				if errTargetMiddleware != nil {
					bl.SaveEvent(b.ID, t.ID, errTargetMiddleware)
					log.Warn().
						Str("phase", "target-middleware-err").
						Str("id", t.ID).
						Msg("Target middleware failed: " + errTargetMiddleware.Error())
					continue
				}
				log.Info().
					Str("phase", "target-middleware-ok").
					Str("id", t.ID).
					Msg("Finish target middleware " + t.ID)
			}
			if errTargetMiddleware != nil {
				continue
			}

			// Save backup to target
			log.Info().
				Str("phase", "save-start").
				Str("id", t.ID).
				Msg("Start save " + t.ID)
			err = t.Save(targetData)
			if err != nil {
				bl.SaveEvent(b.ID, t.ID, err)
				log.Warn().
					Str("phase", "save-err").
					Str("id", t.ID).
					Msg("Save " + t.ID + " failed: " + err.Error())
				continue
			}
			bl.SaveEvent(b.ID, t.ID, err)
			log.Info().
				Str("phase", "save-ok").
				Str("id", t.ID).
				Msg("Finish save " + t.ID)
		}
	}

	output.BackupLogToOutput(bl)

	// Send Notifications
	if config.Notification != nil {
		config.Notification.SendNotification(bl)
	}
}
