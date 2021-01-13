package database

import (
	"database/sql"
	"log"

	// importing the go-sqlite3 driver in the database package
	_ "github.com/mattn/go-sqlite3"
)

// ConnectToDatabase method coonects to the sqlite database and returns the database and error
func ConnectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./recipe.db")

	if err != nil {
		log.Fatal("An error occured when connecting to database", err)
	}

	db.SetMaxOpenConns(1)

	return db, nil
}