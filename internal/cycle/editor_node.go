package cycle

import (
	"fmt"
	"time"

	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/internal/editor"
	"github.com/ali-furkan/wo/pkg/cycle"
)

func NewNodeEditor() *cycle.CycleNode {
	cn := cycle.NewCycleNode()
	cn.Type = cycle.OnCycleStart

	cn.AddExe(scanEditor)

	return cn
}

func scanEditor(ctx *cmdutil.CmdContext) error {
	c, err := ctx.Config()
	if err != nil {
		return err
	}

	t := c.Get("last_scan_editor").(time.Time)

	if time.Since(t) < 1 {
		return nil
	}

	ne, err := editor.Scan()
	if err != nil {
		return err
	}

	editors := c.Get("editors").(map[string]map[string]interface{})
	if len(ne) != len(editors) {
		res := make(map[string]map[string]string)
		for _, editor := range ne {
			res[editor.Name]["id"] = editor.Name
			res[editor.Name]["exec"] = editor.Exec
		}
		editorErr := c.Set("editors", res)
		scanEditorErr := c.Set("last_scan_editor", time.Now())
		if editorErr != nil || scanEditorErr != nil {
			return fmt.Errorf("cycle editor node err: %s, %s", editorErr.Error(), scanEditorErr.Error())
		}
	}

	return nil
}
