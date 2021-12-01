package gitlab

import (
	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "gitlab <username> <password>"
	CmdShortDesc = "GitLab provider"
	CmdLongDesc  = "Connect your GitLab account"
)

type AuthGitLabOpts struct {
	Ctx *cmdutil.CmdContext

	username string
	password string
}

func NewCmdAuthGitLab(ctx *cmdutil.CmdContext) *cobra.Command {
	opts := &AuthGitLabOpts{
		Ctx: ctx,
	}

	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			username, password := args[0], args[1]
			opts.username = username
			opts.password = password

			return authGitLab(opts)
		},
	}

	return cmd
}

func authGitLab(opts *AuthGitLabOpts) error {
	c, err := opts.Ctx.Config()
	if err != nil {
		return err
	}

	err = c.Set("auth.gitlab.username", opts.username)
	if err != nil {
		return err
	}
	err = c.Set("auth.gitlab.password", opts.password)
	if err != nil {
		return err
	}

	return nil
}
