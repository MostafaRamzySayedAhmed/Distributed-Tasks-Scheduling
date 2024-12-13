package main

import (
	"fmt"
	"sync"
	"time"
)

// Task structure to represent a task
type Task struct {
	ID int
}

func worker(id int, tasks <-chan Task, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		// Simulating task processing
		time.Sleep(2 * time.Second)
		results <- fmt.Sprintf("Worker %d completed task %d", id, task.ID)
	}
}

func main() {
	// Number of workers
	numWorkers := 3
	// Channel for tasks and results
	tasks := make(chan Task, 5)
	results := make(chan string, 5)

	// WaitGroup to wait for all workers to complete their tasks
	var wg sync.WaitGroup

	// Start the workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Add tasks to the task channel
	for i := 1; i <= 5; i++ {
		tasks <- Task{ID: i}
	}
	close(tasks)

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print the results
	for result := range results {
		fmt.Println(result)
	}
}
