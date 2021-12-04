package wrkpool

import(
	"fmt"
)

//Task encapsulate a work item.
type Task struct{
	Err error
	Data interface{}
	f 	func(interface{}) error
}

// NewTask initialize a new task
func (f func(interface{}), error, data interface{}) *Task {
	return &Task(f: f, Data: data)
}

// process the task
func process(workerID int, task *Task)  {
	fmt.Printf("Worker %d process task %v\n", workerID, task.Data)
	task.Err = task.f(task.Data)
}