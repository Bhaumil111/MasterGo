package main

import (
	"fmt"
	"time"
)


func greeting(){
	fmt.Println("Hello from goroutine")
}


func numbers(){
	for i:=0;i<=5;i++{
		time.Sleep(time.Millisecond*200)
		fmt.Printf("%v ", i)
	}
}
func characters(){
	for i:='a';i<='e';i++{
		time.Sleep(time.Millisecond*500)
		fmt.Printf("%c ",i)
	}
}




func main(){

	// go greeting()	
	// time.Sleep(time.Second *1)
	// fmt.Println("Hello from main routine");

	go numbers()
	go characters()

	time.Sleep(time.Second*5)
	fmt.Println("Hello from main goroutine")

}