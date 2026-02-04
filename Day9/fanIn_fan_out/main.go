package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker = Fan-Out unit
// Takes jobs from 'jobs' channel
// Sends results to 'results' channel
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs { // worker keeps taking jobs until channel closed
		fmt.Println("Worker", id, "picked job", job)

		time.Sleep(1 * time.Second) // simulate work

		output := job * 10
		fmt.Println("Worker", id, "finished job", job, "-> result", output)

		results <- output // Fan-In happens here
	}

	fmt.Println("Worker", id, "exiting")
}

func main() {
	jobs := make(chan int)    // job queue
	results := make(chan int) // result collector
	var wg sync.WaitGroup

	// ---------------- FAN-OUT ----------------
	// Start 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send 6 jobs
	go func() {
		for j := 1; j <= 6; j++ {
			fmt.Println("Main sending job", j)
			jobs <- j // may block if all workers busy
		}
		close(jobs) // no more jobs
	}()

	// ---------------- FAN-IN ----------------
	// Close results channel AFTER all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results from one place
	for r := range results {
		fmt.Println("Main received result:", r)
	}

	fmt.Println("All work completed")
}




