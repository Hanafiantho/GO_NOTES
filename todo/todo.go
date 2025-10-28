package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Content)
}

func (todo Todo) Save() error {
	filename := "todo" + time.Now().String() + ".json"

	json, err := json.Marshal(todo)

	if err != nil {
		fmt.Print(err)
		return err
	}

	fmt.Print("Success save data to file")
	return os.WriteFile(filename, json, 0644)
}

func New(content string) (*Todo, error) {
	if content == "" {
		return nil, errors.New("content cannot be empty")
	}

	return &Todo{
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
