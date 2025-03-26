package cmd

type UserConfig struct {
	ProjectName   string
	ProjectType   string
	RootLocation  string
	GitProvider   string
	CloudProvider string
}

type Folder struct {
	Name    string
	Files   []string
	Folders []Folder
}

var DefaultProjectStructure = Folder{
	Name: "",
	Files: []string{
		"main.py",
		"requirements.txt",
		"README.md",
	},
	Folders: []Folder{
		{
			Name: "src",
			Files: []string{
				"__init__.py",
			},
			Folders: []Folder{
				{
					Name: "core",
					Files: []string{
						"__init__.py",
						"config.py",
						"database.py",
					},
				},
				{
					Name: "models",
					Files: []string{
						"__init__.py",
						"user.py",
						"base.py",
					},
				},
				{
					Name: "services",
					Files: []string{
						"__init__.py",
						"auth.py",
						"api.py",
					},
				},
			},
		},
		{
			Name: "tests",
			Files: []string{
				"__init__.py",
				"conftest.py",
			},
			Folders: []Folder{
				{
					Name: "unit",
					Files: []string{
						"__init__.py",
						"test_auth.py",
						"test_api.py",
					},
				},
				{
					Name: "integration",
					Files: []string{
						"__init__.py",
						"test_database.py",
					},
				},
			},
		},
	},
}

var BasicProjectStructure = Folder{
	Name: "",
	Files: []string{
		"main.py",
		"requirements.txt",
		"README.md",
	},
	Folders: []Folder{
		{
			Name: "src",
			Files: []string{
				"__init__.py",
			},
			Folders: []Folder{
				{
					Name: "core",
					Files: []string{
						"__init__.py",
						"config.py",
					},
				},
				{
					Name: "models",
					Files: []string{
						"__init__.py",
						"base.py",
					},
				},
			},
		},
		{
			Name: "tests",
			Files: []string{
				"__init__.py",
				"test_main.py",
			},
		},
	},
}

var FlaskProjectStructure = Folder{
	Name: "",
	Files: []string{
		"app.py",
		"requirements.txt",
		"README.md",
	},
	Folders: []Folder{
		{
			Name: "app",
			Files: []string{
				"__init__.py",
			},
			Folders: []Folder{
				{
					Name: "templates",
					Files: []string{
						"base.html",
						"index.html",
					},
				},
				{
					Name: "static",
					Files: []string{
						"style.css",
						"script.js",
					},
				},
				{
					Name: "models",
					Files: []string{
						"__init__.py",
						"user.py",
					},
				},
				{
					Name: "routes",
					Files: []string{
						"__init__.py",
						"auth.py",
						"main.py",
					},
				},
				{
					Name: "forms",
					Files: []string{
						"__init__.py",
						"auth.py",
					},
				},
			},
		},
		{
			Name: "tests",
			Files: []string{
				"__init__.py",
				"conftest.py",
				"test_auth.py",
			},
		},
	},
}

var GitProviderConfig = map[string]Folder{
	"github": {
		Name: ".github",
		Files: []string{
			"README.md",
		},
		Folders: []Folder{
			{
				Name: "workflows",
				Files: []string{
					"ci.yml",
					"cd.yml",
				},
			},
		},
	},
	"gitlab": {
		Name: ".gitlab",
		Files: []string{
			"README.md",
		},
		Folders: []Folder{
			{
				Name: "ci",
				Files: []string{
					"ci.yml",
				},
			},
		},
	},
}

var CloudProviderConfig = map[string]Folder{
	"AWS": {
		Name: "aws",
		Files: []string{
			"README.md",
		},
		Folders: []Folder{
			{
				Name: "terraform",
				Files: []string{
					"main.tf",
					"variables.tf",
					"outputs.tf",
				},
			},
			{
				Name: "scripts",
				Files: []string{
					"deploy.sh",
				},
			},
		},
	},
	"GCP": {
		Name: "gcp",
		Files: []string{
			"README.md",
		},
		Folders: []Folder{
			{
				Name: "terraform",
				Files: []string{
					"main.tf",
					"variables.tf",
					"outputs.tf",
				},
			},
			{
				Name: "scripts",
				Files: []string{
					"deploy.sh",
				},
			},
		},
	},
	"AZURE": {
		Name: "azure",
		Files: []string{
			"README.md",
		},
		Folders: []Folder{
			{
				Name: "terraform",
				Files: []string{
					"main.tf",
					"variables.tf",
					"outputs.tf",
				},
			},
			{
				Name: "scripts",
				Files: []string{
					"deploy.sh",
				},
			},
		},
	},
}

// List of project types
var projectTypes = []string{"Flask", "Default", "Basic"}
var GitProviderTypes = []string{"Github", "Azure DevOps", "Gitlab", "Bitbucket", "None"}
var cloudProviderTypes = []string{"AWS", "GCP", "AZURE", "None"}
