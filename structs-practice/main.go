package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs/note"
	"example.com/structs/todo"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}

	userTodo.Display()
	err = userTodo.Save()
	if err != nil {
		fmt.Println("Error saving todo:", err)
		return
	}
	fmt.Println("Todo saved successfully!")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
	fmt.Println("Note saved successfully!")

}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
