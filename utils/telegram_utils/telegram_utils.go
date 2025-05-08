package telegram_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func TelegramSendMessage(botToken string, chatID int64, message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	body, err := json.Marshal(map[string]string{
		"chat_id": fmt.Sprintf("%d", chatID),
		"text":    message,
	})
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
