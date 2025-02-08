package backup

import (
	"github.com/sikalabs/tergum/cmd/root"
	"github.com/sikalabs/tergum/config"
	"github.com/sikalabs/tergum/do_backup"
	"github.com/spf13/cobra"
)

var CmdFlagConfig string
var CmdFlagExtraName string
var FlagDisableTelemetry bool
var FlagJsonLogs bool
var FlagExpandEnv bool
var FlagDoBackupV2 bool
var FlagDebugLogs bool

var Cmd = &cobra.Command{
	Use:     "backup",
	Short:   "Do backup",
	Aliases: []string{"b"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		// Load config from file
		var cfg config.TergumConfig
		cfg.Load(CmdFlagConfig, FlagExpandEnv)

		extraName := cfg.Settings.ExtraName
		if CmdFlagExtraName != "" {
			extraName = CmdFlagExtraName
		}

		if FlagDoBackupV2 || cfg.Settings.UseDoBackupV2 {
			do_backup.DoBackupV2(
				CmdFlagConfig,
				FlagExpandEnv,
				FlagDisableTelemetry,
				extraName,
				FlagJsonLogs,
				FlagDebugLogs,
			)
			return
		}
		do_backup.DoBackup(
			CmdFlagConfig,
			FlagExpandEnv,
			FlagDisableTelemetry,
			extraName,
			FlagJsonLogs,
			FlagDebugLogs,
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
	Cmd.Flags().BoolVar(
		&FlagDebugLogs,
		"debug-logs",
		false,
		"Print debug logs",
	)
	Cmd.Flags().BoolVar(
		&FlagExpandEnv,
		"expand-env",
		false,
		"Expand ENV Variables in YAML",
	)
	Cmd.Flags().BoolVar(
		&FlagDoBackupV2,
		"v2",
		false,
		"!! EXPERIMENTAL !! Use DoBackupV2 backup method (DOES NOT WORK YET)",
	)
}
