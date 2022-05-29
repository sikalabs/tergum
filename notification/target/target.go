package target

import (
	"fmt"

	"github.com/sikalabs/tergum/backup_log"
	"github.com/sikalabs/tergum/notification/backend"
	"github.com/sikalabs/tergum/notification/target/email"
	"github.com/sikalabs/tergum/notification/target/slack_webhook"
)

type Target struct {
	OnErrorOnly  bool                        `yaml:"OnErrorOnly"`
	Email        *email.EmailRule            `yaml:"Email"`
	SlackWebhook *slack_webhook.SlackWebhook `yaml:"SlackWebhook"`
}

func (r Target) Validate() error {
	if r.Email != nil {
		err := r.Email.Validate()
		if err != nil {
			return err
		}
		return nil
	}

	if r.SlackWebhook != nil {
		err := r.SlackWebhook.Validate()
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

	if r.SlackWebhook != nil {
		err := r.SlackWebhook.SendNotification(bl, b)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("missing AlertTarget configuratoin")
}
