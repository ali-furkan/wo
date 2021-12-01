package cmd_init

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/internal/space"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "init [space:workspace|workspace]"
	CmdShortDesc = "Init a new works"
	CmdLongDesc  = "Init a new works"
)

type InitOpts struct {
	Ctx *cmdutil.CmdContext

	Space         string
	Name          string
	ID            string
	Description   string
	Path          string
	ConfirmSubmit bool
}

func NewCmdInit(ctx *cmdutil.CmdContext) *cobra.Command {
	opts := &InitOpts{
		Ctx: ctx,
	}

	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			if opts.Path == "" {
				opts.Path = wd
			} else if !filepath.IsAbs(opts.Path) {
				path := filepath.Join(wd, filepath.Dir(opts.Path))
				opts.Path = filepath.Clean(path)
			}

			if len(args) == 1 {
				a := strings.Split(args[0], ":")
				if len(a) >= 2 {
					opts.Space = a[0]
					opts.ID = strings.Join(a[1:], ":")
				} else {
					opts.Space = "global"
					opts.ID = args[0]
				}
			} else {
				opts.Name = filepath.Base(opts.Path)
				opts.Space = "global"
			}

			return initWork(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.ConfirmSubmit, "confirm", "y", false, "Skip the confirmation prompt")
	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Name of the work")
	cmd.Flags().StringVar(&opts.Description, "description", "", "Description of the work")

	return cmd
}

func initWork(opts *InitOpts) error {
	c, err := opts.Ctx.Config()
	if err != nil {
		return err
	}

	for id, w := range opts.Ctx.Workspaces() {
		if id == opts.Name {
			return fmt.Errorf("%s work already exists", id)
		}
		if w["path"] == opts.Path {
			return fmt.Errorf("%s path already registered", w["path"])
		}
	}

	t := time.Now()

	ws := space.Workspace{
		ID:          uuid.NewString(),
		Name:        opts.Name,
		Description: opts.Description,
		Path:        opts.Path,
		CreatedAt:   t,
	}

	err = space.InitWorkspace(ws, space.Options{})
	if err != nil {
		return err
	}

	space.PrintTinyStat(ws)

	wsField := fmt.Sprintf("spaces.%s.workspaces.%s")

	return c.Set(wsField, ws)
}
