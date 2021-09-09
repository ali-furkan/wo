package list

import (
	"errors"

	"github.com/ali-furkan/wo/internal/config"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "list"
	CmdShortDesc = "Show your work list"
	CmdLongDesc  = "Show your work list"
)

var CmdAliases = []string{"ls"}

func NewCmdList(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listWorks(cfg)
		},
	}

	return cmd
}

func listWorks(cfg *config.Config) error {
	if cfg.Workspace() == nil {
		return errors.New("works not found")
	}

	workspace := *cfg.Workspace()

	m := &model{
		list: workspace.Works,
	}
	p := tea.NewProgram(m)

	return p.Start()
}
