// ==================================================================================
// Deadlock Case when one ch1 send data and other channel receiving data
// ==================================================================================

// package main

// import "fmt"

// func main(){
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func(){
// 		ch1<-10
// 		fmt.Println("Sending value " , 10)

// 	}()
// 	x:=<-ch2
// 	fmt.Println("Received value for the ch1" ,x)

// }

// ==================================================================================
// Deadlock Case Resolve with the help of the select statement
// ==================================================================================

// package main

// import "fmt"

// func main(){
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func(){
// 		ch2<-20
// 		ch1<-10
// 		fmt.Println("Sending value for ch1 " , 10)
// 		fmt.Println("Sending value for ch2 ", 20 )

// 	}()

// 	select{
// 	case x := <-ch1:
// 		fmt.Println("Received value for ch1" , x)

// 	case y := <-ch2:
// 		fmt.Println("Received value for ch2", y)

// 	}

// }

// ==================================================================================
// Deadlock Case Resolve with the help of the select statement
// ==================================================================================

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func worker ( name string , ch chan string , delay time.Duration){

// 	time.Sleep(delay)
// 	ch<- name+" done "

// }

// func main(){

// 	ch1:= make(chan string)
// 	ch2:= make(chan string)
// 	go worker("Task 1" , ch1, 2 * time.Second)
// 	go worker("Task 2 " , ch2 ,1 *time.Second )

// 	for i:=0;i<5;i++{

// 		select{

// 		case msg:= <-ch1:
// 			fmt.Println("Received msg from channel 1 " ,msg)
// 		case msg:= <-ch2:
// 			fmt.Println("Received msg from channel 2 " ,msg)
// 		case <-time.After(3*time.Second):
// 			fmt.Println("Timeout !")
// 		}

// 	}

// }
// ==================================================================================
// Normal Select Statement
// ==================================================================================

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func goOne(ch chan string){
// 	msg:= "From goOne"
// 	time.Sleep(1*time.Second)
// 	ch<-msg
// }
// func goTwo(ch chan string){
// 	msg:= "From goTwo"
// 	time.Sleep(1*time.Second)
// 	ch<-msg
// }

// func main(){

// 	ch1:= make(chan string)
// 	ch2:= make(chan string)

// 	go goOne(ch1)
// 	go goTwo(ch2)

// 	select{

// 	case msg := <-ch1:
// 		fmt.Println(msg)
// 	case  msg :=<-ch2:
// 		fmt.Println(msg)
// 	}

// }

// ==================================================================================
// Normal Select Statement
// ==================================================================================

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func goOne(ch chan string){
// 	msg:= "From goOne"
// 	time.Sleep(1*time.Second)
// 	ch<-msg
// }
// func goTwo(ch chan string){
// 	msg:= "From goTwo"
// 	time.Sleep(1*time.Second)
// 	ch<-msg
// }

// func main(){

// 	ch1:= make(chan string)
// 	ch2:= make(chan string)

// 	go goOne(ch1)
// 	go goTwo(ch2)

// 	for i:=0;i<2;i++{

// 	select{

// 	case msg := <-ch1:
// 		fmt.Println(msg)
// 	case  msg :=<-ch2:
// 		fmt.Println(msg)
// 	}

// 	}

// }

// ==================================================================================
// Normal Select Statement with deadlock fixed with default
// ==================================================================================

// package main

// import (
// 	"fmt"
// )

// func main(){

// 	ch1:= make(chan string)

// 	select{
// 	case msg:=<-ch1:
// 		fmt.Println(msg)
// 	default:
// 		fmt.Println("Default Statement executed")
// 	}

// }	
// ==================================================================================
// Select statement with receiver and sender
// ==================================================================================

// package main

// import (
// 	"fmt"
// )

// func goSender(ch chan string) {
// 	msg := "From goOne"
// 	ch <- msg
// }
// func goReceiver(ch chan string) {
// 	x := <-ch
// 	fmt.Println("Received data in receiver", x)
// }

// func main() {

// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go goSender(ch1)
// 	go goReceiver(ch2)

// 	for i := 0; i < 2; i++ {

// 		select {

// 		case msg := <-ch1:
// 			fmt.Println(msg)
// 		case ch2 <- "To goTwo Gr":

// 		}

// 	}

// }


//================================ AFter time


// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch1 := make(chan string)
// 	go goOne(ch1)

// 	select {
// 	case msg := <-ch1:
// 		fmt.Println(msg)
// 	case <-time.After(time.Second * 1):
// 		fmt.Println("Timeout")
// 	}
// }

// func goOne(ch chan string) {
// 	time.Sleep(time.Second * 2)
// 	ch <- "From goOne goroutine"
// }
//============================= Break
package main
import "fmt"

func main() {
	ch := make(chan string, 1)
	ch <- "Before break"

	select {
	case msg := <-ch:
		fmt.Println(msg)
		// break
		fmt.Println("After break")
	default:
		fmt.Println("Default case")
	}
	fmt.Println("Hello from outside")
}