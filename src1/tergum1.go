package src1

import (
	"log"
	"math/rand"
	"time"

	"github.com/sikalabs/tergum/src1/alerting"
	"github.com/sikalabs/tergum/src1/backup"
	tergum_config "github.com/sikalabs/tergum/src1/config"
)

func Tergum1(configPath string) {
	// Seed random library
	rand.Seed(time.Now().UTC().UnixNano())

	var config tergum_config.TergumConfig

	err := tergum_config.LoadConfig(&config, configPath)
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
