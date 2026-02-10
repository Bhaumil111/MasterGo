package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan Job, wg *sync.WaitGroup, db *InMem) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %v shutdown\n", id)
			return

		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %v: job channel closed\n", id)
				return
			}

			fmt.Printf("Worker %v started job %v\n", id, job.ID)
			time.Sleep(time.Second)
			db.Save(job.ID, job.Data)
			fmt.Printf("Worker %v finished job %v\n", id, job.ID)
		}
	}
}
