package cmd

import (
	"os"
	"path/filepath"
	"strings"
)

// createFolderStructure recursively creates folders and files from a Folder structure
func createFolderStructure(basePath string, folder Folder) error {
	// Skip folder creation for root folder since it's already created
	if folder.Name != "" {
		folderPath := filepath.Join(basePath, folder.Name)
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			return err
		}
		println("Created folder: ", folderPath)
		basePath = folderPath
	}

	// Create files in the current folder
	for _, file := range folder.Files {
		filePath := filepath.Join(basePath, file)
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}
		if _, err := os.Create(filePath); err != nil {
			return err
		}
		println("Created file: ", filePath)
	}

	// Recursively create subfolders
	for _, subfolder := range folder.Folders {
		if err := createFolderStructure(basePath, subfolder); err != nil {
			return err
		}
	}

	return nil
}

func CreateProject(config UserConfig) {
	// create the folder at the root location
	projectPath := filepath.Join(config.RootLocation, config.ProjectName)
	err := os.MkdirAll(projectPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	println("Project created at: ", projectPath)

	// Validate project type
	projectType := strings.ToLower(config.ProjectType)

	// Select folder structure
	var folderStructure Folder
	switch projectType {
	case "flask":
		folderStructure = FlaskProjectStructure
	case "basic":
		folderStructure = BasicProjectStructure
	default:
		folderStructure = DefaultProjectStructure
	}

	println("Project Type: ", folderStructure.Name)

	// Create root directory files first
	for _, file := range folderStructure.Files {
		filePath := filepath.Join(projectPath, file)
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}
		if _, err := os.Create(filePath); err != nil {
			panic(err)
		}
		println("Created root file: ", filePath)
	}

	// Create the project structure
	if err := createFolderStructure(projectPath, folderStructure); err != nil {
		panic(err)
	}

	// Create Git provider files
	if gitFolder, exists := GitProviderConfig[strings.ToLower(config.GitProvider)]; exists {
		println("Creating Git provider files for: ", config.GitProvider)
		if err := createFolderStructure(projectPath, gitFolder); err != nil {
			panic(err)
		}
	}

	// Create Cloud provider files
	if cloudFolder, exists := CloudProviderConfig[strings.ToUpper(config.CloudProvider)]; exists {
		println("Creating Cloud provider files for: ", config.CloudProvider)
		if err := createFolderStructure(projectPath, cloudFolder); err != nil {
			panic(err)
		}
	}
}
