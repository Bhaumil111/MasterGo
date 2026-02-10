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

func main() {
	var wg sync.WaitGroup

	jobs := make(chan Job)

	ctx, cancel := context.WithCancel(context.Background())

	db := NewInMem()

	// start workers
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		fmt.Printf("Worker %v on duty\n", i)
		go worker(ctx, i, jobs, &wg, db)
	}

	go producer(ctx, jobs) //start producer in background

	sigCh := make(chan os.Signal, 1) // for graceful shutdown 
	signal.Notify(sigCh, syscall.SIGINT) // ctrl + c to interrupt

	<-sigCh
	cancel()

	fmt.Println("System Interrupted, waiting for workers")
	wg.Wait()

	fmt.Println("All workers done, exiting")
	time.Sleep(time.Second)
}
