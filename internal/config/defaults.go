package config

import (
	"github.com/ali-furkan/wo/internal/editor"
	"github.com/ali-furkan/wo/internal/workspace"
)

var defaultConfigFile = &ConfigFile{
	Editors: []editor.Editor{},
	Workspace: workspace.Workspace{
		DefaultEditor:  "",
		DefaultReadme:  true,
		DefaultRc:      true,
		DefaultLicense: "",
		TempWorkDir:    "",
		WorkDir:        "",
		DefaultGit:     true,
		Works:          []workspace.Work{},
	},
}

const ResourceFileName = ".worc"

var defaultResourceFile = &ResourceFile{
	Name:        "",
	Description: "",
	RunScript:   workspace.Script{},
	Scripts: []workspace.Script{
		{
			Name: "start",
			Run:  "echo 'hello world'",
		},
	},
	Flows: []workspace.Workflows{},
}
