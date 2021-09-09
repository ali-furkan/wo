package cycle

import "github.com/ali-furkan/wo/internal/config"

func NewCycleRoot(cfg *config.Config) *Cycle {
	rootCycle := NewCycle(cfg)

	// nodes
	rootCycle.AddNode(NewNodeUpdate())
	rootCycle.AddNode(NewNodeConfig())
	rootCycle.AddNode(NewNodeConsole())
	rootCycle.AddNode(NewNodeEditor())

	return rootCycle
}
