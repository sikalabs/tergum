package pause

import (
	"time"

	parentcmd "github.com/sikalabs/tergum/cmd/utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "pause",
	Short: "pause",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		for {
			time.Sleep(time.Hour)
		}
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
}
