package cycle

import (
	"time"

	"github.com/ali-furkan/wo/internal/config"
	"github.com/ali-furkan/wo/internal/editor"
	"github.com/ali-furkan/wo/pkg/cycle"
)

func NewNodeEditor() *cycle.CycleNode {
	cn := cycle.NewCycleNode()
	cn.Type = cycle.OnCycleStart

	cn.AddExe(scanEditor)

	return cn
}

func scanEditor(cfg *config.Config) error {
	if time.Since(cfg.Config().Workspace.LastScanEditor) < 1 {
		return nil
	}

	ne, err := editor.Scan()
	if err != nil {
		return err
	}

	e := &cfg.Config().Editors
	if len(ne) != len(*e) {
		*e = ne
		cfg.Config().Workspace.LastScanEditor = time.Now()
	}

	return nil
}
