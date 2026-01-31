// package main

// import (
// 	"fmt"
// 	"time"
// )


// func main(){
// ch :=make(chan int) // creating an unbuffered channel


//  go func(){
// 	time.Sleep(time.Second)
// 	ch<-5 // Sending value to the channel
//  }()


//  x:= <-ch // recieving data in channel 

//  fmt.Println(x)


// }

// ==========================================================
// package main
// import "fmt"
// func main(){
// 	var x int
// 	go func(){
// 		x = 10
// 	}()


// 	fmt.Println(x)

// }
// ==========================================================
// package main

// import "fmt"

// func main(){
// 	ch:=make(chan int)


// 	go func(){
// 		fmt.Println("Reciever waiting")
// 		val:= <-ch
// 		fmt.Printf("Recieved value %v\n", val)

// 	}()


	
// 	ch <-32 // sent 32 to the channel
	
// 	fmt.Printf("Sent %v\n" ,ch)
// }

// ==========================================================


// Sent first then received

// package main
// import "fmt"


// func main(){
	
// 	ch:=make(chan int)
// 	go func(){

// 		ch<-32
// 		fmt.Println("Sent value ",32)

// 	}()
// 	val:= <-ch
// 	fmt.Println("Reciever waiting")
// 	fmt.Println("Recieved value", val)
	

// }
// ==========================================================
// All Goroutines are asleep -deadlock!
// package main


// // import "fmt"


// func main(){
// 	ch:= make(chan int)
// 	ch<-32;
// }

// ==========================================================


// Using channels as signal
// package main


// import "fmt"
// func worker (done chan bool){
// 	fmt.Println("working")
// 	done<-true
// }



// func main(){
// 	done:= make( chan bool)
	
// 	 go worker(done)
// 	 <-done
// 	 fmt.Println("Work done")

// }

// ==========================================================
// package main


// import "fmt"



// func main(){
// 	 ch:= make(chan int)
// 	 go func(){
// 		fmt.Println("Receiver Waiting ")
// 		val:=<-ch
// 		fmt.Println("Received ", val)

// 	 }()
// 	 ch<-42 // Sent data 42

// 	 fmt.Println("Sent data" , 42)

	
// }

// ==========================================================


// package main
// import "fmt"
// func main(){

// ch := make(chan int) //create an unbuffered channels
// go func(){
// 	val :=432
// 	fmt.Println("Sending data")
// 	ch<-val
// 	fmt.Println("Data Sent Successfully")
// }()

// go func(){
// 	fmt.Println("Waiting for the data to received")
// 	val:= <-ch
// 	fmt.Println("Data received" , val)
// }()


// 	fmt.Println("Waiting receiver 2")
// 	val:=<-ch
// 	fmt.Println("Data received" , val)



// select{}

// }


// ==========================================================
// Multiple senders + muliplte receiver

package main
import "fmt"
func main(){

	ch:=make(chan int)
	// G1
	go func() {
		fmt.Println("G1: sending 1")
		ch <- 1
		fmt.Println("G1: send done")
	}()

	// G2
	go func() {
		fmt.Println("G2: sending 2")
		ch <- 2
		fmt.Println("G2: send done")
	}()

	// G3
	go func() {
		fmt.Println("G3: waiting to receive")
		x := <-ch
		fmt.Println("G3: received", x)
	}()

	// G4
	go func() {
		fmt.Println("G4: waiting to receive")
		x := <-ch
		fmt.Println("G4: received", x)
	}()

	select {}


}



