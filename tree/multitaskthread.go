package tree

import (
	"fmt"
	"math"
	"sync"
)

type TaskManager struct {
	//tasks  chan *Task
	tasks  blockingChannel
	result chan *Node
	stop   chan struct{}
}

type blockingChannel struct {
	mu sync.Mutex
	ch chan *Task
}

type Task struct {
	n *Node
}

func (t *TaskManager) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task := <-t.tasks.ch:
			task.n.multiTaskThreadSearchUtil(t)
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

func (n *Node) MultiTaskThreadSearch(workers int, thread int) *Node {
	taskManager := TaskManager{
		tasks:  blockingChannel{ch: make(chan *Task, int(math.Pow(2, float64(thread))))},
		result: make(chan *Node, 1),
		stop:   make(chan struct{}, workers),
	}
	taskManager.tasks.ch <- &Task{
		n,
	}

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
		//m.tasks <- &Task{n.Left}
		addTask(&m.tasks, &Task{n.Left})
	}
	if n.Right != nil {
		addTask(&m.tasks, &Task{n.Right})
		//m.tasks <- &Task{n.Right}
	}
}

func addTask(c *blockingChannel, t *Task) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len((*c).ch) == cap((*c).ch) {
		newChan := make(chan *Task, cap((*c).ch)*2)
		fmt.Println("UPSCALING TO", cap((*c).ch)*2)
		for i := 0; i < cap((*c).ch); i++ {
			newChan <- <-(*c).ch
		}
		(*c).ch = newChan
		return
	}
	(*c).ch <- t
}
