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
	table := output.BackupLogToString(bl)
	body := `
<html>
<body>
<pre style="font: monospace">` + table + `</pre>
</body>
</html>`
	for _, email := range r.Emails {
		err := b.Email.SendMail(
			email,
			"Backup Summary -- "+bl.GlobalSuccessString(),
			body,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
	}
	return nil
}
