// package main
// import "fmt"
// type User struct{
// 	Name string
// }


// func (u User ) sayHello(){
// 	fmt.Printf("Hello %v",u.Name)
// }



// type Helloer interface{
// 	sayHello()
	
// }
// func main(){
// 	// u := User{
// 	// 	Name: "Bhaumil",
// 	// }

// 	// u.sayHello()
// 	var h Helloer
// 	h = User{
// 		Name: "Bhaumil",
// 	}
// 	h.sayHello()
	

// }
package main
import "fmt"


type User struct{
	Name string
}
type Robot struct{
	SurName string
}

func (u User) speak(){
	fmt.Printf("Hello %v\n", u.Name)
}

func (r Robot) speak(){
	
	fmt.Printf("Hello %v\n", r.SurName)
}


// func AnnounceSpeaker(u User){
// 	u.speak()
	
// }
type Speaker interface{
	speak()
}

func AnnounceSpeaker(s Speaker){
	s.speak()
}



func main(){
	u:= User{
		Name: "Bhaumil",
	}
	r := Robot{
		SurName: "Panchal",
	}

	AnnounceSpeaker(u)
	AnnounceSpeaker(r) //Here goes problem without interfaces

}
