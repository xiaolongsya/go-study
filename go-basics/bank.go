package main

import (
	"fmt"
	"os"
)
func main(){
	var balance float64 = 1000.0
	fmt.Println("Welcome to the Bank of Go")
	fmt.Println("Please select an option:")
	fmt.Println("1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Check balance")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your choice:")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Print("Enter amount to deposit: ")
		var amount float64
		fmt.Scan(&amount)
		balance += amount
		fmt.Println("Deposit successful. New balance:", balance)
	} else if choice == 2 {
		fmt.Print("Enter amount to withdraw: ")
		var amount float64
		fmt.Scan(&amount)
		if amount > balance {
			fmt.Println("Insufficient funds")
		} else {
			balance -= amount
			fmt.Println("Withdrawal successful. New balance:", balance)
		}
	} else if choice == 3 {
		fmt.Println("Your balance is:", balance)
	} else if choice == 4 {
		fmt.Println("Thank you for using the Bank of Go")
		os.Exit(0)
	} else {
		fmt.Println("Invalid choice")
	}
}