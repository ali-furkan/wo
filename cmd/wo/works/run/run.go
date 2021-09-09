package run

import (
	"github.com/ali-furkan/wo/internal/config"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "run"
	CmdShortDesc = "Run scripts of work"
	CmdLongDesc  = "Run the specified script in your work"
)

func NewCmdRun(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
	}

	return cmd
}
