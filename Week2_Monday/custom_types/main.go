// package main

// import "fmt"


// func GetUserById(userId int){
// 	fmt.Println("This is my id", userId)
// }
// func GetOrderById(orderId int){
// 	fmt.Println("This is my id", orderId)
// }

// func main(){

// 	userId := 1000
// 	orderId :=2000
// 	// This is serious bug order id passed as userId
// 	GetUserById(orderId)
// 	GetOrderById(userId)



// }


//Now with Custom Types=================

package main
import "fmt"

type UserId int
type OrderId int

func GetUserById(userId UserId){
	fmt.Println("This is my userId", userId)
}

func GetOrderById(orderId OrderId){
	fmt.Println("This is my orderId", orderId)
}


func main(){

	var userId UserId =1000
	var orderId OrderId =2000
	GetOrderById(orderId )
	GetUserById(userId)
}


