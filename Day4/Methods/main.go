// package main

// import "fmt"
// type User struct{
// 	Name string
// }

// // func greet(u User){
// // 	u.Name = "New";
// // }

// func (u User) greet (){
// 	u.Name = "New"
// }


// func main(){
// 	u:= User{
// 		Name: "Bhaumil",
// 	}


// 	fmt.Println(u)
// 	// greet(u)
// 	u.greet()
// 	fmt.Println(u)

// }



package main

import "fmt"
type User struct{
	Name string
}

func (u *User) rename(){
	u.Name = "Panchal"

}



func main(){
 u := User{
	Name: "Bhaumil",
 }
 fmt.Printf("Before %v\n",u)
( &u).rename()
 fmt.Printf("After %v\n",u)

}