package get_chat_id

import (
	"fmt"
	"log"

	parentcmd "github.com/sikalabs/tergum/cmd/utils/telegram"
	"github.com/sikalabs/tergum/utils/telegram_utils"
	"github.com/spf13/cobra"
)

var FlagBotToken string

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
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagBotToken,
		"bot-token",
		"t",
		FlagBotToken,
		"Telegram Bot token",
	)
	Cmd.MarkFlagRequired("bot-token")
}
