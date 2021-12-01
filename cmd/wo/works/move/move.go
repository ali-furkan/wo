package move

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/internal/space"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "move <name> [path]"
	CmdShortDesc = "Edit Work"
	CmdLongDesc  = "Edit Work"
)

type MoveOpts struct {
	Ctx *cmdutil.CmdContext

	Name string
	Path string
}

func NewCmdMove(ctx *cmdutil.CmdContext) *cobra.Command {
	opts := &MoveOpts{
		Ctx: ctx,
	}

	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			opts.Path = "."
			if len(args) > 1 {
				opts.Path = args[1]
			}

			return moveWorks(opts)
		},
	}

	return cmd
}

func moveWorks(opts *MoveOpts) error {
	for _, w := range opts.Ctx.Workspaces() {
		if w["id"] == opts.Name {
			oldPath := w["path"]
			if isMatch, err := filepath.Match(oldPath, w["path"]); isMatch && err != nil {
				return errors.New("%s can't move to the same path")
			}

			err := space.MoveWorkspace(w, opts.Path)
			if err != nil {
				return err
			}

			w["path"] = opts.Path

			fmt.Printf("%s moved '%s' '%s'", w["id"], oldPath, opts.Path)

			return nil
		}
	}

	return errors.New("specified work not found")
}
