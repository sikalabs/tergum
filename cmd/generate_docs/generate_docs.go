package version

import (
	"os"

	"github.com/sikalabs/tergum/cmd/root"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var Cmd = &cobra.Command{
	Use:   "generate-docs",
	Short: "Generate Markdown docs",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		path := "./cobra-docs/"
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
		err = doc.GenMarkdownTree(root.Cmd, path)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
