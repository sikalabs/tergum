package backup

import (
	"github.com/sikalabs/tergum/cmd/root"
	"github.com/sikalabs/tergum/do_backup"
	"github.com/spf13/cobra"
)

var CmdFlagConfig string
var CmdFlagExtraName string
var FlagDisableTelemetry bool
var FlagJsonLogs bool

var Cmd = &cobra.Command{
	Use:     "backup",
	Short:   "Do backup",
	Aliases: []string{"b"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		do_backup.DoBackup(
			CmdFlagConfig,
			FlagDisableTelemetry,
			CmdFlagExtraName,
			FlagJsonLogs,
		)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagConfig,
		"config",
		"c",
		"",
		"Path to config file",
	)
	Cmd.MarkFlagRequired("config")
	Cmd.Flags().StringVarP(
		&CmdFlagExtraName,
		"extra-name",
		"e",
		"",
		"Extra name for easy identification of specific run",
	)
	Cmd.Flags().BoolVar(
		&FlagDisableTelemetry,
		"disable-telemetry",
		false,
		"Disable telemetry",
	)
	Cmd.Flags().BoolVar(
		&FlagJsonLogs,
		"json-logs",
		false,
		"Log output to JSON",
	)
}
