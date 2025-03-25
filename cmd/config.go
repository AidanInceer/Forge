package cmd

type UserConfig struct {
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

type Config struct {
	Projects         Projects
	GitProviders     GitProvider
	CloudProviders   CloudProvider
	FileToFolders    FileToFolders
	FolderToFiles    FolderToFiles
	BoilerplateFiles BoilerplateFiles
}

type Projects struct {
	Project []Project
}

type BoilerplateFiles struct {
	Files  []string
	isFile bool
}

type FolderToFiles struct {
	Folder []string
	isFile bool
}

type FileToFolders struct {
	File   []string
	isFile bool
}

type GitProvider struct {
	GitProvider map[string][]string
}

type CloudProvider struct {
	CloudProvider map[string][]string
}

type Project struct {
	name      string
	folders   Folders
	rootFiles []Files
}

type Folders struct {
	Folder []Folder
}

type Folder struct {
	name  string
	files Files
}

type Files struct {
	File []File
}

type File struct {
	name string
}
