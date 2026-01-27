package main

import (
	"fmt"
	"example.com/structs/Users"
)


func getUserData(prompText string ) string{
	fmt.Print(prompText)
	var value string
	fmt.Scanln(&value)
	return value

}

func main(){


	userfirstName :=getUserData("Please enter your first name: ")
	userlastName :=getUserData("Please enter your last name: ")
	userbirthdate :=getUserData("Please enter your birthdate: ")

	var appUser *users.User

	// appUser, err := newUser(userfirstName, userlastName,userbirthdate)

	appUser , err := users.NewUser(userfirstName, userlastName,userbirthdate)
	admin := users.NewAdmin("bhaumilpanchal@gmail.com", "test123")
	admin.OutputUserData()
	

	if(err !=nil){
		fmt.Println(err)
		return 

	}





	appUser.OutputUserData()
	appUser.ClearedUserDetails()
	appUser.OutputUserData()

}



