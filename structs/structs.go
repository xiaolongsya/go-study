package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := user.GetUserData("Please enter your first name: ")
	userLastName := user.GetUserData("Please enter your last name: ")
	userBirthdate := user.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	userAdminEmail := user.GetUserData("Please enter the admin email: ")
	userAdminPassword := user.GetUserData("Please enter the admin password: ")

	adminUser, err := user.NewAdmin(userAdminEmail, userAdminPassword)
	if err != nil {
		fmt.Println(err)
		return
	}
	// ... do something awesome with that gathered data!

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	adminUser.OutputUserDetails()
	adminUser.ClearUserName()
	adminUser.OutputUserDetails()

}
