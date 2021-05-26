package alerting

import (
	"fmt"
	"os"

	"github.com/sikalabs/tergum/alerting/backend_email"
	"github.com/sikalabs/tergum/backup_log"
)

type Backends struct {
	Email backend_email.BackendEmail
}

type Alert struct {
	Backend string
	Email   backend_email.AlertEmail
}

type Alerting struct {
	Backends Backends
	Alerts   []Alert
}

func SendAlerts(config Alerting, globalLog backup_log.BackupGlobalLog) error {
	for i := 0; i < len(config.Alerts); i++ {
		alert := config.Alerts[i]
		switch alert.Backend {
		case "email":
			err := config.Backends.Email.Validate()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			}
			err = alert.Email.Validate()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			}
			return backend_email.SendAlertEmail(
				config.Backends.Email,
				alert.Email,
				globalLog,
			)
		}
	}
	return nil
}
