package auth

import (
	"context"

	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/internal/config"
	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

func GetInfo(provider string, ctx *cmdutil.CmdContext) (*User, error) {
	c, err := ctx.Config()
	if err != nil {
		return nil, err
	}

	if provider == "github" {
		return getUserDetailsFromGitHub(c)
	} else if provider == "gitlab" {
		return getUserDetailsFromGitLab(c)
	}

	return nil, nil
}

func getUserDetailsFromGitHub(c config.Config) (*User, error) {
	token := c.GetString("auth.github.token")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	userResp, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:  *userResp.Name,
		URL:   *userResp.URL,
		Email: *userResp.Email,
	}

	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		return nil, err
	}

	for _, repo := range repos {
		user.Repos = append(user.Repos, &Repository{
			Name:      *repo.Name,
			URL:       *repo.URL,
			CreatedAt: repo.CreatedAt.String(),
		})
	}

	return user, nil
}

func getUserDetailsFromGitLab(c config.Config) (*User, error) {
	// username, password := c.GetString("auth.gitlab.username"), c.GetString("auth.gitlab.password")

	return &User{}, nil
}
