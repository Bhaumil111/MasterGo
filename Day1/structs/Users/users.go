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


type Admin struct {
	email string
	password string
	User

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

func NewAdmin(email, password string ) Admin{
	return Admin{
		email :email,
		password: password,
		User :User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthdate: "----",
			createdAt: time.Now(),
		},
	}
}


func(u * User) ClearedUserDetails(){

	u.firstName = ""
	u.lastName = ""
}