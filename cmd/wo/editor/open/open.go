package open

import (
	"errors"

	"github.com/ali-furkan/wo/internal/config"
	"github.com/ali-furkan/wo/internal/editor"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "open"
	CmdShortDesc = "Open Editor"
	CmdLongDesc  = "Manage your editors"
)

type OpenOpts struct {
	Config *config.Config

	SelectedEditor string
	Path           string
}

func NewCmdOpen(cfg *config.Config) *cobra.Command {
	opts := &OpenOpts{
		Config: cfg,
	}

	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
		Args:  cobra.MaximumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.SelectedEditor = cfg.Workspace().DefaultEditor
			if len(args) > 0 {
				opts.SelectedEditor = args[0]
			}
			if len(args) > 1 {
				opts.Path = args[1]
			}

			return openEditor(opts)
		},
	}

	return cmd
}

func openEditor(opts *OpenOpts) error {
	editors := opts.Config.Editors(nil)

	for _, e := range *editors {
		if e.Name == opts.SelectedEditor {
			return editor.OpenEditor(e, opts.Path)
		}
	}

	return errors.New("editor open failed: unknown editor")
}
