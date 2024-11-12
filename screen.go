// CleanScreen
// HomePage
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CleanScreen() {
	fmt.Print("\033[H\033[2J")
}

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
