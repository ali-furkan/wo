package remove

import (
	"errors"

	"github.com/ali-furkan/wo/internal/config"
	"github.com/ali-furkan/wo/internal/workspace"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "remove <name> "
	CmdShortDesc = "Remove work on wo"
	CmdLongDesc  = "Remove work on wo"
)

var CmdAliases = []string{"rm"}

type RemoveOpts struct {
	Config *config.Config

	Name  string
	Force bool
}

func NewCmdRemove(cfg *config.Config) *cobra.Command {
	opts := &RemoveOpts{
		Config: cfg,
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
	works := opts.Config.Config().Workspace.Works

	for i, work := range works {
		if work.Name == opts.Name {
			err := workspace.RemoveWork(work.Path, opts.Force)
			if err != nil {
				return err
			}

			works = append(works[:i], works[i+1:]...)
			opts.Config.Config().Workspace.Works = works
			return nil
		}
	}

	return errors.New("work is not found")
}
