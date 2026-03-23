package main

import (
	"fmt"
)

type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	var text str
	fmt.Println("Enter some text:")
	fmt.Scan(&text)
	text.log()
}
