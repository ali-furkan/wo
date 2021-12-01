package cycle

import (
	"sync"

	"github.com/ali-furkan/wo/internal/cmdutil"
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

	exes []func(ctx *cmdutil.CmdContext) error
}

func NewCycleNode() *CycleNode {
	id := uuid.New()

	cn := &CycleNode{
		id: id.String(),
	}

	return cn
}

func (cn *CycleNode) AddExe(exe func(ctx *cmdutil.CmdContext) error) {
	cn.mux.Lock()
	defer cn.mux.Unlock()

	cn.exes = append(cn.exes, exe)
}

func (cn *CycleNode) Run(ctx *cmdutil.CmdContext) (err error) {
	for _, exe := range cn.exes {
		cn.mux.Lock()
		err = exe(ctx)
		cn.mux.Unlock()

		if err != nil {
			break
		}
	}

	return
}
