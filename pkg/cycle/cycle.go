package cycle

import (
	"fmt"
	"sync"

	"github.com/ali-furkan/wo/internal/cmdutil"
)

type Cycle struct {
	mux sync.Mutex

	ctx   *cmdutil.CmdContext
	nodes []*CycleNode
}

func NewCycle(ctx *cmdutil.CmdContext) *Cycle {
	c := &Cycle{
		ctx: ctx,
	}

	return c
}

func (c *Cycle) AddNode(newNode *CycleNode) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.nodes = append(c.nodes, newNode)
}

// Run function runs specific type nodes
func (c *Cycle) Run(t CycleNodeType) {
	for _, node := range c.nodes {
		if node.Type != t {
			continue
		}
		err := node.Run(c.ctx)
		if err != nil {
			fmt.Println("Error:\n", err)
		}
	}
}
