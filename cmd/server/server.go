package run

import (
	"github.com/sikalabs/tergum/cmd/root"
	"github.com/sikalabs/tergum/server"
	"github.com/spf13/cobra"
)

var FlagAddr string

var Cmd = &cobra.Command{
	Use:     "server",
	Short:   "Run server",
	Aliases: []string{"s"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		server.Server(FlagAddr)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagAddr,
		"address",
		"a",
		":8000",
		"Bind address (host:port)",
	)
}
