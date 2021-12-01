package list

import (
	"errors"
	"fmt"

	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "list"
	CmdShortDesc = "Show list of editor"
	CmdLongDesc  = "Manage your editors"

	ErrNotFoundEditor = "wo cant found editor. you can download one of the following editors"
)

func NewCmdSetEditor(ctx *cmdutil.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listEditors(ctx)
		},
	}

	return cmd
}

func listEditors(ctx *cmdutil.CmdContext) error {
	c, err := ctx.Config()
	if err != nil {
		return err
	}

	editors := c.Get("editors").(map[string]map[string]string)
	if len(editors) == 0 {
		return errors.New(ErrNotFoundEditor)
	}

	res := fmt.Sprintf("Showing %d list of edit\n\n", len(editors))

	def_editor := c.GetString("defaults.editor")

	for _, e := range editors {
		res += fmt.Sprintf("%s - %s", e["id"], e["exec"])
		if e["id"] == def_editor {
			res += fmt.Sprintln(" 'current editor'")
		} else {
			res += "\n"
		}
	}

	fmt.Println(res)

	return nil
}
