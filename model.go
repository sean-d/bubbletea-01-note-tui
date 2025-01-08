package main

import tea "github.com/charmbracelet/bubbletea"

// enumerated states for the model
const (
	_ = iota
	listView
	titleView
	bodyView
)

type model struct {
	state uint
}

// NewModel takes in a datastore and returns a model using said Store
func NewModel() model {
	return model{state: listView}
}

// Init can return a Cmd that could perform some initial I/O. For now, we don't need to do any I/O, so for the command,
// we'll just return nil, which translates to "no command."
func (m model) Init() tea.Cmd {
	return nil
}

/*
Update is called when ”things happen.” Its job is to look at what has happened and return an updated model in response.
It can also return a Cmd to make more things happen, but for now don't worry about that part.

In our case, when a user presses the down arrow, Update’s job is to notice that the down arrow was pressed and move
the cursor accordingly (or not).

The “something happened” comes in the form of a Msg, which can be any type. Messages are the result of some
I/O that took place, such as a keypress, timer tick, or a response from a server.

We usually figure out which type of Msg we received with a type switch, but you could also use a type assertion.
*/
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
