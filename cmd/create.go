package cmd

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func CreateProject(conf Config) {
	// read the yaml file
	projectConfigPath := "./templates/config/config.yaml"

	// load the yaml file
	config, err := LoadProjectConfig(projectConfigPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Loaded Project Config: %+v\n", config)

	// create the folder at the roote location

	// create new path variable to set root for files to be added to

	// create the project with core files from the config.

	// create the git provider files from the config.

	// create the cloud provider files from the config.

	// add default files

}



func LoadProjectConfig(filePath string) (*ProjectYaml, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config ProjectYaml
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
