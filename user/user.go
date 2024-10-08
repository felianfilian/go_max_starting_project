package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstname string
	created time.Time
}


type Admin struct {
	email string
	User
}

func NewUser(firstname string) (*User, error) {
	if firstname == "" {
		return nil, errors.New("firstname needed")
	} else {
		return &User {
			firstname: firstname,
			created: time.Now(),
		}, nil
	}
}

func NewAdmin(email string) Admin{
	return Admin {
		email: email,
		User: User {
			firstname: "Admin",
		},
	}
}

func (u *User) OutputDetails() {
	fmt.Println(u.firstname)
	fmt.Println(u.created)
}

func (u *User) ClearUserData() {
	u.firstname = ""
}