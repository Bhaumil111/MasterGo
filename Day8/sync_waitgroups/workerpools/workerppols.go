package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	for job := range jobs {
		fmt.Println("Worker", id, "picked job", job)

		time.Sleep(1 * time.Second) // simulate work

		fmt.Println("Worker", id, "finished job", job)
		wg.Done()
	}

	fmt.Println("Worker", id, "exiting")
}

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	// Start 2 workers
	for w := 1; w <= 2; w++ {
		go worker(w, jobs, &wg)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		wg.Add(1)
		jobs <- j
	}

	close(jobs)

	fmt.Println("Main waiting...")
	wg.Wait()
	fmt.Println("All jobs done")
}
