package backend_email

import (
	"errors"

	gosendmail "github.com/ondrejsika/gosendmail/lib"
	"github.com/sikalabs/tergum/backup_log"
)

type BackendEmail struct {
	SmtpHost string
	SmtpPort string
	Email    string
	Password string
}

type AlertEmail struct {
	Emails []string
}

func (config *BackendEmail) Validate() error {
	if config.SmtpHost == "" {
		return errors.New("email backend requires smtpHost")
	}
	if config.SmtpPort == "" {
		return errors.New("email backend requires smtpPort")
	}
	return nil
}

func (config *AlertEmail) Validate() error {
	if len(config.Emails) == 0 {
		return errors.New("email alert requires emails (list)")
	}
	return nil
}

func sendMail(config BackendEmail, to string, subject string, body string) error {
	finalSubject := "[tergum] " + subject
	finalBody := body + "\n\n--\ntergum"
	err := gosendmail.GoSendMail(
		config.SmtpHost,
		config.SmtpPort,
		config.Email,
		config.Password,
		to,
		finalSubject,
		finalBody,
	)
	return err
}

func SendAlertEmail(
	backendConfig BackendEmail,
	alert AlertEmail,
	globalLog backup_log.BackupGlobalLog,
) error {
	table := backup_log.GlobalLogToString(globalLog)
	for _, email := range alert.Emails {
		sendMail(
			backendConfig,
			email,
			"Backup Summary -- "+globalLog.SuccessString(),
			table,
		)
	}
	return nil
}
