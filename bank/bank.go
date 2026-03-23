package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	balance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------------")
		panic("Something went wrong")
	}

	fmt.Println("Welcome to the Bank of Go")
	fmt.Println(randomdata.PhoneNumber())
	for {
		presentOptions()
		var choice int
		fmt.Print("Enter your choice:")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Print("Enter amount to deposit: ")
			var amount float64
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			balance += amount
			fileops.WriteFloatToFile(balance, accountBalanceFile)
			fmt.Println("Deposit successful. New balance:", balance)
		} else if choice == 2 {
			fmt.Println("Your balance is:", balance)
			fmt.Print("Enter amount to withdraw: ")
			var amount float64
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			if amount > balance {
				fmt.Println("Insufficient funds")
				continue
			}
			balance -= amount
			fileops.WriteFloatToFile(balance, accountBalanceFile)
			fmt.Println("Withdrawal successful. New balance:", balance)

		} else if choice == 3 {
			fmt.Println("Your balance is:", balance)
		} else if choice == 4 {
			break
		} else {
			fmt.Println("Invalid choice")
		}
	}

	fmt.Println("Thank you for using the Bank of Go")
}
