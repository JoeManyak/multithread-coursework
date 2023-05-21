package tree

import (
	"math/rand"
	"sync"
)

var r = rand.New(rand.NewSource(0))

type Node struct {
	IsNeeded bool
	Left     *Node
	Right    *Node
}

func GenerateTree(depth int) Node {
	MainNode := Node{IsNeeded: false}
	MainNode.setupChildren(depth)
	MainNode.generateSearchPlace(depth / 2)
	return MainNode
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

func (n *Node) generateSearchPlace(depthLeft int) {
	if depthLeft <= 0 {
		n.IsNeeded = true
		return
	}
	move := r.Intn(2)
	if move == 1 {
		n.Left.generateSearchPlace(depthLeft - 1)
	} else {
		n.Right.generateSearchPlace(depthLeft - 1)
	}
}

func (n *Node) SingleThreadSearch() *Node {
	res := make(chan *Node, 1)
	n.singleThreadSearchUtil(res)
	select {
	case node := <-res:
		return node
	default:
		return nil
	}
}

func (n *Node) singleThreadSearchUtil(res chan *Node) {
	if n.IsNeeded {
		res <- n
		return
	}
	if n.Left != nil {
		n.Left.singleThreadSearchUtil(res)
	}
	if n.Right != nil {
		n.Right.singleThreadSearchUtil(res)
	}
}

func (n *Node) MultiThreadSearch(threads int) *Node {
	wg := sync.WaitGroup{}
	ch := make(chan *Node, 1)
	threadController := make(chan struct{}, threads)

	stopper := false

	wg.Add(1)

	threadController <- struct{}{}
	go n.multiThreadSearchUtil(&wg, ch, &stopper, threadController)

	waiter := make(chan struct{})
	go waitAsChan(waiter, &wg)

	select {
	case res := <-ch:
		return res
	case _ = <-waiter:
		return nil
	}
}

func waitAsChan(ch chan struct{}, wg *sync.WaitGroup) {
	wg.Wait()
	ch <- struct{}{}
}

func (n *Node) multiThreadSearchUtil(wg *sync.WaitGroup, ch chan *Node, stopper *bool, controller chan struct{}) {
	defer func() {
		<-controller
		wg.Done()
	}()

	if n.IsNeeded {
		*stopper = true
		ch <- n
		close(controller)
		return
	}
	if n.Left != nil {
		wg.Add(1)
		controller <- struct{}{}
		go n.Left.multiThreadSearchUtil(wg, ch, stopper, controller)
	}
	if n.Right != nil {
		wg.Add(1)
		controller <- struct{}{}
		go n.Right.multiThreadSearchUtil(wg, ch, stopper, controller)
	}
}
