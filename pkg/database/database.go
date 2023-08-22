package database

import (
	"database/sql"
	// _ "github.com/lib/pq"  // for postegre sql
	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "/home/frap/Desktop/makerspace-api/database.db")
	if err != nil {
		return nil, err
	}

	// Return the database connection
	return db, nil
}
