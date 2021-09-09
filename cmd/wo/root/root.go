package root

import (
	cmdEditor "github.com/ali-furkan/wo/cmd/wo/editor"
	cmdUpdate "github.com/ali-furkan/wo/cmd/wo/update"
	cmdVersion "github.com/ali-furkan/wo/cmd/wo/version"
	cmdWorks "github.com/ali-furkan/wo/cmd/wo/works"
	"github.com/ali-furkan/wo/internal/config"
	"github.com/ali-furkan/wo/internal/version"
	"github.com/spf13/cobra"
)

func NewCmdRoot(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,

		SilenceErrors: true,
		SilenceUsage:  true,
		Example:       CmdExample,
		Version:       version.GetVersion(),
	}

	cmd.PersistentFlags().Bool("help", false, "Show help for command")
	cmd.PersistentFlags().Bool("version, v", false, "Show wo version")

	cmd.AddCommand(cmdUpdate.NewCmdUpdate())
	cmd.AddCommand(cmdVersion.NewCmdVersion())
	cmd.AddCommand(cmdEditor.NewCmdEditor(cfg))
	cmdWorks.InitCmdWorks(cmd, cfg)

	return cmd
}
