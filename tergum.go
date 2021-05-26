package main

import (
	"errors"
	"flag"
	"log"

	"github.com/sikalabs/tergum/alerting"
	"github.com/sikalabs/tergum/backup"
	tergum_config "github.com/sikalabs/tergum/config"
)

func main() {
	// Backup parameters from config file
	path := flag.String("config", "", "tergum config file (json)")

	flag.Parse()

	if *path == "" {
		log.Fatal(errors.New("tergum require config file (-config)"))
	}

	var config tergum_config.TergumConfig

	err := tergum_config.LoadConfig(&config, *path)
	if err != nil {
		log.Fatal(err)
	}

	err = tergum_config.ValidateConfigVersion(&config)
	if err != nil {
		log.Fatal(err)
	}

	globalLog, err := backup.BackupAndSaveAll(config.Backups)
	if err != nil {
		log.Fatal(err)
	}

	alerting.SendAlerts(
		config.Alerting,
		globalLog,
	)

}
