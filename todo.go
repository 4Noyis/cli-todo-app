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
	Title        string
	Completed    bool
	CreatedAt    time.Time
	CompleteTime *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:        title,
		Completed:    false,
		CreatedAt:    time.Now(),
		CompleteTime: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) ValidateIndex(index int) error {
	if index < 0 || index > len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) Delete(index int) error {
	t := *todos

	if err := t.ValidateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) Toggle(index int) error {
	t := *todos

	if err := t.ValidateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	/*
		isCompleted değişkeni true ise !isCompleted false olur.
		isCompleted değişkeni false ise !isCompleted true olur.
	*/
	if !isCompleted {
		completionTime := time.Now()
		t[index].CompleteTime = &completionTime
	}

	t[index].Completed = !isCompleted

	return nil

}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos

	if err := t.ValidateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil

}

func (todos *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "X"
		completeTime := ""

		if t.Completed {
			completed = "+"
			if t.CompleteTime != nil {
				completeTime = t.CompleteTime.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completeTime)
	}

	table.Render()
}
