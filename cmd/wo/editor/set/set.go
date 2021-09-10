package set

import (
	"errors"
	"fmt"

	"github.com/ali-furkan/wo/internal/config"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "set"
	CmdShortDesc = "Set default editor"
	CmdLongDesc  = "Manage your editors"
)

type SetEditorOpts struct {
	Config *config.Config

	SelectedEditor string
}

func NewCmdSetEditor(cfg *config.Config) *cobra.Command {
	opts := &SetEditorOpts{
		Config: cfg,
	}

	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.SelectedEditor = args[0]

			return setEditor(opts)
		},
	}

	return cmd
}

func setEditor(opts *SetEditorOpts) error {
	c := opts.Config.Config()

	for _, editor := range c.Editors {
		if editor.Name == opts.SelectedEditor {
			c.Workspace.DefaultEditor = editor.Name

			fmt.Printf("\n%s set default editor\n", editor.Name)

			return nil
		}
	}

	err := fmt.Sprintf("specified `%s` editor not found", opts.SelectedEditor)
	return errors.New(err)
}
