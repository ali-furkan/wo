package cycle

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/internal/update"
	"github.com/ali-furkan/wo/internal/version"
	"github.com/ali-furkan/wo/pkg/cycle"
	"github.com/fatih/color"
)

func NewNodeUpdate() *cycle.CycleNode {
	cn := cycle.NewCycleNode()
	cn.Name = "update"
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

func checkUpdateCycle(ctx *cmdutil.CmdContext) error {
	releaseInfo, err := update.CheckForUpdate()
	if err != nil && err.Error() == update.ErrInternetConn {
		return nil
	}

	if err != nil {
		return err
	}

	if releaseInfo != nil {
		fmt.Println(checkUpdatefmt(version.CurVersion.String(), releaseInfo.Version, releaseInfo.InfoURL))
	}

	return nil
}
