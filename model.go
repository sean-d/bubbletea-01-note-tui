package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

// enumerated states for the Model
const (
	_ = iota
	listView
	titleView
	bodyView
)

/*
	Model

state: see const above. default to listView
store: the db where notes are stored
notes: list of type []Note
currentNote: the note currently selected/being viewed
currentIndex: when scrolling up and down through notes, this represents the active/selected note
textArea: textarea.Model used for the content display of a note
textInput: textinput.Model used for entering the title of a note being created
*/
type Model struct {
	state        uint
	store        *Store
	notes        []Note
	currentNote  Note
	currentIndex int
	textArea     textarea.Model
	textInput    textinput.Model
}

// NewModel takes in a datastore and returns a Model using said Store that contains all available notes.
func NewModel(store *Store) Model {
	notes, err := store.GetNotes()

	if err != nil {
		log.Fatalf("uanble to get notes when creating a Model: %v", err)
	}

	return Model{
		state:     listView,
		store:     store,
		notes:     notes,
		textArea:  textarea.New(),
		textInput: textinput.New(),
	}
}

// Init can return a Cmd that could perform some initial I/O. For now, we don't need to do any I/O, so for the command,
// we'll just return nil, which translates to "no command."
func (m Model) Init() tea.Cmd {
	return nil
}

/*
Update is called when "things happen." Its job is to look at what has happened and return an updated Model in response.
It can also return a Cmd to make more things happen, but for now don't worry about that part.

In our case, when a user presses the down arrow, Update’s job is to notice that the down arrow was pressed and move
the cursor accordingly (or not).

The “something happened” comes in the form of a Msg, which can be any type. Messages are the result of some
I/O that took place, such as a keypress, timer tick, or a response from a server.

We usually figure out which type of Msg we received with a type switch, but you could also use a type assertion.



what is happening below:
- we update the Model's text area and input with the supplied message, assigning the update text accordingly and capturing the cmd
- we put the returned command into a slice
- switching:
-- switch based on msg type and look for KeyMsg (key presses)
-- string representation of the keypress is saved in key
-- switch on the current state as different keys will mean different things based on what state is current
-- switch on the key that was pressed and provide functionality based on that key value
- return the Model and all commands via Batch
*/

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		commands []tea.Cmd
		command  tea.Cmd
	)

	m.textArea, command = m.textArea.Update(msg)
	commands = append(commands, command)

	m.textInput, command = m.textInput.Update(msg)
	commands = append(commands, command)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String() //up, down, ctrl-c, etc...
		switch m.state {
		case listView:
			switch key {
			case "q":
				return m, tea.Quit
			case "n":
				m.textInput.SetValue("") // clear value...
				m.textInput.Focus()      // give focus
				m.currentNote = Note{}   // current note is now a new Note
				m.state = titleView
				// ... create a new note
			case "up", "k":
				// if the current highlighted note is not at the top of the list, move up.
				if m.currentIndex > 0 {
					m.currentIndex -= 1
				}
			case "down", "j":
				// if the current highlighted note is not at the bottom of the list, move down.
				if m.currentIndex < len(m.notes)-1 {
					m.currentIndex += 1
				}
			case "enter":
				m.currentNote = m.notes[m.currentIndex] // set currentNote to be what is selected when pressing enter
				m.textArea.SetValue(m.currentNote.Body) // set textArea to the body of the current note
				m.textArea.Focus()                      // may as well give it focus
				m.textArea.CursorEnd()                  // puts cursor at the end of the input field.
				m.state = bodyView
			}
		case titleView:
		case bodyView:
		}
	}
	return m, tea.Batch(commands...)
}
