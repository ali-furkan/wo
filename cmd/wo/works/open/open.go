package open

import (
	"errors"

	"github.com/ali-furkan/wo/internal/config"
	"github.com/ali-furkan/wo/internal/editor"
	"github.com/ali-furkan/wo/internal/workspace"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "open"
	CmdShortDesc = "Open work with wo"
	CmdLongDesc  = "Open work with wo"
)

var CmdAliases = []string{"o"}

type OpenOpts struct {
	Config *config.Config

	WorkName       string
	SelectedEditor string
}

func NewCmdOpen(cfg *config.Config) *cobra.Command {
	opts := &OpenOpts{
		Config: cfg,
	}

	cmd := &cobra.Command{
		Use:     CmdUsage,
		Short:   CmdShortDesc,
		Long:    CmdLongDesc,
		Args:    cobra.MaximumNArgs(2),
		Aliases: CmdAliases,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.WorkName = args[0]
			}
			if len(args) > 1 {
				opts.SelectedEditor = args[1]
			}
			return openWork(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.WorkName, "name", "n", "", "Work name to operate on")
	cmd.Flags().StringVarP(&opts.SelectedEditor, "editor", "e", cfg.Config().Workspace.DefaultEditor, "Open work with specified editor")

	return cmd
}

func openWork(opts *OpenOpts) error {
	editors := opts.Config.Config().Editors
	var work *workspace.Work

	for _, w := range opts.Config.Config().Workspace.Works {
		if w.Name == opts.WorkName {
			work = &w
			break
		}
	}

	if work == nil {
		return errors.New("open work failed: unknown work")
	}

	for _, e := range editors {
		if e.Name == opts.SelectedEditor {
			return editor.OpenEditor(e, work.Path)
		}
	}

	if len(editors) > 0 {
		return editor.OpenEditor(editors[0], work.Path)
	}

	return errors.New("open work failed: unknown editor")
}
