package tree

import (
	"sync"
)

type TaskManager struct {
	tasks  chan *Node
	result chan *Node
	stop   chan struct{}
}

func (t *TaskManager) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task := <-t.tasks:
			task.multiTaskThreadSearchUtil(t)
			break
		case <-t.stop:
			//fmt.Println("GOT STOP SIGNAL")
			t.stop <- struct{}{}
			return
		default:
			t.stop <- struct{}{}
			//fmt.Println("NO TASK, ENDING THREAD...")
			return
		}
	}
}

func (n *Node) MultiTaskThreadSearch(workers int, taskSize int) *Node {
	taskManager := TaskManager{
		tasks:  make(chan *Node, taskSize),
		result: make(chan *Node, 1),
		stop:   make(chan struct{}, workers),
	}
	taskManager.tasks <- n

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go taskManager.Run(&wg)
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

func (n *Node) multiTaskThreadSearchUtil(
	m *TaskManager,
) {
	if n.IsNeeded {
		m.result <- n
		m.stop <- struct{}{}
		return
	}
	if n.Left != nil {
		m.tasks <- n.Left
	}
	if n.Right != nil {
		m.tasks <- n.Right
	}
}
