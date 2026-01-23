package main

import (

"fmt"
"time"

)




type user struct{
	firstName string
	lastName string 
	birthdate string
	createdAt time.Time
}

func getUserData(prompText string ) string{
	fmt.Print(prompText)
	var value string
	fmt.Scan(&value)
	return value

}

func(u * user) clearedUserDetails(){

	u.firstName = ""
	u.lastName = ""
}
func main(){


	userfirstName :=getUserData("Please enter your first name: ")
	userlastName :=getUserData("Please enter your last name: ")
	userbirthdate :=getUserData("Please enter your birthdate: ")

	var appUser user


	appUser = user{
		firstName: userfirstName,
		lastName: userlastName,
		birthdate: userbirthdate,
		createdAt: time.Now() ,
	}

	appUser.outputUserData()
	appUser.clearedUserDetails()
	appUser.outputUserData()

}



func (u *user) outputUserData(){

	fmt.Println(u.firstName, u.lastName, u.birthdate)

}