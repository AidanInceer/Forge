package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "forge",
	Short: "Forge is a CLI tool create the template for youre next project",
}

// RegisterCommands registers all commands with the root command
func RegisterCommands() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(placeCmd)
}

func Execute() {
	RegisterCommands()
	rootCmd.Execute()
}
