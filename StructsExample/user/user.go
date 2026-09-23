package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func NewAdmin(email, password string) Admin{
	return Admin {
		email: email,
		password: password,
		User{
			firstName: "ADMIN",
			lastName: "Admin",
			birthDate: "12/04/1983",
			createdAt: time.Now()	
		}
	}
}
func (u *User) OutputUserDetails(){
	fmt.Println(u.firstName, u.lastName,u.birthDate)
}

func ( u *User) ClearUserName(){
	u.firstName = ""
	u.lastName = ""
}

func NewUser(firstName,lastName,birthDate string) (*User,error){
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("FirstName, Last Name and Birthdate cannot be Null")
	}

	return &User {
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	},nil
}