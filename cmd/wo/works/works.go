package works

import (
	"github.com/ali-furkan/wo/cmd/wo/works/cmd_init"
	"github.com/ali-furkan/wo/cmd/wo/works/create"
	"github.com/ali-furkan/wo/cmd/wo/works/list"
	"github.com/ali-furkan/wo/cmd/wo/works/move"
	"github.com/ali-furkan/wo/cmd/wo/works/open"
	"github.com/ali-furkan/wo/cmd/wo/works/remove"
	"github.com/ali-furkan/wo/cmd/wo/works/run"
	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/spf13/cobra"
)

func InitCmdWorks(cmd *cobra.Command, ctx *cmdutil.CmdContext) *cobra.Command {

	cmd.AddCommand(create.NewCmdCreate(ctx))
	cmd.AddCommand(cmd_init.NewCmdInit(ctx))
	cmd.AddCommand(remove.NewCmdRemove(ctx))
	cmd.AddCommand(open.NewCmdOpen(ctx))
	cmd.AddCommand(run.NewCmdRun(ctx))
	cmd.AddCommand(list.NewCmdList(ctx))
	cmd.AddCommand(move.NewCmdMove(ctx))

	return cmd
}
