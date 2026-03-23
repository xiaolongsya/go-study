package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) (Admin, error) {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "User",
			birthdate: "01/01/2000",
			createdAt: time.Now(),
		},
	}, nil
}

func New(firstName, lastName, birthdate string) (User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return User{}, errors.New("Please enter the firstName and lastName and birthdate to create a new user")
	}
	return User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func (user User) OutputUserDetails() {
	fmt.Printf("First Name: %s\n", user.firstName)
	fmt.Printf("Last Name: %s\n", user.lastName)
	fmt.Printf("Birthdate: %s\n", user.birthdate)
	fmt.Printf("Created At: %s\n", user.createdAt.Format("2006-01-02 15:04:05"))
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
