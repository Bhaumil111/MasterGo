package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"example.com/kafka_clone/consumer"
	"example.com/kafka_clone/inMemory"
	"example.com/kafka_clone/jobs"
	"example.com/kafka_clone/producer"
)
// main function to start producer and consumer workers
func main() {
	var wg sync.WaitGroup

	jobChan := make(chan jobs.Job ,3) // channel for jobs with buffer size 3

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := inMemory.NewInMem() // create in-memory db

	// start workers
	numWorkers := 3 // number of consumer workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		fmt.Printf("Worker %v on duty\n", i)
		go consumer.Worker(ctx, i, jobChan, &wg, db)
	}

	go producer.Producer(ctx, jobChan) //start producer in background

	sigCh := make(chan os.Signal, 1) // for graceful shutdown 
	signal.Notify(sigCh, syscall.SIGINT) // ctrl + c to interrupt

	<-sigCh // wait for interrupt signal
	cancel()

	fmt.Println("System Interrupted, waiting for workers")
	wg.Wait()

	fmt.Println("All workers done, exiting")
	time.Sleep(time.Second)

	fmt.Println("Final state of in-memory DB:")
	db.PrintAll() // print all data in in-memory db
}
