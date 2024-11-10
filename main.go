package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func HomePage() string {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	InputValue := ""
	quit := true

	for quit {
		fmt.Print(todos.Print())
		fmt.Printf("1- Show To-Do List\n2- Mark Done Selected To-do\n3- Add New To-Do\n4- Delete To-Do\n5- Edit To-Do\n")
		fmt.Print("Enter a command: ")

		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')

		var input string

		splittedString := strings.Split(line, "\n")

		if splittedString[0] == "exit" || line == "exit" {
			fmt.Println("[Info] Bye!")
			os.Exit(0)
		} else {
			input = splittedString[0]
			quit = false
			InputValue = input
		}
	}
	return InputValue

}

func TextInputField(headerText string) string {
	const exitIndicator = "b"
	var inputValue string

	for {
		fmt.Print(headerText + "-> ")

		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		input := strings.TrimSpace(line)

		if input == "b" {
			fmt.Print("\033[H\033[2J")
			HomePage()
			return exitIndicator
		}

		inputValue = input
		break
	}
	return inputValue
}

func IndexInputField(headerText string) int {
	const exitIndicator = -1
	var inputValue int

	for {
		fmt.Print(headerText + "--> ")

		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		input := strings.TrimSpace(line)

		if input == "b" {
			fmt.Print("\033[H\033[2J")
			HomePage()
			return exitIndicator
		}

		// Try to parse the input as an integer
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("[Error] Please enter a valid integer or type 'exit' to quit.")
			continue // Prompt again if input is invalid
		}

		inputValue = value
		break
	}
	return inputValue
}

func waitForQuit() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		if input == "b\n" {
			fmt.Print("\033[H\033[2J")
			return
		}
	}
}

func cleanScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	quit := true
	for quit {
		inputValue := HomePage()
		// SHOW TABLE
		if inputValue == "1" {
			cleanScreen()
			fmt.Print(todos.Print())
			fmt.Printf("type 'b' for go back to menu:")
			waitForQuit()
		}
		// MARK DONE SELECTED TODO
		if inputValue == "2" {
			cleanScreen()
			fmt.Print(todos.Print())
			indexValue := IndexInputField("Type index number of seleceted to-do ")
			if indexValue == -1 {
				continue
			} else {
				todos.Toggle(indexValue)
				storage.Save(todos)
				cleanScreen()
				fmt.Print(todos.Print())
				fmt.Printf("type 'b' for go back to menu:")
				waitForQuit()
			}
		}
		// ADD NEW TODO
		if inputValue == "3" {
			newTodo := TextInputField("Type title of new to-do ")
			if newTodo == "b" {
				continue
			} else {
				todos.Add(newTodo)
				storage.Save(todos)
				cleanScreen()
				fmt.Print(todos.Print())
				fmt.Printf("type 'b' for go back to menu:")
				waitForQuit()
			}
		}
		// DELETE TODO
		if inputValue == "4" {
			cleanScreen()
			fmt.Print(todos.Print())
			indexValue := IndexInputField("Type index number for deleting to-do ")
			if indexValue == -1 {
				continue
			} else {
				todos.Delete(indexValue)
				storage.Save(todos)
				fmt.Print(todos.Print())
				fmt.Printf("type 'b' for go back to menu: ")
				waitForQuit()
			}
		}
		// EDIT TODO
		if inputValue == "5" {
			cleanScreen()
			fmt.Print(todos.Print())
			indexValue := IndexInputField("Type index number of selected to-do for editting ")
			if indexValue == -1 {
				continue
			} else {
				editedText := TextInputField("Type new title of edited to-do ")
				if editedText == "b" {
					continue
				} else {
					todos.Edit(indexValue, editedText)
					storage.Save(todos)
					fmt.Print(todos.Print())
					fmt.Printf("type 'b' for go back to menu: ")
					waitForQuit()
				}
			}

		} else {
		}
	}

}
