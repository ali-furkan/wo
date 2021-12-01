package set

import (
	"fmt"

	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "set"
	CmdShortDesc = "Set default editor"
	CmdLongDesc  = "Manage your editors"
)

type SetEditorOpts struct {
	Ctx *cmdutil.CmdContext

	SelectedEditor string
}

func NewCmdSetEditor(ctx *cmdutil.CmdContext) *cobra.Command {
	opts := &SetEditorOpts{
		Ctx: ctx,
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
	config, err := opts.Ctx.Config()
	if err != nil {
		return err
	}

	field := fmt.Sprintf("editors.%s", opts.SelectedEditor)
	editor := config.Get(field).(map[string]string)
	if editor == nil {
		return fmt.Errorf("specified '%s' editor not found", opts.SelectedEditor)
	}

	err = config.Set("defaults.editor", editor["id"])
	if err != nil {
		return err
	}

	fmt.Printf("\n%s set default editor\n", editor["id"])
	return nil
}
