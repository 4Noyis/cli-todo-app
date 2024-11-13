// Run
// handleUserInput

package app

import (
	"fmt"

	"github.com/4Noyis/cli-todo-app/models"
	"github.com/4Noyis/cli-todo-app/storage"
	"github.com/4Noyis/cli-todo-app/ui"
	"github.com/4Noyis/cli-todo-app/utils"
)

func Run() {
	ui.CleanScreen()
	todos := models.Todos{}
	storage := storage.NewStorage[models.Todos]("data/todos.json")
	storage.Load(&todos)

	for {
		ui.CleanScreen()
		inputValue := ui.HomePage()
		handleUserInput(inputValue, &todos, storage)
	}

}

func handleUserInput(inputValue string, todos *models.Todos, storage *storage.Storage[models.Todos]) {
	switch inputValue {
	case "1":
		ui.CleanScreen()
		fmt.Print(todos.Print())
		fmt.Printf("Type 'b' to go back to the menu: ")
		utils.WaitForQuit()

	case "2":
		ui.CleanScreen()
		fmt.Print(todos.Print())
		for {
			indexValue := utils.IndexInputField("Type the index number of the selected to-do")
			if indexValue == -1 {
				break
			} else if err := todos.ValidateIndex(indexValue); err != nil {
				continue
			} else {
				todos.Toggle(indexValue)
				storage.Save(*todos)
				ui.CleanScreen()
				fmt.Print(todos.Print())
				fmt.Printf("Type 'b' to go back to the menu: ")
				utils.WaitForQuit()
				break
			}
		}

	case "3":
		newTodo := utils.TextInputField("Type the title of the new to-do")
		if newTodo == "b" {
			return
		} else {
			todos.Add(newTodo)
			storage.Save(*todos)
			ui.CleanScreen()
			fmt.Print(todos.Print())
			fmt.Printf("Type 'b' to go back to the menu: ")
			utils.WaitForQuit()
		}

	case "4":
		ui.CleanScreen()
		fmt.Print(todos.Print())
		for {
			indexValue := utils.IndexInputField("Type the index number for deleting the to-do")
			if indexValue == -1 {
				break
			} else if err := todos.ValidateIndex(indexValue); err != nil {
				continue
			} else {
				todos.Delete(indexValue)
				storage.Save(*todos)
				fmt.Print(todos.Print())
				fmt.Printf("Type 'b' to go back to the menu: ")
				utils.WaitForQuit()
				break
			}
		}

	case "5":
		ui.CleanScreen()
		fmt.Print(todos.Print())
		for {
			indexValue := utils.IndexInputField("Type the index number of the to-do to edit")
			if indexValue == -1 {
				break
			} else if err := todos.ValidateIndex(indexValue); err != nil {
				continue
			} else {
				editedText := utils.TextInputField("Type the new title for the to-do")
				if editedText == "b" {
					break
				} else {
					todos.Edit(indexValue, editedText)
					storage.Save(*todos)
					fmt.Print(todos.Print())
					fmt.Printf("Type 'b' to go back to the menu: ")
					utils.WaitForQuit()
					break
				}
			}
		}

	default:
		ui.CleanScreen()
	}
}
