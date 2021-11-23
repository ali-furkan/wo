package cycle

import (
	"github.com/ali-furkan/wo/internal/config"
	"github.com/ali-furkan/wo/pkg/cycle"
)

func NewNodeConsole() *cycle.CycleNode {
	cn := cycle.NewCycleNode()

	cn.AddExe(consoleCycle)

	return cn
}

func consoleCycle(cfg *config.Config) error {

	return nil
}
