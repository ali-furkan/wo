package cycle

import (
	"sync"

	"github.com/ali-furkan/wo/internal/config"
	"github.com/google/uuid"
)

type CycleNodeType uint

const (
	OnCycleStart CycleNodeType = iota
	OnCycleShutdown
)

type CycleNode struct {
	mux sync.Mutex

	id   string
	Type CycleNodeType
	Name string

	exes []func(cfg *config.Config) error
}

func NewCycleNode() *CycleNode {
	id := uuid.New()

	cn := &CycleNode{
		id: id.String(),
	}

	return cn
}

func (cn *CycleNode) AddExe(exe func(cfg *config.Config) error) {
	cn.mux.Lock()
	defer cn.mux.Unlock()

	cn.exes = append(cn.exes, exe)
}

func (cn *CycleNode) Run(cfg *config.Config) (err error) {
	for _, exe := range cn.exes {
		cn.mux.Lock()
		err = exe(cfg)
		cn.mux.Unlock()

		if err != nil {
			break
		}
	}

	return
}
