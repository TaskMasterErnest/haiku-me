package models

import (
	"database/sql"
	"time"
)

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

	return nil, nil
}

// return 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
