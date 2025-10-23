package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`     // (using struct tag)
	Content   string    `json:"content"`   // (using struct tag)
	CreatedAt time.Time `json:"createdAt"` // (using struct tag)
}

func (note Note) Display() {
	fmt.Printf("%v\n%v", note.Title, note.Content)
}

func (note Note) Save() error {
	// creating file name based on note title
	fileName := strings.ReplaceAll(note.Title, " ", "_") // converting all space on title into _
	fileName = strings.ToLower(fileName) + ".json"       // convert all upper case into lower case

	// conver note into json
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content cannot be empty")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
