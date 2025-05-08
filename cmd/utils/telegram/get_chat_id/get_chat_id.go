package get_chat_id

import (
	"fmt"
	"log"
	"os"

	parentcmd "github.com/sikalabs/tergum/cmd/utils/telegram"
	"github.com/sikalabs/tergum/utils/telegram_utils"
	"github.com/spf13/cobra"
)

var FlagBotToken string
var FlagSendToChat bool

var Cmd = &cobra.Command{
	Use:   "get-chat-id",
	Short: "Get Telegram chat ID",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		chatID, err := telegram_utils.TelegramGetLastChatID(FlagBotToken)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(chatID)
		if FlagSendToChat {
			telegram_utils.TelegramSendMessage(FlagBotToken, chatID, fmt.Sprintf("%d", chatID))
		}
	},
}

func init() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagBotToken,
		"bot-token",
		"t",
		botToken,
		"Telegram Bot token, can be set via TELEGRAM_BOT_TOKEN env var",
	)
	if botToken == "" {
		Cmd.MarkFlagRequired("bot-token")
	}
	Cmd.Flags().BoolVarP(
		&FlagSendToChat,
		"send-to-chat",
		"s",
		false,
		"Send message to chat",
	)
}
