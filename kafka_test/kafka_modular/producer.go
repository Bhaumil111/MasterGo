package main

import (
	"context"
	"fmt"
	"time"
)

func producer(ctx context.Context, jobs chan<- Job) {
	jobID := 1

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Jobs producer shutdown")
			close(jobs)
			return

		default:
			job := Job{
				ID:   jobID,
				Data: fmt.Sprintf("Data %v", jobID),
			}

			fmt.Printf("Added job %v to the channel\n", jobID)
			jobs <- job

			jobID++
			time.Sleep(5 * time.Second)
		}
	}
}
