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
		DefaultLicense: "",
		TempWorkDir:    "",
		WorkDir:        "",
		DefaultGit:     true,
		Works:          []workspace.Work{},
	},
}
