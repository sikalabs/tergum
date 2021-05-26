package backend_email

import (
	"errors"
	"fmt"
	"os"

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
	rawMessage := []byte("To: " + to + "\r\n" +
		"From: " + config.Email + "\r\n" +
		"Subject: " + finalSubject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html\r\n" +
		"Content-Disposition: inline\r\n" +
		"\r\n" +
		finalBody + "\r\n")
	err := gosendmail.GoRawSendMail(
		config.SmtpHost,
		config.SmtpPort,
		config.Email,
		config.Password,
		to,
		rawMessage,
	)
	return err
}

func SendAlertEmail(
	backendConfig BackendEmail,
	alert AlertEmail,
	globalLog backup_log.BackupGlobalLog,
) error {
	table := backup_log.GlobalLogToString(globalLog)
	body := `
<html>
<body>
<pre style="font: monospace">` + table + `</pre>
</body>
</html>`
	for _, email := range alert.Emails {
		err := sendMail(
			backendConfig,
			email,
			"Backup Summary -- "+globalLog.SuccessString(),
			body,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
	}
	return nil
}
