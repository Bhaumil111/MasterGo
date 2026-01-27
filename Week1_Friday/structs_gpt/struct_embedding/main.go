// package main

// import "fmt"

// // type Address struct {
// // 	City  string
// // 	State string
// // }

// // type User struct {
// // 	Name string
// // 	Address
// // }
// type A struct {
// 	Value int
// }

// type B struct {
// 	Value int
// }

// type C struct {
// 	A
// 	B
// }

// func main() {

// 	// user := User{
// 	// 	Name: "Bhaumi",
// 	// 	Address: Address{
// 	// 		City:  "Ahmedabad",
// 	// 		State: "Gujarat",
// 	// 	},
// 	// }
// 	// fmt.Println("Name:", user.Name)
// 	// fmt.Println("City:", user.City)
// 	// fmt.Println("State:", user.State)
// 	// fmt.Println("State:", user.Address.State)

// 	C := C{
// 		A : A{
// 			Value: 32,
// 		},
// 		B : B{
// 			Value: 31,
// 		},
// 	}

// fmt.Println(C.A.Value)
// fmt.Println(C.B.Value)
	
// }
package main

import "fmt"

type Logger struct{}

func (l Logger) Log(msg string) {
	fmt.Println("LOG:", msg)
}

type Service struct {
	Logger
	Name string
}

func main() {
	s := Service{Name: "UserService"}
	s.Log("service started")
}
