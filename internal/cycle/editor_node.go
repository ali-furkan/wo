package cycle

import (
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

	tStr := c.GetString("last_scan_editor")
	t, err := time.Parse(time.RFC1123, tStr)
	if err == nil {
		if time.Since(t) < 1 {
			return nil
		}
	}

	ne, err := editor.Scan()
	if err != nil {
		return err
	}

	editors := c.Get("editors").(map[interface{}]interface{})
	if len(ne) != len(editors) {
		res := make(map[string]map[string]string)
		for _, editor := range ne {
			res[editor.Name] = make(map[string]string)
			res[editor.Name]["id"] = editor.Name
			res[editor.Name]["exec"] = editor.Exec
		}
		err = c.Set("editors", res)
		if err != nil {
			return err
		}
		err = c.Set("last_scan_editor", time.Now().String())
		if err != nil {
			return err
		}
	}

	return nil
}
