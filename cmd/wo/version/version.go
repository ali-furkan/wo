package version

import (
	"fmt"

	"github.com/ali-furkan/wo/internal/version"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "version"
	CmdShortDesc = "Show wo version"
)

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		RunE:  versionRun,
	}

	return cmd
}

func versionRun(cmd *cobra.Command, args []string) error {
	fmt.Println("Wo CLI", version.CurVersion.String())
	return nil
}
