package cmd

import (
	"github.com/esteam85/interviews-tracker-cli/cmd/add"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "itracker",
	Short: "Interview Tracker, to manage interviews for devs",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommands(cmd *cobra.Command) {
	cmd.AddCommand(add.AddCmd)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommands(rootCmd)
}
