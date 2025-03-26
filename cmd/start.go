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

		conf := UserConfig{
			RootLocation:  rootLocation,
			ProjectName:   ProjectName,
			ProjectType:   projectType,
			GitProvider:   gitProvider,
			CloudProvider: cloudProvider,
		}

		fmt.Printf("Root Project: %s\n", conf.RootLocation)
		// CreateProject(conf)
		CreateProject(conf)
	},
}

func init() {
	startCmd.Flags().StringP("start", "s", "", "Starts the generation of a new project (config selection and setup)")
}
