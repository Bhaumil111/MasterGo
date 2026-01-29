package main 
import "fmt"

func main(){
// 	m:= make(map[string]int)
// 	m["Bhaumil"] = 1
// 	m["Panchal"] = 2

// 	fmt.Println("Learning about map")
// 	fmt.Println(m)




	// for  k , v := range(m){
	// 	println(k,v)
	// }
	// for  k , v := range(m){
	// 	println(k,v)
	// }
	// var m map[string]int{
	// 	"a":1,
	// 	"b":2,
	// }
	var m map[string] int

	fmt.Println(m["a"])
	// m["a"] = 1

	for k , v :=range(m){
		fmt.Println(k,v)

	}

}
