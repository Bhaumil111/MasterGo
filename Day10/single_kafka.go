///////////////////////////// first workerpool + graceful shutdown ////////

// package main

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"os/signal"
// 	"sync"
// 	"syscall"
// 	"time"
// )

// // worker functions to process jobs and listen for shutdown signal from main function
// func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for { // infinite loop to keep worker running until shutdown signal is received

// 		select {
// 		case <-ctx.Done(): // listen for shutdown signal from main function
// 			fmt.Printf("Worker %v Shutdown\n", id)

// 			return
// 		case job := <-jobs: // listen for jobs from the jobs channel

// 			fmt.Printf("Worker %v start processing Job %v\n ", id, job)
// 			time.Sleep(3000 * time.Millisecond)
// 			fmt.Printf("Worker %v finished job %v\n", id, job)

// 		}

// 	}
// }

// func main() {

// 	ctx, cancel := context.WithCancel(context.Background())
// 	var wg sync.WaitGroup

// 	// making job channel
// 	jobs := make(chan int)

// 	// now creating a sig channel to listen for interrupt signals from the operating system like ctrl +c

// 	sigCh := make( chan os.Signal , 1) // creating a channel to receive OS signals,

// 	signal.Notify( sigCh , syscall.SIGINT) // this line tells the signal package to relay on signal like ctrl +c to interrput.

// 	// number of workers
// 	numWorkers := 3 //  I have numworkers running concurrently

// 	for i := 1; i <= numWorkers; i++ {

// 		wg.Add(1)
// 		fmt.Printf("Worker %v ready on duty\n", i)
// 		go worker(ctx, i, jobs, &wg) // start worker goroutine

// 	}

// 	// adding jobs to the jobs channel
// 	go func() {

// 		// for j := 1; j <= 5; j++ {
// 		// 	fmt.Printf("Added Job %v to the channel\n", j)
// 		// 	jobs <- j // send job to the jobs channel
// 		// 	time.Sleep(5000 * time.Millisecond)

// 		// }

// 		jobID := 1
// 		for {

// 			fmt.Printf("Added Job %v to the channel\n", jobID)
// 			jobs <- jobID // send job to the jobs channel
// 			jobID++
// 			time.Sleep(5000 * time.Millisecond)

// 		}
// 	}() // this goroutine is responsible for adding jobs to the jobs channel. It runs concurrently with the worker goroutines and sends jobs to the channel every 5 seconds. The workers will pick up these jobs and process them as they become available.

// 	// fmt.Println("I am triggering cancel after 8 second...")
// 	// time.Sleep(8000 * time.Millisecond)

// 	<-sigCh // it means main function block until any ctrl + c signal is received

// 	cancel()

// 	fmt.Println("Program is shutting down... Waiting for workers to finish ongoing jobs")

// 	wg.Wait()

// 	fmt.Println("All workers have finished. Exiting program.")
// 	time.Sleep(time.Millisecond * 100)

// }

///////////////////////// Now i am implementing map (in mem) store with prod , cons model like (kafka ) with all above features//////



// prod ---- ////// ---- consumer 


package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Job struct {
	ID   int
	Data string
}

// my in memory db in which prodcuer will save data , consumer will read data
type InMem struct {
	mu   sync.Mutex
	data map[int]string
}

func (inMemDb *InMem) Save(id int, data string) {
	inMemDb.mu.Lock()         // lock
	defer inMemDb.mu.Unlock() // mutex must be unlocked

	inMemDb.data[id] = data // saving data to map
}

func worker(ctx context.Context, id int, jobs <-chan Job, wg *sync.WaitGroup, inMemDb *InMem) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done(): // listen for shutdown signal
			fmt.Printf("Worker %v shutdown\n", id)
			return
		case job := <-jobs: // listen for job from jobs channel
			fmt.Printf("Worker %v started processing job %v \n", id, job.ID) // process job
			time.Sleep(4000 * time.Millisecond)
			inMemDb.Save(job.ID, job.Data) // save job data to inMem db

			fmt.Printf("Worker %v finshed job %v\n", id, job.ID)

		}

	}

}

func main() {

	var wg sync.WaitGroup
	jobs := make(chan Job , 3)

	ctx, cancel := context.WithCancel(context.Background()) // creating root context

	// using map as reference type so that all workers can access the same map
	inMemDb := InMem{

		data: make(map[int]string),
	}

	//starting workers
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		fmt.Printf("Worker %v on duty \n", i)
		go worker(ctx, i, jobs, &wg, &inMemDb) // worker goroutine
	}

	// addding jobs to the Job channel

	go func() { // here we use goroutine to add jobs in background concurrently
		jobID := 1

		for  {
			select {

			case <-ctx.Done():
				fmt.Printf("Jobs producer shutdown\n")
				return
			default:
				fmt.Printf("Added job %v to the channel\n", jobID)
				job:= Job{

					ID:jobID,
					Data : fmt.Sprintf("Data %v", jobID),
				}

				jobs <- job // send job to jobs channel

				jobID++
				time.Sleep(2000 * time.Millisecond)



			}

		}

	}()

	//creating a sigChan for graceful shutdown

	sigCh := make(chan os.Signal, 1)     //Buffer of 1 so it blocks only when there is no space
	signal.Notify(sigCh, syscall.SIGINT) // ctrl + c to interrupt

	<-sigCh // block main until signal is received
	cancel()

	fmt.Println("System Interrupted  waiting for workers to finish ongoing stuff")
	wg.Wait()

	fmt.Println("All workers done ,exiting now")
	time.Sleep(time.Millisecond * 1000)

}
