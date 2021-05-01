package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	r := NewRepository("/Users/adammitha/Downloads")
	err := r.AddImage("https://image.freepik.com/free-vector/shining-circle-purple-lighting-isolated-dark-background_1441-2396.jpg")
	if err != nil {
		log.Fatal(err)
	}
}

// Init initializes the command line interface to create or interact with an image repository
func (r Repository) Init() tea.Cmd {
	return nil
}
