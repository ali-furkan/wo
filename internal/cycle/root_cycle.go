package cycle

import (
	"github.com/ali-furkan/wo/internal/config"
	"github.com/ali-furkan/wo/pkg/cycle"
)

func NewCycleRoot(cfg *config.Config) *cycle.Cycle {
	rootCycle := cycle.NewCycle(cfg)

	// nodes
	rootCycle.AddNode(NewNodeUpdate())
	rootCycle.AddNode(NewNodeConfig())
	rootCycle.AddNode(NewNodeConsole())
	rootCycle.AddNode(NewNodeEditor())

	return rootCycle
}
