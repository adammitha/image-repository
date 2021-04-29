package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	fmt.Println("Hello, world!")
}

// Init initializes the command line interface to create or interact with an image repository
func (r Repository) Init() tea.Cmd {
	return nil
}
