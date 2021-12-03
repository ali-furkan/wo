package cmdutil

import (
	"fmt"

	"github.com/ali-furkan/wo/internal/config"
)

func GetWorkspacesFromConfig(c config.Config) map[string]map[string]string {
	spaces := c.Get("spaces").(map[string]interface{})

	workspaces := make(map[string]map[string]string)

	for sid := range spaces {
		spaceWorkspaces, ok := spaces["workspaces"].(map[interface{}]interface{})
		if !ok {
			continue
		}
		for wid, ws := range spaceWorkspaces {
			field := fmt.Sprintf("%s:%s", sid, wid)
			workspaces[field] = ws.(map[string]string)
		}
	}

	return workspaces
}
