package remove

import (
	"fmt"

	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/internal/space"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "remove <name> "
	CmdShortDesc = "Remove work on wo"
	CmdLongDesc  = "Remove work on wo"

	ErrWsNotFound = "%s workspace not found"
)

var CmdAliases = []string{"rm"}

type RemoveOpts struct {
	Ctx *cmdutil.CmdContext

	Name  string
	Force bool
}

func NewCmdRemove(ctx *cmdutil.CmdContext) *cobra.Command {
	opts := &RemoveOpts{
		Ctx: ctx,
	}

	cmd := &cobra.Command{
		Use:     CmdUsage,
		Aliases: CmdAliases,
		Short:   CmdShortDesc,
		Long:    CmdLongDesc,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]

			return removeWork(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Work name to operate on")
	cmd.Flags().BoolVarP(&opts.Force, "force", "f", false, "remove directories and their files permanently")

	return cmd
}

func removeWork(opts *RemoveOpts) error {
	workspaces := opts.Ctx.Workspaces()

	for id, ws := range workspaces {
		if ws["id"] == opts.Name {
			err := space.RemoveWorkspace(ws["path"], opts.Force)
			if err != nil {
				return err
			}

			delete(workspaces, id)
			return nil
		}
	}

	return fmt.Errorf(ErrWsNotFound, opts.Name)
}
