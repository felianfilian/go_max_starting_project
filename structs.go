package main

import (
	"fmt"
	"structs/user"
)



func main() {
	userFirstName := getUserData("Please enter your first name: ")
	// created := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.NewUser(userFirstName )
	
	
	//user, err = &user.User()


	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputDetails()
	appUser.ClearUserData()
	//user.outputDetails()
}


func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
