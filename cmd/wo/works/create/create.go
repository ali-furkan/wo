package create

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ali-furkan/wo/internal/config"
	"github.com/ali-furkan/wo/internal/workspace"
	"github.com/google/uuid"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "create <path>"
	CmdShortDesc = "Create a new works"
	CmdLongDesc  = "Create a new works"
)

type CreateOpts struct {
	Config *config.Config

	Name              string
	Path              string
	Description       string
	Git               bool
	Readme            bool
	RCFile            bool
	Internal          bool
	Temporary         bool
	ConfirmSubmit     bool
	Template          string
	GitIgnoreTemplate string
	LicenseTemplate   string
}

func NewCmdCreate(cfg *config.Config) *cobra.Command {
	opts := &CreateOpts{
		Config: cfg,
	}

	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := homedir.Expand(args[0])
			if err != nil {
				return err
			}

			if opts.Name == "" {
				opts.Name = filepath.Base(path)
			}

			if !filepath.IsAbs(path) && (strings.HasPrefix(path, ".") || cfg.Config().Workspace.WorkDir == "") {
				d, err := os.Getwd()
				if err != nil {
					return err
				}

				path = filepath.Join(d, path)
			} else if (!filepath.IsAbs(path) && cfg.Config().Workspace.WorkDir != "") || opts.Internal {
				path = filepath.Join(cfg.Config().Workspace.WorkDir, path)
			}

			opts.Path = filepath.Clean(path)
			if !strings.HasSuffix(opts.Path, "/") {
				opts.Path += "/"
			}

			return createWork(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.ConfirmSubmit, "confirm", "y", false, "Skip the confirmation prompt")
	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Name of the work")
	cmd.Flags().StringVar(&opts.Description, "description", "", "Description of the work")
	cmd.Flags().BoolVar(&opts.Git, "git", cfg.Config().Workspace.DefaultGit, "Init git at work")
	cmd.Flags().BoolVar(&opts.Readme, "readme", cfg.Config().Workspace.DefaultReadme, "Create readme at folder of work")
	cmd.Flags().BoolVar(&opts.Internal, "internal", false, "Create work to workspace dir")
	cmd.Flags().BoolVar(&opts.Temporary, "temporary", false, "Create work to temporary dir")
	cmd.Flags().BoolVar(&opts.RCFile, "rc", cfg.Config().Workspace.DefaultRc, "Create resource file for your work")
	cmd.Flags().StringVar(&opts.Template, "template", "", "Install work with template")
	cmd.Flags().StringVarP(&opts.LicenseTemplate, "license", "l", "", "Specify SPDX License for work (SPDX ID)")

	return cmd
}

func createWork(opts *CreateOpts) error {
	for _, w := range opts.Config.Config().Workspace.Works {
		if w.Name == opts.Name {
			errTxt := fmt.Sprintf("%s work already exists", w.Name)
			return errors.New(errTxt)
		}
		if w.Path == opts.Path {
			errTxt := fmt.Sprintf("%s path already registered", w.Path)
			return errors.New(errTxt)
		}
	}

	t := time.Now()
	id := uuid.NewString()

	workType := workspace.Created
	if opts.Template != "" {
		workType = workspace.Template
	}
	if opts.Temporary {
		workType = workspace.Temporary
	}

	work := workspace.Work{
		ID:          id,
		Name:        opts.Name,
		Description: opts.Description,
		Path:        opts.Path,
		License:     opts.LicenseTemplate,
		Type:        workType,
		InitGit:     opts.Git,
		InitReadme:  opts.Readme,
		CreatedAt:   t,
		UpdatedAt:   t,
	}

	err := workspace.CreateWork(work)
	if err != nil {
		return err
	}
	fmt.Println(opts.RCFile)
	if opts.RCFile {
		err := config.CreateDefaultRCFile(opts.Path)
		if err != nil {
			return err
		}
	}

	workspace.PrintTinyStat(work)

	opts.Config.Config().Workspace.Works = append(opts.Config.Config().Workspace.Works, work)

	return nil
}
