package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	store := &Store{
		conn: nil,
	} // initialize a store

	if err := store.Init(); err != nil {
		log.Fatalf("unable to init store: %v, err")
	}

	m := NewModel(store)

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		log.Fatalf("unable to run tui: %v", err)
	}
}
