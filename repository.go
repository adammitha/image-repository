package main

import tea "github.com/charmbracelet/bubbletea"

// Repository stores the application's state
type Repository struct {
	root string
}

// NewRepository initializes the application's state with an image repository located at the path `root`
func NewRepository(root string) *Repository {
	return &Repository{
		root: root,
	}
}

// Init initializes the command line interface to create or interact with an image repository
func (r Repository) Init() tea.Cmd {
	return nil
}
