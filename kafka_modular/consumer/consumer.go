package consumer

import (
	"context"
	"fmt"
	"sync"
	"time"
	"example.com/kafka_clone/inMemory"
	"example.com/kafka_clone/jobs"
)
// Now this is my consumer Worker function which will consume jobs from the channel and process them
func Worker(ctx context.Context, id int, jobChan <-chan jobs.Job, wg *sync.WaitGroup, db *inMemory.InMem) {
	defer wg.Done() // signal that this worker is done when the function returns

	for {
		select {
		case <-ctx.Done(): // listen for shutdown signal
			fmt.Printf("Worker %v shutdown\n", id)
			return

		case job, ok := <-jobChan: // receive job from channel
			if !ok {
				fmt.Printf("Worker %v: job channel closed\n", id)
				return
			}

			fmt.Printf("Worker %v started job %v\n", id, job.ID)
			time.Sleep(4*time.Second)
			db.Save(job.ID, job.Data) //save job 
			fmt.Printf("Worker %v finished job %v\n", id, job.ID)
		}
	}
}
