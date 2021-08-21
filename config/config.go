package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/sikalabs/tergum/backup"
	"github.com/sikalabs/tergum/notification"
)

const MIN_CONFIG_VERSION = 3
const MAX_CONFIG_VERSION = 3

type TergumConfigMeta struct {
	SchemaVersion int
}

type TergumConfig struct {
	Meta         TergumConfigMeta
	Backups      []backup.Backup
	Notification *notification.Notification
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
	// Validate Schema Version
	err := c.ValidateSchemaVersion()
	if err != nil {
		return err
	}

	// Validate all Backups
	for _, b := range c.Backups {
		err := b.Validate()
		if err != nil {
			return err
		}
	}

	if c.Notification != nil {
		err := c.Notification.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *TergumConfig) ValidateSchemaVersion() error {
	if c.Meta.SchemaVersion < MIN_CONFIG_VERSION {
		return fmt.Errorf(
			"your config schemaVersion %d is lower than minimum schemaVersion %d",
			c.Meta.SchemaVersion,
			MIN_CONFIG_VERSION,
		)
	}
	if c.Meta.SchemaVersion > MAX_CONFIG_VERSION {
		return fmt.Errorf(
			"your config schemaVersion %d is greather than minimum schemaVersion %d",
			c.Meta.SchemaVersion,
			MAX_CONFIG_VERSION,
		)
	}
	return nil
}
