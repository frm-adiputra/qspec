// Package sample provides sample query.
package sample

// DO NOT EDIT: Code generated by "qspec samples/sample.yml"

import "database/sql"

const (
	_UpdateStmt = `UPDATE Sample
SET
  ID = ?,
  Title = ?,
  Content = ?
WHERE ID = ?;
`
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

// Update updates sample.
func Update(id, title, content interface{}) (sql.Result, error) {
	r, err := getDB().Exec(_UpdateStmt, id, title, content, id)
	return r, err
}

// PreparedUpdate represents the prepared Update statement.
type PreparedUpdate struct {
	stmt *sql.Stmt
}

// PrepareUpdate returns the prepared Update statement.
func PrepareUpdate() (*PreparedUpdate, error) {
	stmt, err := getDB().Prepare(_UpdateStmt)
	if err != nil {
		return nil, err
	}

	return &PreparedUpdate{stmt}, nil
}

// Close closes the prepared Update statement.
func (p *PreparedUpdate) Close() error {
	return p.stmt.Close()
}

// Exec executes the prepared Update statement.
func (p *PreparedUpdate) Exec(id, title, content interface{}) (sql.Result, error) {
	r, err := p.stmt.Exec(id, title, content, id)
	return r, err
}
