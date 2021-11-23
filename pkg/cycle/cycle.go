package cycle

import (
	"fmt"
	"sync"

	"github.com/ali-furkan/wo/internal/config"
)

type Cycle struct {
	mux    sync.Mutex
	config *config.Config

	nodes []*CycleNode
}

func NewCycle(cfg *config.Config) *Cycle {
	c := &Cycle{
		config: cfg,
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
		err := node.Run(c.config)
		if err != nil {
			fmt.Println("Error:\n", err)
		}
	}
}
