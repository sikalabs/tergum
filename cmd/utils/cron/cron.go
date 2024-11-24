package cron

import (
	"log"
	"os"
	"os/exec"
	"strings"

	robfig_cron "github.com/robfig/cron/v3"

	parentcmd "github.com/sikalabs/tergum/cmd/utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "cron",
	Short: "cron <cron-expression> <command> [args...]",
	Args:  cobra.MinimumNArgs(2),
	Run: func(c *cobra.Command, args []string) {
		cron(args)
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
}

func cron(args []string) {
	cronExpression := args[0]
	command := args[1]
	args = args[2:]

	c := robfig_cron.New()
	_, err := c.AddFunc(cronExpression, func() {
		log.Printf("Executing command: %s %s", command, strings.Join(args, " "))
		cmd := exec.Command(command, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Printf("Error executing command: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}

	log.Printf("Cron scheduler started with expression: %s, command: %s %s", cronExpression, command, strings.Join(args, " "))
	c.Start()

	// Keep the program running
	select {}
}
