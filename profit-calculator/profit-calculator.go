package main

import (
	"errors"
	"fmt"
	"os"
)

const fileName = "results.txt"

func main() {
	revenue, err := getUserInput("Revenue: ")

	expenses, err := getUserInput("Expenses: ")

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	WriteResult(ebt, profit, ratio)
	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("The number must be positive")
	}
	return userInput, nil
}

func WriteResult(ebt, profit, ratio float64) {
	resultText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile(fileName, []byte(resultText), 0644)
}
