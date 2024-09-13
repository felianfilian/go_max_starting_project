package main

import (
	"fmt"
	"time"
)

type User struct {
	firstname string
	created time.Time
}

func (u User) outputDetails() {
	fmt.Println(u.firstname)
}

func (u *User) clearUserData() {
	u.firstname = ""
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	// created := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var user User
	
	user = User{
		firstname: firstName,
		created: time.Now(),
	}

	user.outputDetails()
	user.clearUserData()
	user.outputDetails()
}


func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
