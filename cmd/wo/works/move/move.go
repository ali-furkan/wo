package move

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/ali-furkan/wo/internal/config"
	"github.com/ali-furkan/wo/internal/workspace"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "move <name> [path]"
	CmdShortDesc = "Edit Work"
	CmdLongDesc  = "Edit Work"
)

type MoveOpts struct {
	Config *config.Config

	Name string
	Path string
}

func NewCmdMove(cfg *config.Config) *cobra.Command {
	opts := &MoveOpts{
		Config: cfg,
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
	ws := opts.Config.Config().Workspace

	for _, w := range ws.Works {
		if w.Name == opts.Name {
			oldPath := w.Path
			if isMatch, err := filepath.Match(oldPath, w.Path); isMatch && err != nil {
				return errors.New("%s can't move to the same path")
			}

			err := workspace.MoveWork(&w, opts.Path)
			if err != nil {
				return err
			}

			fmt.Printf("%s moved '%s' '%s'", w.Name, oldPath, opts.Path)

			return nil
		}
	}

	return errors.New("specified work not found")
}
