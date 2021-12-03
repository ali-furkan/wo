package cmdutil

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/ali-furkan/wo/internal/config"
)

type CmdContext struct {
	// Extensions

	Config      func() (config.Config, error)
	SpaceRC     func() (config.Config, error)
	WorkspaceRC func() (config.Config, error)
}

func NewCmdContext() *CmdContext {
	c := &CmdContext{
		Config: configFunc(),
	}

	c.SpaceRC = spaceRCFunc(c)
	c.WorkspaceRC = wsRCFunc(c)

	return c
}

func (cc *CmdContext) Workspaces() map[string]map[string]string {
	c, err := cc.Config()
	if err != nil {
		return nil
	}
	return GetWorkspacesFromConfig(c)
}

func (cc *CmdContext) Editors() map[string]map[string]string {
	c, err := cc.Config()
	if err != nil {
		return nil
	}
	return c.Get("editors").(map[string]map[string]string)
}

func (cc *CmdContext) Spaces() map[string]map[string]interface{} {
	c, err := cc.Config()
	if err != nil {
		return nil
	}
	return c.Get("spaces").(map[string]map[string]interface{})
}

func (cc *CmdContext) Defaults() map[string]string {
	c, err := cc.Config()
	if err != nil {
		return nil
	}
	return c.Get("defaults").(map[string]string)
}

func configFunc() func() (config.Config, error) {
	var (
		cachedConfig config.Config
		cachedError  error
	)
	return func() (config.Config, error) {
		if cachedConfig != nil || cachedError != nil {
			return cachedConfig, cachedError
		}

		cachedConfig, cachedError = config.NewGlobalConfig()
		if errors.Is(cachedError, os.ErrNotExist) {
			cachedConfig, cachedError = config.NewBlankConfig()
		}

		return cachedConfig, cachedError
	}
}

func spaceRCFunc(cc *CmdContext) func() (config.Config, error) {
	var (
		cachedRC  config.Config
		cachedErr error
	)
	return func() (config.Config, error) {
		if cachedRC != nil || cachedErr != nil {
			return cachedRC, cachedErr
		}

		c, cachedErr := cc.Config()
		spaces := c.Get("spaces").(map[string]map[string]string)

		wd, err := os.Getwd()
		if err != nil {
			cachedErr = err
			return cachedRC, cachedErr
		}

		for _, space := range spaces {
			if strings.HasPrefix(wd, space["root_dir"]) || strings.HasPrefix(wd, space["temp_dir"]) {
				source := filepath.Join(space["root_dir"], config.SpaceRCFileName)
				cachedRC, cachedErr = config.NewConfig(source, config.SpaceRCSchema)
				break
			}
		}

		return cachedRC, cachedErr
	}
}

func wsRCFunc(cc *CmdContext) func() (config.Config, error) {
	var (
		cachedRC  config.Config
		cachedErr error
	)
	return func() (config.Config, error) {
		if cachedRC != nil || cachedErr != nil {
			return cachedRC, cachedErr
		}

		wd, err := os.Getwd()
		if err != nil {
			cachedErr = err
			return cachedRC, cachedErr
		}

		workspaces := cc.Workspaces()

		for _, ws := range workspaces {
			if strings.HasPrefix(wd, ws["path"]) {
				source := filepath.Join(ws["path"], config.WorkspaceFileName)
				cachedRC, cachedErr = config.NewConfig(source, config.WorkspaceRCSchema)
				break
			}
		}

		return cachedRC, cachedErr
	}
}
