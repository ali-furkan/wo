package root

import (
	cmdAuth "github.com/ali-furkan/wo/cmd/wo/auth"
	cmdEditor "github.com/ali-furkan/wo/cmd/wo/editor"
	cmdUpdate "github.com/ali-furkan/wo/cmd/wo/update"
	cmdVersion "github.com/ali-furkan/wo/cmd/wo/version"
	cmdWorks "github.com/ali-furkan/wo/cmd/wo/works"
	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/internal/version"
	"github.com/spf13/cobra"
)

func NewCmdRoot(ctx *cmdutil.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		Long:  CmdLongDesc,

		SilenceErrors: true,
		SilenceUsage:  true,
		Example:       CmdExample,
		Version:       version.CurVersion.String(),
	}

	cmd.PersistentFlags().Bool("help", false, "Show help for command")
	cmd.PersistentFlags().Bool("version, v", false, "Show wo version")

	cmd.AddCommand(cmdUpdate.NewCmdUpdate())
	cmd.AddCommand(cmdVersion.NewCmdVersion())
	cmd.AddCommand(cmdEditor.NewCmdEditor(ctx))
	cmd.AddCommand(cmdAuth.NewCmdAuth(ctx))
	cmdWorks.InitCmdWorks(cmd, ctx)

	return cmd
}
