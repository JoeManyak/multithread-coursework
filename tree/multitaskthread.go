package tree

import (
	"math"
	"time"
)

type TaskManager struct {
	tasks  chan *Node
	result chan *Node
	stop   chan struct{}
}

type Task struct {
	n *Node
}

func (t *TaskManager) Run() {
	for task := range t.tasks {
		task.multiTaskThreadSearchUtil(t)
	}
	//fmt.Println("ded")
}

func (n *Node) MultiTaskThreadSearch(workers int, deep int) *Node {
	taskManager := TaskManager{
		tasks:  make(chan *Node, int(math.Pow(2, float64(deep)))),
		result: make(chan *Node, 1),
		stop:   make(chan struct{}, workers),
	}
	taskManager.tasks <- n

	for i := 0; i < workers; i++ {
		go taskManager.Run()
	}

	ch := make(chan struct{}, 2)
	go checker(taskManager.tasks, ch)

	select {
	case r := <-taskManager.result:
		return r
	case <-ch:
		ch <- struct{}{}
		return nil
	}
}

func (n *Node) multiTaskThreadSearchUtil(
	m *TaskManager,
) {
	if n.IsNeeded {
		m.result <- n

		close(m.tasks)
		return
	}
	if n.Left != nil {
		m.tasks <- n.Left
	}
	if n.Right != nil {
		m.tasks <- n.Right
	}
}

func checker(ch chan *Node, stop chan struct{}) {
	for {
		time.Sleep(time.Second)
		if len(ch) == 0 {
			select {
			case <-stop:
				return
			default:
				close(ch)
				stop <- struct{}{}
			}
		}
	}
}
