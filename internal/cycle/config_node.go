package cycle

import "github.com/ali-furkan/wo/internal/config"

func NewNodeConfig() *CycleNode {
	cn := NewCycleNode()
	cn.Type = OnCycleShutdown

	cn.AddExe(configWriteCycle)

	return cn
}

func configWriteCycle(cfg *config.Config) error {
	return cfg.Write()
}
