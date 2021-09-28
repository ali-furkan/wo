package config

import (
	"github.com/ali-furkan/wo/internal/editor"
	"github.com/ali-furkan/wo/internal/workspace"
)

type ConfigFile struct {
	Editors   []editor.Editor      `yaml:"editors"`
	Workspace workspace.Workspace  `yaml:"workspace"`
}
