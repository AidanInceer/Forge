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

		
		rootLocation := getLocation()
		ProjectName := getProjectName()
		projectType := getProjectType()
		gitProvider := getGitProvider()
		cloudProvider := getCloudProvider()

		fmt.Printf("Root Project: %s\nName: %s\nProject Type: %s\nGit Provider: %s\nCloud Provider: %s\n", rootLocation, ProjectName, projectType, gitProvider, cloudProvider)
	},
}

func init() {
	startCmd.Flags().StringP("start", "s", "", "Starts the generation of a new project (config selection and setup)")
	rootCmd.AddCommand(startCmd)
}
