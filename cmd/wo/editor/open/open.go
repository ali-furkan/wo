package open

import (
	"errors"

	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/internal/editor"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "open"
	CmdShortDesc = "Open Editor"
	CmdLongDesc  = "Manage your editors"
)

type OpenOpts struct {
	Ctx *cmdutil.CmdContext

	SelectedEditor string
	Path           string
}

func NewCmdOpen(ctx *cmdutil.CmdContext) *cobra.Command {
	opts := &OpenOpts{
		Ctx: ctx,
	}

	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
		Args:  cobra.MaximumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.SelectedEditor = ctx.Defaults()["editor"]
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
	editors := opts.Ctx.Editors()

	for _, e := range editors {
		if e["id"] == opts.SelectedEditor {
			return editor.OpenEditor(editor.Editor{
				Name: e["id"],
				Exec: e["exec"],
			}, opts.Path)
		}
	}

	return errors.New("editor open failed: unknown editor")
}
