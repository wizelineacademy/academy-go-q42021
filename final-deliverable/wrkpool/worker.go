package wrkpool

import (
	task "final-deliverable/task"
	"fmt"
	"sync"
)

// Worker handles all the work
type Worker struct {
	ID       int
	taskChan chan *Task
	quit     chan bool
}

// NewWorker returns new instance of worker
func NewWorker(channel chan *Task, ID int) *Worker {
	return &Worker{
		ID:       ID,
		taskChan: channel,
		quit:     make(chan bool),
	}
}

// start the worker
func (wr *Worker) Start(wg *sync.WaitGroup) {
	fmt.Printf("Starting worker %d\n", wr.ID)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for task := range wr.taskChan {
			process(wr.ID, task)
		}
	}()
}
