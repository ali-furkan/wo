package cycle

import (
	"time"

	"github.com/ali-furkan/wo/internal/cmdutil"
	"github.com/ali-furkan/wo/internal/editor"
	"github.com/ali-furkan/wo/pkg/cycle"
)

const CheckFreqTime = 5 * 60

func NewNodeEditor() *cycle.CycleNode {
	cn := cycle.NewCycleNode()
	cn.Name = "editor"
	cn.Type = OnCycleStart

	cn.AddExe(scanEditor)

	return cn
}

func scanEditor(ctx *cmdutil.CmdContext) error {
	c, err := ctx.Config()
	if err != nil {
		return err
	}

	t, ok := c.Get("last_scan_editor").(int64)
	if !ok {
		t = 0
	}
	if time.Now().Unix() < t+CheckFreqTime {
		return nil
	}

	ne, err := editor.Scan()
	if err != nil {
		return err
	}

	editors := c.Get("editors").(map[interface{}]interface{})
	if len(ne) != len(editors) {
		res := make(map[string]map[string]string)
		for _, editor := range ne {
			res[editor.Name] = editor.Map()
		}
		err = c.Set("editors", res)
		if err != nil {
			return err
		}
		err = c.Set("last_scan_editor", time.Now().Unix())
		if err != nil {
			return err
		}
	}

	if defEditor := c.GetString("defaults.editor"); defEditor == "" {
		return c.Set("defaults.editor", ne[0].ID)
	}

	return nil
}
