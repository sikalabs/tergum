package email

import (
	"fmt"
	"os"

	"github.com/sikalabs/tergum/backup_log"
	"github.com/sikalabs/tergum/backup_log/backup_log/output"
	"github.com/sikalabs/tergum/notification/backend"
)

type EmailRule struct {
	Emails []string `yaml:"Emails"`
	SendOK bool     `yaml:"SendOK"`
}

func (r EmailRule) Validate() error {
	if len(r.Emails) == 0 {
		return fmt.Errorf("must have at least one target email")
	}
	return nil
}

func (r EmailRule) SendNotification(
	bl backup_log.BackupLog,
	b backend.Backend,
) error {
	// Skip sending email on successfull backups
	// if SendOK is not set or set to false
	if bl.GlobalSuccess() && !r.SendOK {
		logSkipped()
		return nil
	}
	table := output.BackupLogToString(bl)
	errorTable := output.BackupErrorLogToString(bl)
	subject := "Backup Summary -- " + bl.GlobalSuccessString()
	if bl.ExtraName != "" {
		subject = "[" + bl.ExtraName + "] " + subject
	}
	body := `
<html>
<body>
<pre style="font: monospace">` + table + `</pre>
<pre style="font: monospace">` + errorTable + `</pre>
</body>
</html>`
	for _, email := range r.Emails {
		err := b.Email.SendMail(
			email,
			subject,
			body,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
	}
	logSend()
	return nil
}
