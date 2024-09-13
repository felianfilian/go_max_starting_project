package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstname string
	created time.Time
}

func newUser(firstname string) (*User, error) {
	if firstname == "" {
		return nil, errors.New("firstname needed")
	} else {
		return &User {
			firstname: firstname,
			created: time.Now(),
		}, nil
	}
}

func (u User) outputDetails() {
	fmt.Println(u.firstname)
	fmt.Println(u.created)
}

func (u *User) clearUserData() {
	u.firstname = ""
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	// created := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var user *User
	user, err := newUser(firstName)

	if err != nil {
		fmt.Println(err)
		return
	}

	user.outputDetails()
	user.clearUserData()
	//user.outputDetails()
}


func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
