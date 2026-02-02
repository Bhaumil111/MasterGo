package main


import "fmt"

func main(){
	ch:= make(chan int , 2)


	go func (){
		fmt.Println("G1 sending 1 ")
		ch<-1
		fmt.Println("G1 sent 1 ")

		fmt.Println("G2 sending 2")
		ch<-2
		fmt.Println("G2 sent 2")

		fmt.Println("G3 sending 3")
		ch<-3
		fmt.Println("G3 sent 3 ")
	}()
	go func() {
		fmt.Println("G2: receiving...")
		x := <-ch
		fmt.Println("G2: received", x)

		fmt.Println("G2: receiving...")
		x = <-ch
		fmt.Println("G2: received", x)

		fmt.Println("G2: receiving...")
		x = <-ch
		fmt.Println("G2: received", x)
	}()
	select{}


	

}