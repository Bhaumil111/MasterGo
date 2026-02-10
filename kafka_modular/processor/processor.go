package processor

import (
	"context"
	"time"

	"example.com/kafka_clone/inMemory"
	"example.com/kafka_clone/jobs"
)

type Processor struct {
	db *inMemory.InMem
}


//  creates a new processor instance with reference to inmemory
func New(db *inMemory.InMem) *Processor {
	return &Processor{db: db}
}

// func (p *Processor) Process(ctx context.Context, job jobs.Job) error {
// 	// simulate work
// 	select {
// 	case <-time.After(4 * time.Second):
// 	case <-ctx.Done():
// 		return ctx.Err()
// 	}

// 	p.db.Save(job.ID, job.Data)
// 	return nil
// }


func (p * Processor) Process(ctx context.Context , job jobs.Job) error{
	// select {
	// case <-time.After(4*time.Second): 

	// case <-ctx.Done(): // this case is for graceful shutdown, means stop ongoing work if shutdown signal is received

	// 	return ctx.Err()
		
		
	// }

	// Simulate work with a sleep
	time.Sleep(4*time.Second) // After processing, save the result to the in-memory database

	p.db.Save(job.ID, job.Data)
	return nil
}