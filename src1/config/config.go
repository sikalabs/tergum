package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/sikalabs/tergum/src1/alerting"
	"github.com/sikalabs/tergum/src1/backup"
)

const MIN_CONFIG_VERSION = 2
const MAX_CONFIG_VERSION = 2

type TergumConfigMeta struct {
	SchemaVersion int
}

type TergumConfig struct {
	Meta     TergumConfigMeta
	Backups  []backup.Backups
	Alerting alerting.Alerting
}

func LoadConfig(config *TergumConfig, path string) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func ValidateConfigVersion(config *TergumConfig) error {
	if config.Meta.SchemaVersion < MIN_CONFIG_VERSION {
		return fmt.Errorf(
			"your config schemaVersion %d is lower than minimum schemaVersiion %d",
			config.Meta.SchemaVersion,
			MIN_CONFIG_VERSION,
		)
	}
	if config.Meta.SchemaVersion > MAX_CONFIG_VERSION {
		return fmt.Errorf(
			"your config schemaVersion %d is greather than minimum schemaVersiion %d",
			config.Meta.SchemaVersion,
			MAX_CONFIG_VERSION,
		)
	}
	return nil
}
