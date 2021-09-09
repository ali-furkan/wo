package cycle

import "github.com/ali-furkan/wo/internal/config"

func NewNodeConsole() *CycleNode {
	cn := NewCycleNode()

	cn.AddExe(consoleCycle)

	return cn
}

func consoleCycle(cfg *config.Config) error {

	return nil
}
