package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	appNameStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6C11")).Background(lipgloss.Color("#241734")).Padding(0, 1)
	bottomMenuStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#791E94"))
	noteTitleStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#920075")).Underline(true)
	noteStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#EE227D")).Underline(true)
	deleteStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	enumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F706CF")).MarginRight(1)
)

func (m Model) View() string {

	s := appNameStyle.Render("THE MOST BAREBONES BUBBLETEA MINIMUM VIABLE SOLUTION NOTES APP") + "\n\n"

	if m.state == addNoteView {
		s += "note title:\n\n"
		s += m.textInput.View() + "\n\n"
		s += bottomMenuStyle.Render("enter - submit title, esc - discard")
	}

	if m.state == noteView {
		s += fmt.Sprintf("%s:\n\n", noteTitleStyle.Render(m.currentNote.Title))
		s += m.textArea.View() + "\n\n"
		s += bottomMenuStyle.Render("ctrl+s - save, , ctrl+d - delete, esc - discard")
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

			s += fmt.Sprintf("%s %s | %s\n\n", enumeratorStyle.Render(prefix), noteTitleStyle.Render(note.Title), noteStyle.Render(shortBody))
		}

		s += bottomMenuStyle.Render("n - new note, q/esc - quit")

	}
	if m.state == deleteView {
		s += noteTitleStyle.Render("delete note:") + "\n\n"
		/*
			display the note to be deleted.
			ask to confirm at bottom
			switch on the key y/n
			if yes, run delete and return to listView
			if no, return to bodyView

			add "d" to bodyView as well

			maybe have a flag set so we track where we came from...so if src was list, we return there. if source was body, we return there

		*/
		s += m.textArea.View() + "\n\n"
		s += deleteStyle.Render("ctrl+d - confirm delete, esc - discard")
	}
	return s
}
