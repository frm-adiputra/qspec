package model

//go:generate go-bindata -o assets/assets.go -pkg assets -prefix _assets _assets/
//go:generate qspec blog.yaml

import (
	"database/sql"

	// Import SQLite3 driver.
	_ "github.com/mattn/go-sqlite3"

	"github.com/frm-adiputra/qspec/examples/model/assets"
	"github.com/frm-adiputra/qspec/examples/model/blog"
)

var db *sql.DB

// OpenDB opens database and sets all models database.
func OpenDB(dataSource string) error {
	_db, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		return err
	}

	db = _db
	blog.DB = _db
	return nil
}

// CloseDB closes the database.
func CloseDB() error {
	err := db.Close()
	if err != nil {
		return err
	}

	db = nil
	blog.DB = nil
	return nil
}

// InitDB initializes database.
func InitDB() error {
	sql, err := assets.Asset("tables.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sql))
	if err != nil {
		return err
	}
	return nil
}
