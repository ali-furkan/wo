package cycle

import (
	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/pkg/cycle"
)

const (
	OnCycleStart cycle.CycleNodeType = iota
	OnCycleShutdown
)

func NewCycleRoot(ctx *cmdutil.CmdContext) *cycle.Cycle {
	rootCycle := cycle.NewCycle(ctx)

	// nodes
	rootCycle.AddNode(NewNodeUpdate())
	rootCycle.AddNode(NewNodeConfig())
	rootCycle.AddNode(NewNodeConsole())
	rootCycle.AddNode(NewNodeEditor())

	return rootCycle
}
