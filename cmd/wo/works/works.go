package works

import (
	"github.com/ali-furkan/wo/cmd/wo/works/cmd_init"
	"github.com/ali-furkan/wo/cmd/wo/works/create"
	"github.com/ali-furkan/wo/cmd/wo/works/list"
	"github.com/ali-furkan/wo/cmd/wo/works/move"
	"github.com/ali-furkan/wo/cmd/wo/works/open"
	"github.com/ali-furkan/wo/cmd/wo/works/remove"
	"github.com/ali-furkan/wo/cmd/wo/works/run"
	"github.com/ali-furkan/wo/internal/config"
	"github.com/spf13/cobra"
)

func InitCmdWorks(cmd *cobra.Command, cfg *config.Config) *cobra.Command {

	cmd.AddCommand(create.NewCmdCreate(cfg))
	cmd.AddCommand(cmd_init.NewCmdInit(cfg))
	cmd.AddCommand(remove.NewCmdRemove(cfg))
	cmd.AddCommand(open.NewCmdOpen(cfg))
	cmd.AddCommand(run.NewCmdRun(cfg))
	cmd.AddCommand(list.NewCmdList(cfg))
	cmd.AddCommand(move.NewCmdMove(cfg))

	return cmd
}
