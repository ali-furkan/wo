package run

import (
	"errors"
	"fmt"

	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/internal/space"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "run [script]"
	CmdShortDesc = "Run scripts of work"
	CmdLongDesc  = "Run the specified script in your work"

	// Errors
	ErrResourceFileNotFound = "resource file not found"
	ErrWsActionNotFound     = "workspace '%s' action not found"
)

type RunOpts struct {
	Ctx *cmdutil.CmdContext

	Name       string
	WorkingDir string
	Watch      bool
	Quiet      bool
	Args       []string
	Env        []string
}

func NewCmdRun(ctx *cmdutil.CmdContext) *cobra.Command {
	opts := &RunOpts{
		Ctx: ctx,
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
	wsRC, err := opts.Ctx.WorkspaceRC()
	if err != nil {
		return errors.New(ErrResourceFileNotFound)
	}

	actionMap := wsRC.Get("run").(map[string]interface{})

	if opts.Name != "" {
		actions := wsRC.Get("actions").([]map[string]interface{})
		for _, a := range actions {
			if a["name"].(string) == opts.Name {
				actionMap = a
				break
			}
		}
	}

	if actionMap["name"] == "" || actionMap["run"] == "" {
		return fmt.Errorf(ErrWsActionNotFound, opts.Name)
	}

	action := new(space.Action)
	action.Args = actionMap["args"].([]string)
	action.Env = actionMap["env"].([]string)
	action.Name = actionMap["name"].(string)
	action.Run = actionMap["run"].(string)
	action.Workingdir = actionMap["working_dir"].(string)

	if len(opts.Env) > 0 {
		action.Env = append(action.Env, opts.Env...)
	}
	if len(opts.Args) > 0 {
		action.Args = append(action.Args, opts.Args...)
	}

	return space.RunAction(*action, opts.Quiet)
}
