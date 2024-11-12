// Textinputfield
// Indexinputfield
// Waitforquit

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
			return exitIndicator
		} else {
			inputValue = input
			break
		}
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
			return exitIndicator
		}

		// Try to parse the input as an integer
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("[Error] Please enter a valid integer or type 'b' to quit.")
			continue // Prompt again if input is invalid
		}

		inputValue = value
		break
	}
	return inputValue
}
func WaitForQuit() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		if input == "b\n" {
			CleanScreen()
			return
		}
	}
}
