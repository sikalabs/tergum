package slack_webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sikalabs/tergum/backup_log"
	"github.com/sikalabs/tergum/backup_log/backup_log/output"
	"github.com/sikalabs/tergum/notification/backend"
)

type SlackWebhook struct {
	URLs   []string `yaml:"URLs"`
	SendOK bool     `yaml:"SendOK"`
}

func (t SlackWebhook) Validate() error {
	if len(t.URLs) == 0 {
		return fmt.Errorf("must have at least one webhook URL")
	}
	return nil
}

func (t SlackWebhook) SendNotification(
	bl backup_log.BackupLog,
	b backend.Backend,
) error {
	// Skip sending email on successfull backups
	// if SendOK is not set or set to false
	if bl.GlobalSuccess() && !t.SendOK {
		logSkipped()
		return nil
	}
	table := output.BackupLogToString(bl)
	errorTable := output.BackupErrorLogToString(bl)

	subject := "Backup Summary -- " + bl.GlobalSuccessString()
	if bl.ExtraName != "" {
		subject = "[" + bl.ExtraName + "] " + subject
	}

	text := subject + "\n\n" +
		"```\n" + table + "\n" + errorTable + "\n```"

	for _, url := range t.URLs {
		values := map[string]string{"text": text}
		json_data, err := json.Marshal(values)
		_ = err
		resp, err := http.Post(url, "application/json",
			bytes.NewBuffer(json_data))
		_ = err
		_ = resp
	}

	logSend()
	return nil
}
