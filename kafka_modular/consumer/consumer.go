package consumer

import (
	"context"
	"fmt"
	"sync"

	"example.com/kafka_clone/inMemory"
	"example.com/kafka_clone/jobs"
	"example.com/kafka_clone/processor"
)

// Now this is my consumer Worker function which will consume jobs from the channel and process them
func Worker(ctx context.Context, id int, jobChan <-chan jobs.Job, wg *sync.WaitGroup, db *inMemory.InMem, p *processor.Processor) {
	defer wg.Done() // signal that this worker is done when the function returns

	for {
		select {
		case <-ctx.Done(): // listen for shutdown signal
			
			fmt.Printf("Consumer %v received shutdown signal\n", id)
			return

		case job, ok := <-jobChan: // receive job from channel
			if !ok {
				
				fmt.Printf("Consumer %v: job channel closed, exiting\n", id)
				return
			}

			// fmt.Printf("Consumer %v started job %v\n", id, job.ID)
			// time.Sleep(4*time.Second)
			// db.Save(job.ID, job.Data) //save job 
			// fmt.Printf("Consumer %v finished job %v\n", id, job.ID)



			// Process the job using the processor
			if err := p.Process(ctx , job); err != nil{
				fmt.Printf("Consumer %v: error processing job %v: %v\n", id, job.ID, err)
				continue
			}

			fmt.Printf("Consumer %v finished job %v\n", id, job.ID)


		}
	}
}
