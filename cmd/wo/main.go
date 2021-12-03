package wo

import (
	"fmt"

	"github.com/ali-furkan/wo/cmd/wo/root"
	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/internal/cycle"
)

func Run() int {
	ctx := cmdutil.NewCmdContext()

	rootCycle := cycle.NewCycleRoot(ctx)

	rootCycle.Run(cycle.OnCycleStart)

	rootCmd := root.NewCmdRoot(ctx)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	rootCycle.Run(cycle.OnCycleShutdown)

	return 0
}
