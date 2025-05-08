package telegram

import (
	"fmt"
	"io"

	"github.com/sikalabs/tergum/utils/telegram_utils"
)

type TelegramTarget struct {
	BotToken string `yaml:"BotToken" json:"BotToken,omitempty"`
	ChatID   int64  `yaml:"ChatID" json:"ChatID,omitempty"`
	FileName string `yaml:"FileName" json:"Suffix,omitempty"`
}

func (t TelegramTarget) Validate() error {
	if t.BotToken == "" {
		return fmt.Errorf("TelegramTarget requires BotToken")
	}
	if t.ChatID == 0 {
		return fmt.Errorf("TelegramTarget requires ChatID")
	}
	if t.FileName == "" {
		return fmt.Errorf("TelegramTarget requires FileName")
	}
	return nil
}

func (t TelegramTarget) Save(data io.ReadSeeker) error {
	return telegram_utils.TelegramSendMessageWithFile(t.BotToken, t.ChatID, "", t.FileName, data)
}
