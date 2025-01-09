package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

// Note definition: unique ID, user submitted title and note body as strings
type Note struct {
	ID    int64
	Title string
	Body  string
}

// Store represents the database that contains the db connection
type Store struct {
	conn *sql.DB
}

// Init sets things up and will create the table if it does not exist
func (s *Store) Init() error {
	var err error

	s.conn, err = sql.Open("sqlite3", "notes.db")

	if err != nil {
		return err
	}

	createTableStatement := `CREATE TABLE IF NOT EXISTS notes (
        id INTEGER NOT NULL PRIMARY KEY,
        title TEXT NOT NULL,
        body TEXT NOT NULL
	);`

	_, err = s.conn.Exec(createTableStatement)

	if err != nil {
		return err
	}
	return nil
}

// GetNotes returns all notes as a list of notes and an error.
func (s *Store) GetNotes() ([]Note, error) {
	var err error
	var notes []Note

	s.conn, err = sql.Open("sqlite3", "notes.db")

	if err != nil {
		return []Note{}, err
	}

	getNotesStatement := `SELECT * FROM notes`

	rows, err := s.conn.Query(getNotesStatement)
	defer rows.Close()

	if err != nil {
		return []Note{}, err
	}

	for rows.Next() {
		var note Note
		_ = rows.Scan(&note.ID, &note.Title, &note.Body)
		notes = append(notes, note)
	}
	return notes, nil
}

// SaveNote takes in a Note and saves it. Returns any errors.
func (s *Store) SaveNote(note Note) error {
	if note.ID == 0 {
		note.ID = time.Now().UTC().UnixNano() // generate a unique ID
	}

	/*
		upsert is an insert clause that behaves like an update or no-op depending on constraints.

		excluded results in the title and body still being updated as if there were no conflict with the id field.
		So if we have note with ID 1 and we are changing the title or body, the changed fields will update for 1 even though
		1 already exists. the ID will not update.
	*/
	upsertStatement := `INSERT INTO notes (id, title, body)
    VALUES (?, ?, ?)
    ON CONFLICT(id) DO UPDATE
    SET title=excluded.title, body=excluded.body`

	_, err := s.conn.Exec(upsertStatement, note.ID, note.Title, note.Body)

	if err != nil {
		return err
	}
	return nil
}
