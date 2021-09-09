package config

import (
	"github.com/ali-furkan/wo/internal/auth"
	"github.com/ali-furkan/wo/internal/editor"
	"github.com/ali-furkan/wo/internal/workspace"
)

type ConfigFile struct {
	Auth      map[string]auth.Auth `yaml:"auth"`
	Editors   []editor.Editor      `yaml:"editors"`
	Workspace workspace.Workspace  `yaml:"workspace"`
}
