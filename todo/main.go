package main

import (
	"fmt"
	"time"
)

type Todo struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt time.Time
}

var todos []Todo

func GetByID(id int) (*Todo, bool) {
	for i := range todos {
		if todos[i].ID == id {
			return &todos[i], true
		}
	}
	return nil, false
}

func Add(title string) {
	newTodo := Todo{
		ID:        len(todos) + 1,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	todos = append(todos, newTodo)
}
func List() {
	fmt.Println("---My To-Dos ---")
	for _, todo := range todos {
		fmt.Printf("ID: %d | Task: %q | Done: %t | CreatedAt: %s\n", todo.ID, todo.Title, todo.Done, todo.CreatedAt)
	}
}
func main() {
	Add("Write Go")
	Add("Buy Milk")
	List()
	GetByID(1)
}
