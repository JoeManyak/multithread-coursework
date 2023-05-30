package tree

import (
	"sync"
)

type TaskManager struct {
	tasks  chan *Node
	result chan *Node
	stop   chan struct{}
	depth  int
}

func (t *TaskManager) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task := <-t.tasks:
			task.multiTaskThreadSearchUtil(t, t.depth)
			break
		case <-t.stop:
			t.stop <- struct{}{}
			return
		default:
			t.stop <- struct{}{}
			return
		}
	}
}

func (n *Node) MultiTaskThreadSearch(workers int, taskSize int, depth int) *Node {
	taskManager := TaskManager{
		tasks:  make(chan *Node, taskSize),
		result: make(chan *Node, 1),
		stop:   make(chan struct{}, workers),
		depth:  depth,
	}
	taskManager.tasks <- n

	var wg sync.WaitGroup
	wg.Add(workers)
	go taskManager.Run(&wg)
	for i := 0; i < workers-1; i++ {
		t := <-taskManager.tasks
		go taskManager.Run(&wg)
		taskManager.tasks <- t
	}

	ch := make(chan struct{})
	go waitAsChan(ch, &wg)

	select {
	case r := <-taskManager.result:
		return r
	case <-ch:
		return nil
	}
}

func waitAsChan(ch chan struct{}, wg *sync.WaitGroup) {
	wg.Wait()
	ch <- struct{}{}
}

func (n *Node) multiTaskThreadSearchUtil(
	m *TaskManager,
	depth int,
) {
	if n.IsNeeded {
		m.result <- n
		m.stop <- struct{}{}
		return
	}

	if n.Left != nil {
		if depth == 0 {
			m.tasks <- n.Left
		} else {
			n.Left.multiTaskThreadSearchUtil(m, depth-1)
		}
	}
	if n.Right != nil {
		if depth == 0 {
			m.tasks <- n.Right
		} else {
			n.Right.multiTaskThreadSearchUtil(m, depth-1)
		}
	}
}
