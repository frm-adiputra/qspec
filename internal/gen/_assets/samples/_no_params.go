// Package sample provides sample model.
package sample

// DO NOT EDIT: Code generated by "qspec samples/sample.yml"

import "database/sql"

const (
	_SelectStmt = `SELECT * FROM Sample;`
)

var (
	// DB defines the database that will be used for this model.
	DB *sql.DB
)

func getDB() *sql.DB {
	if DB == nil {
		panic("DB not set")
	}
	return DB
}

// SelectResult represents the result of Select query.
type SelectResult struct {
	ID      string `json:"id,string"`
	Title   string
	Content string
}

// Select select by ID and Title.
func Select() (SelectResult, error) {
	var v SelectResult
	err := getDB().QueryRow(_SelectStmt).Scan(&v.ID, &v.Title, &v.Content)
	if err != nil {
		return SelectResult{}, err
	}

	return v, nil
}

// PreparedSelect represents the prepared Select statement.
type PreparedSelect struct {
	stmt *sql.Stmt
}

// PrepareSelect returns the prepared Select statement.
func PrepareSelect() (*PreparedSelect, error) {
	stmt, err := getDB().Prepare(_SelectStmt)
	if err != nil {
		return nil, err
	}

	return &PreparedSelect{stmt}, nil
}

// Close closes the prepared Select statement.
func (p *PreparedSelect) Close() error {
	return p.stmt.Close()
}

// Query executes the prepared Select statement.
func (p *PreparedSelect) Query() (SelectResult, error) {
	var v SelectResult
	err := p.stmt.QueryRow().Scan(&v.ID, &v.Title, &v.Content)
	if err != nil {
		return SelectResult{}, err
	}

	return v, nil
}
