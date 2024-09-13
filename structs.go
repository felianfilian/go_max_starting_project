package main

import (
	"fmt"
	"time"
)

type User struct {
	firstname string
	created time.Time
}

func newUser(firstname string) *User {
	return &User {
		firstname: firstname,
		created: time.Now(),
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

	var user *User = newUser(firstName)

	user.outputDetails()
	user.clearUserData()
	//user.outputDetails()
}


func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
