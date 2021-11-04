package run

import (
	"errors"

	"github.com/ali-furkan/wo/internal/config"
	"github.com/ali-furkan/wo/internal/workspace"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "run [script]"
	CmdShortDesc = "Run scripts of work"
	CmdLongDesc  = "Run the specified script in your work"

	// Errors
	ErrResourceFileNotFound = "resource file not found"
	ErrScriptNotFound       = "script not found"
)

type RunOpts struct {
	Config *config.Config

	Name       string
	WorkingDir string
	Watch      bool
	Quiet      bool
	Args       []string
	Env        []string
}

func NewCmdRun(cfg *config.Config) *cobra.Command {
	opts := &RunOpts{
		Config: cfg,
	}

	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.Name = args[0]
			}

			if len(args) > 1 && len(opts.Args) == 0 {
				opts.Args = args[1:]
			}

			return runWork(opts)
		},
	}

	cmd.Flags().StringArrayVar(&opts.Env, "env", []string{}, "Env")
	cmd.Flags().StringArrayVar(&opts.Args, "args", []string{}, "Args")
	cmd.Flags().BoolVarP(&opts.Watch, "watch", "w", false, "Restart")
	cmd.Flags().BoolVarP(&opts.Quiet, "quiet", "q", false, "Quite")
	cmd.Flags().StringVarP(&opts.WorkingDir, "working-dir", "d", "", "Dir")

	return cmd
}

func runWork(opts *RunOpts) error {
	if opts.Config.Resource() == nil {
		return errors.New(ErrResourceFileNotFound)
	}

	var s *workspace.Script

	if opts.Name == "" {
		s = &opts.Config.Resource().RunScript
	}

	for _, script := range opts.Config.Resource().Scripts {
		if script.Name == opts.Name {
			s = &script
			break
		}
	}

	if s == nil {
		return errors.New(ErrScriptNotFound)
	}

	if len(opts.Env) > 0 {
		s.Env = append(s.Env, opts.Env...)
	}
	if len(opts.Args) > 0 {
		s.Args = append(s.Args, opts.Args...)
	}

	return workspace.RunScript(*s, opts.Quiet)
}
