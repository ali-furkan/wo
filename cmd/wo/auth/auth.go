package auth

import (
	"github.com/ali-furkan/wo/cmd/wo/auth/github"
	"github.com/ali-furkan/wo/cmd/wo/auth/gitlab"
	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "auth"
	CmdShortDesc = "Authorize"
	CmdLongDesc  = "Authorize your git service provider"
)

func NewCmdAuth(ctx *cmdutil.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
	}

	cmd.AddCommand(github.NewCmdAuthGitHub(ctx))
	cmd.AddCommand(gitlab.NewCmdAuthGitLab(ctx))

	return cmd
}
