package utils

import (
	"github.com/sikalabs/tergum/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "utils",
	Short:   "Random utils",
	Aliases: []string{"u"},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
