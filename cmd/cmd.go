package cmd

import (
	_ "github.com/sikalabs/tergum/cmd/backup"
	_ "github.com/sikalabs/tergum/cmd/generate_docs"
	"github.com/sikalabs/tergum/cmd/root"
	_ "github.com/sikalabs/tergum/cmd/server"
	_ "github.com/sikalabs/tergum/cmd/utils"
	_ "github.com/sikalabs/tergum/cmd/utils/cron"
	_ "github.com/sikalabs/tergum/cmd/utils/pause"
	_ "github.com/sikalabs/tergum/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
