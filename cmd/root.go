package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "forge",
	Short: "Forge is a CLI tool create the template for youre next project",
}

func Execute() {
	rootCmd.Execute()
}
