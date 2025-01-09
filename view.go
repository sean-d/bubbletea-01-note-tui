package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	appNameStyle    = lipgloss.NewStyle().Background(lipgloss.Color("241")).Padding(0, 1)
	faintStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).Faint(true)
	enumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("100")).MarginRight(1)
)

func (m Model) View() string {

	s := appNameStyle.Render("NOTES APP") + "\n\n"

	if m.state == titleView {
		s += "note title:\n\n"
		s += m.textInput.View() + "\n\n"
		s += faintStyle.Render("enter - save, esc - discard")
	}

	if m.state == bodyView {
		s += "note:\n\n"
		s += m.textInput.View() + "\n\n"
		s += faintStyle.Render("ctrl+s - save, esc - discard")
	}

	if m.state == listView {
		for i, note := range m.notes {
			prefix := " " // default prefix for what's show in listView

			// add a > to the front of an item if the item is what you have selected
			if i == m.currentIndex {
				prefix = ">"
			}

			// shortBody is the note body without new lines...
			shortBody := strings.ReplaceAll(note.Body, "\n", " ")

			// shortBody will only be the first 30 bytes of the note body
			if len(shortBody) > 30 {
				shortBody = shortBody[:30]
			}

			s += fmt.Sprintf("%s %s | %s\n\n", enumeratorStyle.Render(prefix), note.Title, faintStyle.Render(shortBody))
		}

		s += faintStyle.Render("n - new note, q - quit")
	}
	return s
}
