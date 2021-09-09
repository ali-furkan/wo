package editor

import (
	"github.com/ali-furkan/wo/cmd/wo/editor/list"
	"github.com/ali-furkan/wo/cmd/wo/editor/open"
	"github.com/ali-furkan/wo/cmd/wo/editor/set"
	"github.com/ali-furkan/wo/internal/config"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "editor"
	CmdShortDesc = "Manage your editors"
	CmdLongDesc  = "Manage your editors"
)

func NewCmdEditor(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
	}

	cmd.AddCommand(open.NewCmdOpen(cfg))
	cmd.AddCommand(set.NewCmdSetEditor(cfg))
	cmd.AddCommand(list.NewCmdSetEditor(cfg))

	return cmd
}
