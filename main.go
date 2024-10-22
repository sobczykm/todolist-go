package main

import (
	"sobczyk.dev/todolist/storage"
	"sobczyk.dev/todolist/todo"
)

func main() {
	todos := &todo.Todos{}
	storage := storage.Init[todo.Todos]("todos.json")

	defer storage.Save(todos)

	storage.Load(todos)
	todos.Add("Buy Milk")
	todos.Add("Buy Bread")
	todos.Toggle(0)
	todos.Print()
}
