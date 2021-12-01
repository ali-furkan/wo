package cycle

import (
	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/pkg/cycle"
)

func NewNodeConsole() *cycle.CycleNode {
	cn := cycle.NewCycleNode()

	cn.AddExe(consoleCycle)

	return cn
}

func consoleCycle(ctx *cmdutil.CmdContext) error {

	return nil
}
