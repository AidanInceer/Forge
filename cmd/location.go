package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Location string `yaml:"location"`
}

const configFile = "./config.yaml"

func SaveConfig(config Config) error {
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

func LoadConfig() (Config, error) {
	config := Config{}
	file, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{}, nil // Return empty config if file doesn't exist
		}
		return Config{}, err
	}

	err = yaml.Unmarshal(file, &config)
	return config, err
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

		cfg := Config{Location: absPath}
		if err := SaveConfig(cfg); err != nil {
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
