package cmd

type Config struct {
	Location      string
	ProjectName   string
	ProjectType   string
	GitProvider   string
	CloudProvider string
}

// List of project types
var projectTypes = []string{"Flask", "Default", "Basic"}
var GitProviderTypes = []string{"Github", "Azure DevOps", "Gitlab", "Bitbucket", "None"}
var cloudProviderTypes = []string{"AWS", "GCP", "AZURE", "None"}
