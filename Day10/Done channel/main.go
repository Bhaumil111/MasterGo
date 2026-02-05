package main

import (
	"fmt"
	"time"
)


func worker(done <- chan struct{}){
	for {
		select{
		case<-done:
			fmt.Println("Worker Stopped")
			return 
		default:
			fmt.Println("Working")
			time.Sleep(1*time.Second)
		}

	}


}


func main(){

	done  := make(chan struct{})
	go worker(done)
	// time.Sleep(3.*time.Second)
	close(done)
	// var done chan struct{}


	// <-done

	fmt.Println("rest of program hello")
}


