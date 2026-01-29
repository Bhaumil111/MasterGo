package main


import "fmt"


type User struct{
	Name string
	Age int
}

func birthday(u User){
	u.Age++
}


func birthdayAdd(u * User){
	u.Age = 340
}
func main(){
	u := User{"Alice", 30}
	birthday(u)

	fmt.Println(u)
	birthdayAdd(&u)
	fmt.Println(u)

	
}







