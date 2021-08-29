package root

import (
	"github.com/sikalabs/tergum/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "tergum",
	Short: "Tergum Backup, " + version.Version,
}

func init() {}
