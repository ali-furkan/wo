package cmd_init

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ali-furkan/wo/internal/config"
	"github.com/ali-furkan/wo/internal/workspace"
	"github.com/google/uuid"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "init [name]"
	CmdShortDesc = "Init a new works"
	CmdLongDesc  = "Init a new works"
)

type InitOpts struct {
	Config *config.Config

	Name            string
	Description     string
	Path            string
	Git             bool
	Readme          bool
	RCFile          bool
	ConfirmSubmit   bool
	LicenseTemplate string
}

func NewCmdInit(cfg *config.Config) *cobra.Command {
	opts := &InitOpts{
		Config: cfg,
	}

	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := "."
			if len(args) > 0 {
				var err error
				path, err = homedir.Expand(args[0])
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

			path = filepath.Clean(path)

			if opts.Name == "" {
				opts.Name = filepath.Base(path)
			}

			opts.Path = path

			return initWork(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.ConfirmSubmit, "confirm", "y", false, "Skip the confirmation prompt")
	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Name of the work")
	cmd.Flags().StringVar(&opts.Description, "description", "", "Description of the work")
	cmd.Flags().BoolVar(&opts.Git, "git", cfg.Config().Workspace.DefaultGit, "Init git at work")
	cmd.Flags().BoolVar(&opts.RCFile, "rc", cfg.Config().Workspace.DefaultRc, "Create resource file for your work")
	cmd.Flags().BoolVar(&opts.Readme, "readme", cfg.Config().Workspace.DefaultReadme, "Create readme at folder of work")
	cmd.Flags().StringVarP(&opts.LicenseTemplate, "license", "l", "", "Specify SPDX License for work (SPDX ID)")

	return cmd
}

func initWork(opts *InitOpts) error {
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

	work := workspace.Work{
		ID:          uuid.NewString(),
		Name:        opts.Name,
		Description: opts.Description,
		Path:        opts.Path,
		License:     opts.LicenseTemplate,
		Type:        workspace.Init,
		InitGit:     opts.Git,
		InitReadme:  opts.Readme,
		CreatedAt:   t,
		UpdatedAt:   t,
	}

	err := workspace.InitWork(work)
	if err != nil {
		return err
	}

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
