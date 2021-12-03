package open

import (
	"fmt"
	"os"

	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/internal/editor"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "open"
	CmdShortDesc = "Open work with wo"
	CmdLongDesc  = "Open work with wo"

	ErrWorkspaceNotFound = "%s workspace not found"
	ErrUnknownEditor     = "%s editor not found"
)

var CmdAliases = []string{"o"}

type OpenOpts struct {
	Ctx *cmdutil.CmdContext

	WsName         string
	SelectedEditor string
}

func NewCmdOpen(ctx *cmdutil.CmdContext) *cobra.Command {
	opts := &OpenOpts{
		Ctx: ctx,
	}

	cmd := &cobra.Command{
		Use:     CmdUsage,
		Short:   CmdShortDesc,
		Long:    CmdLongDesc,
		Args:    cobra.MaximumNArgs(2),
		Aliases: CmdAliases,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.WsName = args[0]
			}
			if len(args) > 1 {
				opts.SelectedEditor = args[1]
			}
			return openWork(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.WsName, "name", "n", "", "Work name to operate on")
	cmd.Flags().StringVarP(&opts.SelectedEditor, "editor", "e", "", "Open work with specified editor")

	return cmd
}

func openWork(opts *OpenOpts) error {
	c, err := opts.Ctx.Config()
	if err != nil {
		return err
	}

	var wsPath string

	for _, w := range opts.Ctx.Workspaces() {
		if w["id"] == opts.WsName {
			wsPath = w["path"]
			break
		}
	}

	_, err = opts.Ctx.WorkspaceRC()
	if err == nil && wsPath == "" {
		d, _ := os.Getwd()
		wsPath = d
	}

	if wsPath == "" {
		return fmt.Errorf(ErrWorkspaceNotFound, opts.WsName)
	}

	if opts.SelectedEditor == "" {
		opts.SelectedEditor = c.GetString("defaults.editor")
	}
	selectedEditorPath := fmt.Sprintf("editors.%s", opts.SelectedEditor)
	e, ok := c.Get(selectedEditorPath).(map[string]interface{})
	if !ok {
		return fmt.Errorf(ErrUnknownEditor, opts.SelectedEditor)
	}

	selectedEditor := editor.Editor{
		Name: e["id"].(string),
		Exec: e["exec"].(string),
	}

	return editor.OpenEditor(selectedEditor, wsPath)
}
