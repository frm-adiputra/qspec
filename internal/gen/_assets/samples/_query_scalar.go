// Package sample provides sample query.
package sample

// DO NOT EDIT: Code generated by "qspec samples/sample.yml"

import "database/sql"

const (
	_CountStmt         = `SELECT COUNT(*) FROM Sample;`
	_CountDisabledStmt = `SELECT COUNT(*) FROM Sample WHERE disabled = ?;`
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

// Count counts sample.
func Count() (int64, error) {
	var v int64
	err := getDB().QueryRow(_CountStmt).Scan(&v)
	if err != nil {
		return v, err
	}

	return v, nil
}

// PreparedCount represents the prepared Count statement.
type PreparedCount struct {
	stmt *sql.Stmt
}

// PrepareCount returns the prepared Count statement.
func PrepareCount() (*PreparedCount, error) {
	stmt, err := getDB().Prepare(_CountStmt)
	if err != nil {
		return nil, err
	}

	return &PreparedCount{stmt}, nil
}

// Close closes the prepared Count statement.
func (p *PreparedCount) Close() error {
	return p.stmt.Close()
}

// Query executes the prepared Count statement.
func (p *PreparedCount) Query() (int64, error) {
	var v int64
	err := p.stmt.QueryRow().Scan(&v)
	if err != nil {
		return v, err
	}

	return v, nil
}

// CountDisabled counts sample by disabled.
func CountDisabled(disabled interface{}) (int64, error) {
	var v int64
	err := getDB().QueryRow(_CountDisabledStmt, disabled).Scan(&v)
	if err != nil {
		return v, err
	}

	return v, nil
}

// PreparedCountDisabled represents the prepared CountDisabled statement.
type PreparedCountDisabled struct {
	stmt *sql.Stmt
}

// PrepareCountDisabled returns the prepared CountDisabled statement.
func PrepareCountDisabled() (*PreparedCountDisabled, error) {
	stmt, err := getDB().Prepare(_CountDisabledStmt)
	if err != nil {
		return nil, err
	}

	return &PreparedCountDisabled{stmt}, nil
}

// Close closes the prepared CountDisabled statement.
func (p *PreparedCountDisabled) Close() error {
	return p.stmt.Close()
}

// Query executes the prepared CountDisabled statement.
func (p *PreparedCountDisabled) Query(disabled interface{}) (int64, error) {
	var v int64
	err := p.stmt.QueryRow(disabled).Scan(&v)
	if err != nil {
		return v, err
	}

	return v, nil
}
