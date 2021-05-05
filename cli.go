package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Init initializes the command line interface to create or interact with an image repository
func (r *Repository) Init() tea.Cmd {
	return textinput.Blink
}

func (r *Repository) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return r, tea.Quit
		}
	}

	r.textInput, cmd = r.textInput.Update(msg)

	return r, cmd
}

func (r *Repository) View() string {
	return fmt.Sprintf(
		"Please enter the directory of your image repository%s", r.textInput.View(),
	)
}
