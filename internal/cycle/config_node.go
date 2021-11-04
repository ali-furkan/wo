package cycle

import (
	"os"
	"path/filepath"

	"github.com/ali-furkan/wo/internal/config"
)

func NewNodeConfig() *CycleNode {
	cn := NewCycleNode()
	cn.Type = OnCycleShutdown

	cn.AddExe(resourceConfigSyncCycle)
	cn.AddExe(configWriteCycle)

	return cn
}

func configWriteCycle(cfg *config.Config) error {
	return cfg.Write()
}

func resourceConfigSyncCycle(cfg *config.Config) error {
	if cfg.Resource() == nil {
		return nil
	}

	if cfg.Resource().Name == "" && cfg.Resource().Description == "" {
		return nil
	}

	isChange := false

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	path = filepath.Clean(path)

	for _, work := range cfg.Config().Workspace.Works {
		if filepath.Clean(work.Path) != path {
			continue
		}

		if cfg.Resource().Name != work.Name {
			isChange = true
			cfg.Resource().Name = work.Name
		}

		if cfg.Resource().Description != work.Description {
			isChange = true
			cfg.Resource().Description = work.Description
		}

		break
	}

	if isChange {
		return cfg.WriteResourceFile()
	}

	return nil
}
