package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// func main() {
// 	p := tea.NewProgram(InitialInputModel())
// 	if _, err := p.Run(); err != nil {
// 		log.Fatal(err)
// 	}
// }

type (
	errMsg error
)

type Inputmodel struct {
	textInput  textinput.Model
	err        error
	inputValue string // Field to store user input
}

func InitialInputModel() Inputmodel {
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Inputmodel{
		textInput:  ti,
		err:        nil,
		inputValue: "", // Initialize inputValue to an empty string
	}
}

func (m Inputmodel) Init() tea.Cmd {
	return textinput.Blink
}

func (m Inputmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			// Capture the input when Enter is pressed and store it in inputValue
			m.inputValue = m.textInput.Value()
			return m, nil

		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Inputmodel) View() string {
	return fmt.Sprintf(
		"What’s your favorite Pokémon?\n\n%s\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
		"You entered: "+m.inputValue, // Display the entered value
	) + "\n"
}
