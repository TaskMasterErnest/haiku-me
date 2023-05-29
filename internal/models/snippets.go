package models

import (
	"database/sql"
	"errors"
	"time"
)

type SnippetModelInterface interface {
	Insert(title string, content string, expires int) (int, error)
	Get(id int) (*Snippet, error)
	Latest() ([]*Snippet, error)
}

//define Snippet type to hold data for an individual snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// wrapping a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// insert new snippet into database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {

	statement := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(statement, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

// return snippet based on its id
func (m *SnippetModel) Get(id int) (*Snippet, error) {

	statement := `SELECT id, title, content, created, expires FROM snippets 
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(statement, id)

	s := &Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// return 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {

	statement := `SELECT id, title, content, created, expires FROM snippets 
	WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(statement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*Snippet{}

	for rows.Next() {

		s := &Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	// check to see if any errors occured when calling the rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// if OK
	return snippets, nil
}
