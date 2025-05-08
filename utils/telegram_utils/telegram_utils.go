package telegram_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
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

func TelegramSendMessageWithFile(
	botToken string, chatID int64, message string,
	fileName string, file io.Reader,
) error {
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Add chat_id field
	err := writer.WriteField("chat_id", fmt.Sprintf("%d", chatID))
	if err != nil {
		return fmt.Errorf("failed to write chat_id field: %v", err)
	}

	// Add caption/message
	err = writer.WriteField("caption", message)
	if err != nil {
		return fmt.Errorf("failed to write caption field: %v", err)
	}

	// Add the file
	part, err := writer.CreateFormFile("document", fileName)
	if err != nil {
		log.Fatalf("Failed to create form file: %v", err)
	}

	// Copy from io.Reader into the form part
	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("failed to copy file content: %v", err)
	}

	writer.Close()

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument", botToken)

	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()
	return nil
}

func TelegramGetLastChatID(botToken string) (int64, error) {
	type Chat struct {
		ID int64 `json:"id"`
	}

	type Message struct {
		Chat Chat `json:"chat"`
	}

	type Update struct {
		Message Message `json:"message"`
	}

	type Response struct {
		Ok     bool     `json:"ok"`
		Result []Update `json:"result"`
	}

	resp, err := http.Get("https://api.telegram.org/bot" + botToken + "/getUpdates")
	if err != nil {
		return 0, fmt.Errorf("getting updates failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("reading response body failed: %v", err)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return 0, fmt.Errorf("unmarshalling response failed: %v", err)
	}

	if len(response.Result) == 0 {
		return 0, fmt.Errorf("no updates found")
	}

	lastChatID := response.Result[len(response.Result)-1].Message.Chat.ID
	return lastChatID, nil
}
