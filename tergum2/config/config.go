package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/sikalabs/tergum/tergum2/backup"
)

type TergumConfigMeta struct {
	SchemaVersion int
}

type TergumConfig struct {
	Meta    TergumConfigMeta
	Backups []backup.Backup
}

func (c *TergumConfig) Load(path string) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (c TergumConfig) Validate() error {
	// Validate all Backups
	for _, b := range c.Backups {
		err := b.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
