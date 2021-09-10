package list

import (
	"errors"
	"fmt"

	"github.com/ali-furkan/wo/internal/config"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "list"
	CmdShortDesc = "Show list of editor"
	CmdLongDesc  = "Manage your editors"

	ErrNotFoundEditor = "wo cant found editor. you can download one of the following editors"
)

func NewCmdSetEditor(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listEditors(cfg)
		},
	}

	return cmd
}

func listEditors(cfg *config.Config) error {
	editors := cfg.Config().Editors
	if len(editors) == 0 {
		return errors.New(ErrNotFoundEditor)
	}

	res := fmt.Sprintf("Showing %d list of edit\n\n", len(editors))

	def_editor := cfg.Config().Workspace.DefaultEditor

	for _, e := range editors {
		res += fmt.Sprintf("%s - %s", e.Name, e.Exec)
		if e.Name == def_editor {
			res += fmt.Sprintln(" 'current editor'")
		} else {
			res += "\n"
		}
	}

	fmt.Println(res)

	return nil
}
