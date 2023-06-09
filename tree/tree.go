package tree

import (
	"math/rand"
)

var r = rand.New(rand.NewSource(0))

type Node struct {
	IsNeeded bool
	Visited  bool
	Left     *Node
	Right    *Node
}

func GenerateTree(depth int) Node {
	MainNode := Node{IsNeeded: false}
	MainNode.setupChildren(depth)
	return MainNode
}

func (n *Node) RemoveVisitors() {
	n.Visited = false
	if n.Left != nil {
		n.Left.RemoveVisitors()
	}
	if n.Right != nil {
		n.Right.RemoveVisitors()
	}
}

func (n *Node) setupChildren(depthLeft int) {
	if depthLeft <= 0 {
		return
	}

	n.Left = &Node{}
	n.Right = &Node{}

	n.Left.setupChildren(depthLeft - 1)
	n.Right.setupChildren(depthLeft - 1)
}

func (n *Node) GenerateSearchPlace(depthLeft int) {
	if depthLeft <= 0 {
		n.IsNeeded = true
		return
	}
	move := r.Intn(2)
	if move == 1 {
		n.Right.GenerateSearchPlace(depthLeft - 1)
	} else {
		n.Right.GenerateSearchPlace(depthLeft - 1)
	}
}
