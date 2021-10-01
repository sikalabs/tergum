package do_backup

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/sikalabs/tergum/backup_log"
	"github.com/sikalabs/tergum/backup_log/backup_log/output"
	"github.com/sikalabs/tergum/config"
)

func DoBackup(configPath, extraName string) {
	// Seed random library
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("tergum v2")

	if extraName != "" {
		fmt.Println("extra name:", extraName)
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
		data, err := b.Source.Backup()
		if err != nil {
			bl.SaveEvent(b.ID, "---", err)
			continue
		}

		// Process Backup's Middlewares
		var errBackupMiddleware error = nil
		for _, m := range b.Middlewares {
			data, errBackupMiddleware = m.Process(data)
			if errBackupMiddleware != nil {
				bl.SaveEvent(b.ID, "---", errBackupMiddleware)
				continue
			}
		}

		if errBackupMiddleware != nil {
			continue
		}

		for _, t := range b.Targets {
			targetData := data

			// Process Targets's Middlewares
			var errTargetMiddleware error = nil
			for _, m := range t.Middlewares {
				targetData, errTargetMiddleware = m.Process(targetData)
				if errTargetMiddleware != nil {
					bl.SaveEvent(b.ID, t.ID, errTargetMiddleware)
					continue
				}
			}
			if errTargetMiddleware != nil {
				continue
			}

			// Save backup to target
			err = t.Save(targetData)
			if err != nil {
				bl.SaveEvent(b.ID, t.ID, err)
				continue
			}
			bl.SaveEvent(b.ID, t.ID, err)
		}
	}

	output.BackupLogToOutput(bl)

	// Send Notifications
	if config.Notification != nil {
		config.Notification.SendNotification(bl)
	}
}
