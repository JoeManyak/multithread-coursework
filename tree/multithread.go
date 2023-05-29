package tree

import "sync"

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
		wg.Done()
		<-controller
		//clearIfCan(controller)
	}()

	if *stopper {
		for i := 0; i < len(controller); i++ {
			<-controller
		}
		return
	}

	n.Visited = true
	if n.IsNeeded {
		*stopper = true
		ch <- n
		return
	}
	if n.Left != nil {
		if !n.Left.Visited {
			wg.Add(1)
			controller <- struct{}{}
			go n.Left.multiThreadSearchUtil(wg, ch, stopper, controller)
		}
	}
	if n.Right != nil {
		if !n.Right.Visited {
			wg.Add(1)
			go n.Right.multiThreadSearchUtil(wg, ch, stopper, controller)
		}
	}
}

func clearIfCan(c chan struct{}) {
	select {
	case <-c:
		return
	default:
		return
	}
}
