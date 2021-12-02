package cycle

import (
	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/pkg/cycle"
)

func NewNodeConfig() *cycle.CycleNode {
	cn := cycle.NewCycleNode()
	cn.Name = "config"
	cn.Type = OnCycleShutdown

	cn.AddExe(configWriteCycle)

	return cn
}

func configWriteCycle(ctx *cmdutil.CmdContext) error {
	config, err := ctx.Config()
	if err != nil {
		return err
	}

	return config.Write()
}
