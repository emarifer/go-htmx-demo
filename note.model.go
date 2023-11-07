package main

import (
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type ConvertedNote struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
}

func (n *Note) CreateNote() (Note, error) {
	db := GetConnection()

	query := `INSERT INTO notes (title, description, created_at)
		VALUES(?, ?, ?) RETURNING *`

	stmt, err := db.Prepare(query)
	if err != nil {
		return Note{}, err
	}

	defer stmt.Close()

	var newNote Note
	err = stmt.QueryRow(
		n.Title,
		n.Description,
		time.Now().UTC(),
	).Scan(
		&newNote.ID,
		&newNote.Title,
		&newNote.Description,
		&newNote.Completed,
		&newNote.CreatedAt,
	)
	if err != nil {
		return Note{}, err
	}

	/* if i, err := result.RowsAffected(); err != nil || i != 1 {
		return errors.New("error: an affected row was expected")
	} */

	return newNote, nil
}

func (n *Note) GetAllNotes(offset int) ([]Note, error) {
	db := GetConnection()
	query := fmt.Sprintf("SELECT * FROM notes ORDER BY created_at DESC LIMIT 5 OFFSET %d", offset)

	rows, err := db.Query(query)
	if err != nil {
		return []Note{}, err
	}
	// Cerramos el recurso
	defer rows.Close()

	notes := []Note{}
	for rows.Next() {
		rows.Scan(&n.ID, &n.Title, &n.Description, &n.Completed, &n.CreatedAt)

		notes = append(notes, *n)
	}

	return notes, nil
}

func (n *Note) GetNoteById() (Note, error) {
	db := GetConnection()

	query := `SELECT * FROM notes
		WHERE id=?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return Note{}, err
	}

	defer stmt.Close()

	var recoveredNote Note
	err = stmt.QueryRow(
		n.ID,
	).Scan(
		&recoveredNote.ID,
		&recoveredNote.Title,
		&recoveredNote.Description,
		&recoveredNote.Completed,
		&recoveredNote.CreatedAt,
	)
	if err != nil {
		return Note{}, err
	}

	return recoveredNote, nil
}

func (n *Note) UpdateNote() (Note, error) {
	db := GetConnection()

	query := `UPDATE notes SET completed=?
		WHERE id=? RETURNING *`

	stmt, err := db.Prepare(query)
	if err != nil {
		return Note{}, err
	}

	defer stmt.Close()

	var updatedNote Note
	err = stmt.QueryRow(
		!n.Completed,
		n.ID,
	).Scan(
		&updatedNote.ID,
		&updatedNote.Title,
		&updatedNote.Description,
		&updatedNote.Completed,
		&updatedNote.CreatedAt,
	)
	if err != nil {
		return Note{}, err
	}

	return updatedNote, nil
}

func (n *Note) DeleteNote() error {
	db := GetConnection()

	query := `DELETE FROM notes
		WHERE id=?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(n.ID)
	if err != nil {
		return err
	}

	if i, err := result.RowsAffected(); err != nil || i != 1 {
		return errors.New("an affected row was expected")
	}

	return nil
}

func convertDateTime(note Note, timeZone string) ConvertedNote {
	loc, _ := time.LoadLocation(timeZone)
	convertedNote := ConvertedNote{
		ID:          note.ID,
		Title:       note.Title,
		Description: note.Description,
		Completed:   note.Completed,
		CreatedAt:   note.CreatedAt.In(loc).Format(time.RFC822Z),
	}

	return convertedNote
}
