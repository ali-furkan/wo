package cycle

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/ali-furkan/wo/internal/config"
	"github.com/ali-furkan/wo/internal/update"
	"github.com/ali-furkan/wo/internal/version"
	"github.com/fatih/color"
)

func NewNodeUpdate() *CycleNode {
	cn := NewCycleNode()

	cn.Type = OnCycleStart
	cn.AddExe(checkUpdateCycle)

	return cn
}

func checkUpdatefmt(oldVersion, newVersion, releaseURL string) string {
	str := heredoc.Docf(`
	Update Available: %s -> %s
	Run "wo update" to upgrade
	For more information: %s
`, oldVersion, newVersion, releaseURL)
	return color.YellowString(str)
}

func checkUpdateCycle(cfg *config.Config) error {
	releaseInfo, err := update.CheckForUpdate()
	if err != nil {
		return err
	}

	if releaseInfo != nil {
		fmt.Println(checkUpdatefmt(version.GetVersion(), releaseInfo.Version, releaseInfo.InfoURL))
	}

	return nil
}
