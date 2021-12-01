package cmdutil

import (
	"fmt"

	"github.com/ali-furkan/wo/internal/config"
)

func GetWorkspacesFromConfig(c config.Config) map[string]map[string]string {
	spaces := c.Get("spaces").(map[string]interface{})

	for id := range spaces {
		field := fmt.Sprintf("spaces.%s.workspaces", id)
		workspaces := c.Get(field).(map[string]map[string]string)

		return workspaces
	}

	return nil
}
