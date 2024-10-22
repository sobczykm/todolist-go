package cmd

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"sobczyk.dev/todolist/todo"
)

type CmdFlags struct {
	Add    string
	Edit   string
	Delete int
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "adds a new todo")
	flag.StringVar(&cf.Edit, "edit", "", "edits a  todo")
	flag.IntVar(&cf.Delete, "delete", -1, "deletes a todo")
	flag.IntVar(&cf.Toggle, "toggle", -1, "toggles a completion for todo")
	flag.BoolVar(&cf.List, "list", false, "lists todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *todo.Todos) {
	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.Split(cf.Edit, ":")
		if len(parts) != 2 {
			fmt.Println("invalid format for edit, use id:title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("invalid index for edit")
			os.Exit(1)
		}

		todos.Update(index, parts[1])
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	case cf.Delete != -1:
		todos.Delete(cf.Delete)
	default:
		fmt.Println("Invalid command")
	}
}
