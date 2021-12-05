package main

import (
	"academy-go-q42021/final-deliverable/model"
	"academy-go-q42021/final-deliverable/workerpool"
	"fmt"
	"time"
)

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for j := range jobs {
// 		fmt.Println("worker", id, "started  job", j)
// 		time.Sleep(time.Second)
// 		fmt.Println("worker", id, "finished job", j)
// 		results <- j * 2
// 	}
// }

func main() {

	// Prepare the data
	var allData []model.SimpleData
	for i := 0; i < 1000; i++ {
		data := model.SimpleData{ID: i}
		allData = append(allData, data)
	}

	var allTask []*workerpool.Task
	for i := 1; i <= 100; i++ {
		task := workerpool.NewTask(
			func(data interface{}) error {
				taskID := data.(int)
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("Task %d processed\n", taskID)
				return nil
			}, i)
		allTask = append(allTask, task)
	}

	pool := workerpool.NewPool(allTask, 5)
	pool.Run()

	// const numJobs = 5
	// jobs := make(chan int, numJobs)
	// results := make(chan int, numJobs)

	// for w := 1; w <= 3; w++ {
	// 	go worker(w, jobs, results)
	// }

	// for j := 1; j <= numJobs; j++ {
	// 	jobs <- j
	// }
	// close(jobs)

	// for a := 1; a <= numJobs; a++ {
	// 	<-results
	// }
}
