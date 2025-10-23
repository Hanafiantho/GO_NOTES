package main

import (
	"fmt"

	"go-notes/helpers"
	"go-notes/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
	}

	userNote.Display()
	userNote.Save()
}

func getNoteData() (string, string) {
	title := helpers.GetUserInput("Note title: ")
	content := helpers.GetUserInput("Note content: ")

	return title, content
}
