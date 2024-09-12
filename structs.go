package main

import (
	"fmt"
	"time"
)

type User struct {
	firstname string
	lastname string
	created time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	// created := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var user User
	
	user = User{
		firstname: firstName,
		lastname: lastName,
		created: time.Now(),
	}

	outputDetails(user)
}

func outputDetails(user User) {
	fmt.Println()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
