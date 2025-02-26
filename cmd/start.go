package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// File creation command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the generation of a new project",
	Run: func(cmd *cobra.Command, args []string) {
		start, _ := cmd.Flags().GetString("start")
		fmt.Println("Forge activated:", start)
	},
}

func init() {
	startCmd.Flags().StringP("start", "s", "", "Starts the generation of a new project (config selection and setup)")
	rootCmd.AddCommand(startCmd)
}
