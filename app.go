package main

import (
	"fmt"

	"go-notes/helpers"
	"go-notes/note"
	"go-notes/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	userNote, noteErr := note.New(title, content)

	if noteErr != nil {
		fmt.Println(noteErr)
		return
	}

	todoContent := getTodoData()
	userTodo, todoErr := todo.New(todoContent)

	if todoErr != nil {
		fmt.Println(todoErr)
		return
	}

	userNote.Display()
	saveData(userNote)

	userTodo.Display()
	saveData(userTodo)
}

func saveData(data saver) {
	err := data.Save()

	if err != nil {
		fmt.Print("failed to save data")
		return
	}

	fmt.Print("Success!")
}

func getNoteData() (string, string) {
	title := helpers.GetUserInput("Note title: ")
	content := helpers.GetUserInput("Note content: ")

	return title, content
}

func getTodoData() string {
	content := helpers.GetUserInput("Todo Content: ")

	return content
}
