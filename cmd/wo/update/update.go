package update

import (
	"fmt"

	"github.com/ali-furkan/wo/internal/update"
	"github.com/ali-furkan/wo/internal/version"
	"github.com/spf13/cobra"
)

const (
	CmdUsage     = "update"
	CmdShortDesc = "Update Wo"
)

func NewCmdUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   CmdUsage,
		Short: CmdShortDesc,
		RunE:  updateRun,
	}

	return cmd
}

func updateRun(cmd *cobra.Command, args []string) error {
	err := update.Update()
	if err != nil {
		return err
	}

	fmt.Printf("wo upgraded to %s", version.GetVersion())

	return nil
}
