package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Error reading balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Error parsing balance")
	}

	return balance, nil
}
func writeBalenceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}
func main() {
	balance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------------")
		panic("Something went wrong")
	}
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
			writeBalenceToFile(balance)
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
			writeBalenceToFile(balance)
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
