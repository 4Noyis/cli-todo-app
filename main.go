package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func HomePage() string {

	InputValue := ""
	quit := true

	for quit {
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

func TextInputField(printText string) string {
	var inputValue string

	for {
		fmt.Print(printText + "-> ")

		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		input := strings.TrimSpace(line)

		if input == "b" {
			fmt.Print("\033[H\033[2J")
			HomePage()
		}

		inputValue = input
		break
	}
	return inputValue
}

func IndexInputField() int {
	var inputValue int

	for {
		fmt.Print("To-do index: ")

		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		input := strings.TrimSpace(line)

		if input == "b" {
			fmt.Print("\033[H\033[2J")
			HomePage()
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
			fmt.Println("Exiting...")
			fmt.Print("\033[H\033[2J")
			return
		}
	}
}

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	quit := true
	for quit {
		inputValue := HomePage()
		if inputValue == "1" {

			fmt.Print(todos.Print())
			fmt.Printf("type 'b' for go back to menu:")
			waitForQuit()
		}
		if inputValue == "2" {
			indexValue := IndexInputField()
			todos.Toggle(indexValue)
			storage.Save(todos)
			fmt.Print(todos.Print())
			fmt.Printf("type 'b' for go back to menu:")
			waitForQuit()
		}
		if inputValue == "3" {
			newTodo := TextInputField("yo")
			todos.Add(newTodo)
			storage.Save(todos)
			fmt.Print(todos.Print())
			fmt.Printf("type 'b' for go back to menu:")
			waitForQuit()
		}
		if inputValue == "4" {
			indexValue := IndexInputField()
			todos.Delete(indexValue)
			storage.Save(todos)
			fmt.Print(todos.Print())
			fmt.Printf("type 'b' for go back to menu: ")
			waitForQuit()
		}
		if inputValue == "5" {
			fmt.Print(todos.Print())
			indexValue := IndexInputField()
			editedText := TextInputField("yooo")
			todos.Edit(indexValue, editedText)
			storage.Save(todos)
			fmt.Print(todos.Print())
			fmt.Printf("type 'b' for go back to menu: ")
			waitForQuit()
		}
	}

}
