package users

import (
	"fmt"
	"errors"
	"time"
)


type User struct{
	firstName string
	lastName string 
	birthdate string
	createdAt time.Time
}


func (u *User) OutputUserData(){

	fmt.Println(u.firstName, u.lastName, u.birthdate)

}
func NewUser(firstName, lastName, birthDate string ) (*User, error) {

	if(firstName==""|| lastName==""|| birthDate==""){
		return nil, errors.New("Please enter firstName, lastName, birthDate")
	}
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func(u * User) ClearedUserDetails(){

	u.firstName = ""
	u.lastName = ""
}