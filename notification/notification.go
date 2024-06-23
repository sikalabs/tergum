package notification

import (
	"github.com/sikalabs/tergum/backup_log"
	"github.com/sikalabs/tergum/notification/backend"
	"github.com/sikalabs/tergum/notification/target"
)

type Notification struct {
	Backend backend.Backend `yaml:"Backend" json:"Backend,omitempty"`
	Targets []target.Target `yaml:"Targets" json:"Targets,omitempty"`
}

func (a Notification) Validate() error {
	// Validate Backend
	err := a.Backend.Validate()
	if err != nil {
		return err
	}

	// Validate all Targets
	for _, t := range a.Targets {
		err = t.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (a Notification) SendNotification(bl backup_log.BackupLog) error {
	// Send all Targets
	for _, t := range a.Targets {
		err := t.SendNotification(bl, a.Backend)
		if err != nil {
			return err
		}
	}

	return nil
}
