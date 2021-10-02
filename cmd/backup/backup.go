package backup

import (
	"github.com/sikalabs/tergum/cmd/root"
	"github.com/sikalabs/tergum/do_backup"
	"github.com/sikalabs/tergum/src1"
	"github.com/sikalabs/tergum/telemetry"
	"github.com/spf13/cobra"
)

var CmdFlagConfig string
var CmdFlagExtraName string
var FlagTelemetryOrigin string
var CmdFlagImplementation1 bool
var FlagDisableTelemetry bool

var Cmd = &cobra.Command{
	Use:     "backup",
	Short:   "Do backup",
	Aliases: []string{"b"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if CmdFlagImplementation1 {
			src1.Tergum1(CmdFlagConfig)
			return
		}
		do_backup.DoBackup(
			CmdFlagConfig,
			!FlagDisableTelemetry,
			FlagTelemetryOrigin,
			CmdFlagExtraName,
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
	Cmd.Flags().BoolVar(
		&CmdFlagImplementation1,
		"implementation1",
		false,
		"Switch to implementation1 (src1)",
	)
	Cmd.Flags().StringVarP(
		&CmdFlagExtraName,
		"extra-name",
		"e",
		"",
		"Extra name for easy identification of specific run",
	)
	Cmd.Flags().StringVar(
		&FlagTelemetryOrigin,
		"telemetry-origin",
		telemetry.DEFAULT_TELEMETRY_ORIGIN,
		"Origin of telemetry server",
	)
	Cmd.Flags().BoolVar(
		&FlagDisableTelemetry,
		"telemetry-disable",
		false,
		"Disable telemetry",
	)
}
