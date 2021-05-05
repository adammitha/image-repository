package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	r := NewRepository("/Users/adammitha/Downloads")
	cli := tea.NewProgram(r)
	if err := cli.Start(); err != nil {
		log.Fatalf("Alas, there's been an error: %v", err)
	}
}
