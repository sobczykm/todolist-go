package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Name        string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		fmt.Println("Todo index out of range")
		return errors.New("Todo index out of range")
	}

	return nil
}

func (todos *Todos) add(name string) {
	t := *todos
	todo := Todo{
		Name:        name,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(t, todo)
}

func (todos *Todos) delete(index int) error {
	t := *todos

	err := t.validateIndex(index)
	if err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos

	err := t.validateIndex(index)
	if err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if isCompleted {
		t[index].CompletedAt = nil
	} else {
		now := time.Now()
		t[index].CompletedAt = &now
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) update(index int, name string) error {
	t := *todos

	err := t.validateIndex(index)
	if err != nil {
		return err
	}

	t[index].Name = name

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)

	table.SetHeaders("#No.", "Name", "Completed", "Completed At", "Created At")

	for index, t := range *todos {
		completed := "✅"
		completedAt := ""

		if t.Completed {
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		} else {
			completed = "❌"
		}

		table.AddRow(strconv.Itoa(index), t.Name, completed, completedAt, t.CreatedAt.Format(time.RFC1123))
	}
	table.Render()
}
