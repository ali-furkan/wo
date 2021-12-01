package github

import (
	"net/http"
	"time"

	"github.com/ali-furkan/wo/internal/auth"
	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "github"
	CmdShortDesc = "GitHub provider"
	CmdLongDesc  = "Connect your GitHub account"
)

type AuthGitHubOpts struct {
	Ctx *cmdutil.CmdContext

	scopes     []string
	httpClient *http.Client
}

func NewCmdAuthGitHub(ctx *cmdutil.CmdContext) *cobra.Command {
	opts := &AuthGitHubOpts{
		Ctx: ctx,
	}

	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.scopes = []string{"repo", "read:user"}
			opts.httpClient = &http.Client{Timeout: 2 * time.Second}

			return authGitHub(opts)
		},
	}

	return cmd
}

func authGitHub(opts *AuthGitHubOpts) error {
	c, err := opts.Ctx.Config()
	if err != nil {
		return err
	}

	github := &auth.GitHub{
		HttpClient: opts.httpClient,
		Scopes:     opts.scopes,
	}

	accessToken, err := github.RequestCodeAndPollToken()

	if err != nil {
		return err
	}

	return c.Set("auth.github.token", accessToken.Token)
}
