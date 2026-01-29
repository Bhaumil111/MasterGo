package main

import "fmt"
func addOne(x interface{}){
	v, ok :=x.(int)
	if !ok {
		fmt.Println("x is not int")
		return 
	}

	fmt.Println(v+1)
}
func main(){
	addOne(122)
	addOne("hello")

}