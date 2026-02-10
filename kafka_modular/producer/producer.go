package producer 
import (
	"context"
	"fmt"
	"time"
	"example.com/kafka_clone/jobs"
)
// this is my producer function which will produce jobs and send them to the channel
func Producer(ctx context.Context, jobChan chan<- jobs.Job) {
	jobID := 1

	for {
		select {
		case <-ctx.Done():// listen for shutdown signal
			fmt.Println("Jobs producer shutdown")
			close(jobChan)
			return

		default: // produce jobs
			// fmt.Printf("Added job %v to the channell\n", jobID)
			fmt.Printf("Producer producing job %v \n", jobID)
			job := jobs.Job{
				ID:   jobID,
				Data: fmt.Sprintf("Data %v", jobID),
			}

			jobChan <- job

			jobID++
			time.Sleep(2* time.Second)
		}
	}
}
