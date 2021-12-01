package wo

import (
	"fmt"

	"github.com/ali-furkan/wo/cmd/wo/root"
	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/internal/cycle"
	cycle_pkg "github.com/ali-furkan/wo/pkg/cycle"
)

func Run() int {
	ctx := cmdutil.NewCmdContext()

	rootCycle := cycle.NewCycleRoot(ctx)

	rootCycle.Run(cycle_pkg.OnCycleStart)

	rootCmd := root.NewCmdRoot(ctx)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	rootCycle.Run(cycle_pkg.OnCycleShutdown)

	return 0
}
