// Run
// handleUserInput

package main

import (
	"fmt"
)

func Run() {
	CleanScreen()
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	for {
		CleanScreen()
		inputValue := HomePage()
		handleUserInput(inputValue, &todos, storage)
	}

}

func handleUserInput(inputValue string, todos *Todos, storage *Storage[Todos]) {
	switch inputValue {
	case "1":
		CleanScreen()
		fmt.Print(todos.Print())
		fmt.Printf("Type 'b' to go back to the menu: ")
		WaitForQuit()

	case "2":
		CleanScreen()
		fmt.Print(todos.Print())
		for {
			indexValue := IndexInputField("Type the index number of the selected to-do")
			if indexValue == -1 {
				break
			} else if err := todos.ValidateIndex(indexValue); err != nil {
				continue
			} else {
				todos.Toggle(indexValue)
				storage.Save(*todos)
				CleanScreen()
				fmt.Print(todos.Print())
				fmt.Printf("Type 'b' to go back to the menu: ")
				WaitForQuit()
				break
			}
		}

	case "3":
		newTodo := TextInputField("Type the title of the new to-do")
		if newTodo == "b" {
			return
		} else {
			todos.Add(newTodo)
			storage.Save(*todos)
			CleanScreen()
			fmt.Print(todos.Print())
			fmt.Printf("Type 'b' to go back to the menu: ")
			WaitForQuit()
		}

	case "4":
		CleanScreen()
		fmt.Print(todos.Print())
		for {
			indexValue := IndexInputField("Type the index number for deleting the to-do")
			if indexValue == -1 {
				break
			} else if err := todos.ValidateIndex(indexValue); err != nil {
				continue
			} else {
				todos.Delete(indexValue)
				storage.Save(*todos)
				fmt.Print(todos.Print())
				fmt.Printf("Type 'b' to go back to the menu: ")
				WaitForQuit()
				break
			}
		}

	case "5":
		CleanScreen()
		fmt.Print(todos.Print())
		for {
			indexValue := IndexInputField("Type the index number of the to-do to edit")
			if indexValue == -1 {
				break
			} else if err := todos.ValidateIndex(indexValue); err != nil {
				continue
			} else {
				editedText := TextInputField("Type the new title for the to-do")
				if editedText == "b" {
					break
				} else {
					todos.Edit(indexValue, editedText)
					storage.Save(*todos)
					fmt.Print(todos.Print())
					fmt.Printf("Type 'b' to go back to the menu: ")
					WaitForQuit()
					break
				}
			}
		}

	default:
		CleanScreen()
	}
}
