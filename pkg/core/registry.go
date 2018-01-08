package core

import (
	"sync"
)

type Node struct{}

type NodeRegistry struct {
	m     *sync.RWMutex
	Nodes []*Node
}

func (nd *NodeRegistry) RegisterNode() error {
	return nil
}

func (nd *NodeRegistry) DeregisterNode() error {
	return nil
}
