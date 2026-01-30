package main

import (
	"fmt"
	"time"
)


func printOne(num int ){
	fmt.Println(num)
}




func main(){


	go printOne(1) 
	go printOne(2)
	go printOne(3)
	time.Sleep(time.Millisecond *1 )
	fmt.Println("Learning go routines")




	



}