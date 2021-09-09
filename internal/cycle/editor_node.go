package cycle

import (
	"time"

	"github.com/ali-furkan/wo/internal/config"
	"github.com/ali-furkan/wo/internal/editor"
)

func NewNodeEditor() *CycleNode {
	cn := NewCycleNode()
	cn.Type = OnCycleStart

	cn.AddExe(scanEditor)

	return cn
}

func scanEditor(cfg *config.Config) error {
	if time.Since(cfg.Workspace().LastScanEditor) < 1 {
		return nil
	}

	ne, err := editor.Scan()
	if err != nil {
		return err
	}

	e := cfg.Editors(nil)
	if len(ne) != len(*e) {
		cfg.Editors(&ne)
		cfg.Workspace().LastScanEditor = time.Now()
	}

	return nil
}
