package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func getProjectName() string {
	// Get user input for project name
	validate := func(input string) error {
		if strings.Contains(input, " ") {
			return errors.New("Input must not contain spaces")
		}
		return nil
	}
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	p := promptui.Prompt{
		Label:     "Enter a project name: ",
		Validate:  validate,
		Templates: templates,
		IsConfirm: false,
	}

	projectName, err := p.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "\n"
	}
	return projectName

}

func getProjectType() string {
	// Get user input for project type
	prompt := promptui.Select{
		Label: "Select a project type:",
		Items: projectTypes,
	}
	_, projectType, err := prompt.Run()
	if err != nil {
		fmt.Println("Selection failed:", err)
		return ""
	}
	return projectType
}

func getGitProvider() string {
	// Get user input for git provider
	prompt := promptui.Select{
		Label: "Select a git provider",
		Items: GitProviderTypes,
	}
	_, gitProvider, err := prompt.Run()
	if err != nil {
		fmt.Println("Selection failed:", err)
		return ""
	}
	return gitProvider
}

func getCloudProvider() string {
	// Get user input for cloud provider
	prompt := promptui.Select{
		Label: "Select a cloud provider",
		Items: cloudProviderTypes,
	}
	_, cloudProvider, err := prompt.Run()
	if err != nil {
		fmt.Println("Selection failed:", err)
		return ""
	}
	return cloudProvider
}
