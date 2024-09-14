package main

import (
	"fmt"
)



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
