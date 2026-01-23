package main

import "fmt"



func main(){




	// result:= func( a , b int )int{

	// 	return a+b
	// }(2,3)

	// add := func(a, b int)int{
	// 	return a+b
	// }
	// fmt.Println(result)
	// fmt.Println(add(2,3))


	for i:=0 ;i<3;i++{
		func(){
			fmt.Println(i)
		}()
	}

}