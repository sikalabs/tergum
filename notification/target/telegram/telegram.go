package telegram

import (
	"fmt"

	"github.com/sikalabs/tergum/backup_log"
	"github.com/sikalabs/tergum/backup_log/backup_log/telegram_output"
	"github.com/sikalabs/tergum/notification/backend"
	"github.com/sikalabs/tergum/utils/telegram_utils"
)

type Telegram struct {
	BotToken string `yaml:"BotToken" json:"BotToken,omitempty"`
	ChatID   int64  `yaml:"ChatID" json:"ChatID,omitempty"`
	SendOK   bool   `yaml:"SendOK" json:"SendOK,omitempty"`
}

func (t Telegram) Validate() error {
	if t.BotToken == "" {
		return fmt.Errorf("telegram: BotToken is required")
	}
	if t.ChatID == 0 {
		return fmt.Errorf("telegram: ChatID is required")
	}
	return nil
}

func (t Telegram) SendNotification(
	bl backup_log.BackupLog,
	b backend.Backend,
) error {
	// Skip sending email on successfull backups
	// if SendOK is not set or set to false
	if bl.GlobalSuccess() && !t.SendOK {
		logSkipped()
		return nil
	}
	table := telegram_output.BackupLogTelegram(bl)
	errorTable := telegram_output.BackupErrorLogTelegram(bl)

	extraName := ""
	if bl.ExtraName != "" {
		extraName = "=== " + bl.ExtraName + " ===\n\n"
	}

	text := bl.GlobalSuccessEmoji() +
		bl.GlobalSuccessEmoji() +
		bl.GlobalSuccessEmoji() +
		"\n\n" +
		extraName +
		table + errorTable + "\n"

	err := telegram_utils.TelegramSendMessage(
		t.BotToken,
		t.ChatID,
		text,
	)
	if err != nil {
		return fmt.Errorf("telegram: %w", err)
	}

	logSend()
	return nil
}
