package cmd

type Config struct {
	RootLocation  string
	ProjectName   string
	ProjectType   string
	GitProvider   string
	CloudProvider string
}

// List of project types
var projectTypes = []string{"Flask", "Default", "Basic"}
var GitProviderTypes = []string{"Github", "Azure DevOps", "Gitlab", "Bitbucket", "None"}
var cloudProviderTypes = []string{"AWS", "GCP", "AZURE", "None"}

type ProjectYaml struct {
	Projects         ProjectDetails
	GitProviders     map[string][]interface{}
	CloudProviders   map[string][]interface{}
	FileToFolders    []string
	FolderToFiles    []string
	BoilerplateFiles []string
}

type ProjectDetails struct {
	Core  []map[string][]string
	Tests []map[string][]string
	Docs  []string
	Lib   []string
	Logs  []map[string][]string
	Tools []string
	Files []string
}
