package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs/note"
	"example.com/structs/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}

	err = outputData(userTodo)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}

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

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the data failed:", err)
		return err
	}
	fmt.Println("Data saved successfully!")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}
