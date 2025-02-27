package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type baseConfig struct {
	Location string `yaml:"location"`
}

const configFile = "./config.yaml"

func saveConfig(config baseConfig) error {
	file, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer file.Close()

	yamlData, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	_, err = file.Write(yamlData)
	return err
}

func loadConfig() (baseConfig, error) {
	config := baseConfig{}
	file, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return baseConfig{}, nil // Return empty config if file doesn't exist
		}
		return baseConfig{}, err
	}

	err = yaml.Unmarshal(file, &config)
	return config, err
}

func handleLocationInput() string {
	// Get user input for project name
	validate := func(input string) error {
		if strings.Contains(input, " ") {
			return errors.New("Input must not contain spaces")
		}
		return nil
	}

	p := promptui.Prompt{
		Label:    "Enter the root location for your projects",
		Validate: validate,
	}

	rootLocation, err := p.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return rootLocation
}

func resetLocation() string {
	prompt := promptui.Prompt{
		Label:     "Do you want to reset the location?",
		IsConfirm: true,
	}
	resetLocation, _ := prompt.Run()
	return resetLocation
}

func getLocation() string {
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	// need to handle the case where the user is wants to change the location if even if it is already set
	reset := resetLocation()
	fmt.Println(reset)
	if reset == "y" {
		loc := overwriteLocation(config)
		return loc
	} else if config.Location != "" {
		return config.Location
	} else {
		locs := overwriteLocation(config)
		return locs
	}

}

func overwriteLocation(config baseConfig) string {
	// Otherwise, prompt the user
	input := handleLocationInput()

	// Convert to absolute path
	absPath, err := filepath.Abs(input)
	if err != nil {
		fmt.Println("Error resolving absolute path:", err)
		os.Exit(1)
	}

	// Save new Location to config
	config.Location = absPath
	if err := saveConfig(config); err != nil {
		fmt.Println("Error saving config:", err)
		os.Exit(1)
	}

	return absPath

}

// File creation command
var placeCmd = &cobra.Command{
	Use:   "place",
	Short: "`place` the root location for future projects e.g. in a .../coding/ folder.	 Allows the user to redefine the default location for new projects without having to start up all of forge.",
	Run: func(cmd *cobra.Command, args []string) {
		location, _ := cmd.Flags().GetString("location")
		absPath, err := filepath.Abs(location)
		if err != nil {
			fmt.Println("Error getting absolute path:", err)
			return
		}

		cfg := baseConfig{Location: absPath}
		if err := saveConfig(cfg); err != nil {
			fmt.Println("Error saving config:", err)
			return
		}

		fmt.Println("Location set to:", absPath)
	},
}

func init() {
	placeCmd.Flags().StringP("location", "l", ".", "Sets the absolute root path of the project for future projects by default e.g. in a .../coding/ folder")
	rootCmd.AddCommand(placeCmd)
}
