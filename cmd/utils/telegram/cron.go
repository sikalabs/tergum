package telegram

import (
	parentcmd "github.com/sikalabs/tergum/cmd/utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "telegram",
	Short: "Telegram utils",
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
}
