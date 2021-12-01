package create

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/internal/space"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "create <space:workspace|workspace>"
	CmdShortDesc = "Create a new works"
	CmdLongDesc  = "Create a new works"

	// Errors
	ErrInvalidWorkspaceName = "invalid workspace name"
	ErrInvalidSpaceName     = "invalid space name"
)

type CreateOpts struct {
	Ctx *cmdutil.CmdContext

	Space           string
	Name            string
	ID              string
	Path            string
	Description     string
	Git             string
	Gitignore       string
	LicenseTemplate string
	Readme          string
	External        bool
	Temporary       bool
	ConfirmSubmit   bool
}

func NewCmdCreate(ctx *cmdutil.CmdContext) *cobra.Command {
	opts := &CreateOpts{
		Ctx: ctx,
	}

	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			a := strings.Split(args[0], ":")
			if len(a) >= 2 {
				opts.Space = a[0]
				opts.ID = strings.Join(a[1:], ":")
			} else {
				opts.Space = "global"
				opts.ID = args[0]
			}

			if err := validation.Validate(opts.ID, space.WorkspaceNameValidationRules...); err != nil {
				return errors.New(ErrInvalidWorkspaceName)
			}

			c, err := ctx.Config()
			if err != nil {
				return err
			}

			spacePathField := fmt.Sprintf("spaces.%s.root_dir", opts.Space)

			path := c.GetString(spacePathField)
			if path == "" {
				return errors.New(ErrInvalidSpaceName)
			}
			if opts.External {
				path, err = homedir.Expand(opts.Path)
				if err != nil {
					return err
				}
			}

			if !filepath.IsAbs(path) {
				d, err := os.Getwd()
				if err != nil {
					return err
				}

				path = filepath.Join(d, path)
			}

			path = filepath.Join(path, opts.ID)

			opts.Path = filepath.Clean(path)
			opts.Path += "/"

			return createWork(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.ConfirmSubmit, "confirm", "y", false, "Skip the confirmation prompt")
	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Name of the work")
	cmd.Flags().StringVar(&opts.Description, "description", "", "Description of the work")
	cmd.Flags().BoolVar(&opts.External, "external", false, "Create the workspace outside space dir")
	cmd.Flags().BoolVar(&opts.Temporary, "temporary", false, "Create work to temporary dir")
	cmd.Flags().StringVarP(&opts.LicenseTemplate, "license", "l", "", "Specify SPDX License for work (SPDX ID)")

	return cmd
}

func createWork(opts *CreateOpts) error {
	c, err := opts.Ctx.Config()
	if err != nil {
		return err
	}

	for _, w := range opts.Ctx.Workspaces() {
		if w["id"] == opts.ID {
			errTxt := fmt.Sprintf("%s work already exists", opts.ID)
			return errors.New(errTxt)
		}
		if w["path"] == opts.Path {
			errTxt := fmt.Sprintf("%s path already registered", opts.Path)
			return errors.New(errTxt)
		}
	}

	t := time.Now()
	id := uuid.NewString()

	ws := space.Workspace{
		ID:          id,
		Name:        opts.Name,
		Description: opts.Description,
		Path:        opts.Path,
		CreatedAt:   t,
	}

	wsOpts := space.Options{
		Readme:    opts.Readme,
		License:   opts.LicenseTemplate,
		Git:       opts.Git,
		Gitignore: opts.Gitignore,
	}

	err = space.CreateWorkspace(ws, wsOpts)
	if err != nil {
		return err
	}

	space.PrintTinyStat(ws)

	wsField := fmt.Sprintf("spaces.%s.workspaces.%s", opts.Space, opts.ID)

	return c.Set(wsField, ws)
}
