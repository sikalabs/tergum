package target

import (
	"fmt"

	"github.com/sikalabs/tergum/backup_log"
	"github.com/sikalabs/tergum/notification/backend"
	"github.com/sikalabs/tergum/notification/target/email"
)

type Target struct {
	OnErrorOnly bool             `yaml:"OnErrorOnly"`
	Email       *email.EmailRule `yaml:"Email"`
}

func (r Target) Validate() error {
	if r.Email != nil {
		err := r.Email.Validate()
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("missing AlertTarget configuratoin")
}

func (r Target) SendNotification(
	bl backup_log.BackupLog,
	b backend.Backend,
) error {
	if r.Email != nil {
		err := r.Email.SendNotification(bl, b)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("missing AlertTarget configuratoin")
}
