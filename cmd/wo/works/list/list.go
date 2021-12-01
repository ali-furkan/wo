package list

import (
	"errors"

	"github.com/ali-furkan/wo/internal/cmdutil"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "list"
	CmdShortDesc = "Show your work list"
	CmdLongDesc  = "Show your work list"
)

var CmdAliases = []string{"ls"}

func NewCmdList(ctx *cmdutil.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listWorks(ctx)
		},
	}

	return cmd
}

func listWorks(ctx *cmdutil.CmdContext) error {
	ws := ctx.Workspaces()

	if len(ws) == 0 {
		return errors.New("works not found")
	}

	m := &model{
		list: ws,
	}
	p := tea.NewProgram(m)

	return p.Start()
}
