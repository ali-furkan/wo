package config

import (
	"github.com/ali-furkan/wo/internal/editor"
	"github.com/ali-furkan/wo/internal/workspace"
)

type ConfigFile struct {
	Editors   []editor.Editor     `yaml:"editors"`
	Workspace workspace.Workspace `yaml:"workspace"`
}

type ResourceFile struct {
	Name        string                `yaml:"name"`
	Description string                `yaml:"description"`
	RunScript   workspace.Script      `yaml:"run"`
	Scripts     []workspace.Script    `yaml:"scripts"`
	Flows       []workspace.Workflows `yaml:"flows"`
}
