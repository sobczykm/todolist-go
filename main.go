package main

import (
	"sobczyk.dev/todolist/cmd"
	"sobczyk.dev/todolist/storage"
	"sobczyk.dev/todolist/todo"
)

func main() {
	todos := &todo.Todos{}
	storage := storage.Init[todo.Todos]("todos.json")
	storage.Load(todos)

	defer storage.Save(todos)

	cmdFlags := cmd.NewCmdFlags()
	cmdFlags.Execute(todos)
}
