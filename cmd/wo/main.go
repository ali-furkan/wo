package wo

import (
	"fmt"

	"github.com/ali-furkan/wo/cmd/wo/root"
	"github.com/ali-furkan/wo/internal/config"
	"github.com/ali-furkan/wo/internal/cycle"
)

func Run() int {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	rootCycle := cycle.NewCycleRoot(cfg)

	rootCycle.Run(cycle.OnCycleStart)

	rootCmd := root.NewCmdRoot(cfg)
	err = rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	rootCycle.Run(cycle.OnCycleShutdown)

	return 0
}
