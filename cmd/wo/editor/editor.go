package editor

import (
	"github.com/ali-furkan/wo/cmd/wo/editor/list"
	"github.com/ali-furkan/wo/cmd/wo/editor/open"
	"github.com/ali-furkan/wo/cmd/wo/editor/set"
	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "editor"
	CmdShortDesc = "Manage your editors"
	CmdLongDesc  = "Manage your editors"
)

func NewCmdEditor(ctx *cmdutil.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
	}

	cmd.AddCommand(open.NewCmdOpen(ctx))
	cmd.AddCommand(set.NewCmdSetEditor(ctx))
	cmd.AddCommand(list.NewCmdSetEditor(ctx))

	return cmd
}
