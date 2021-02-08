package database

import (
	"database/sql"

	// importing the go-sqlite3 driver in the database package
	_ "github.com/mattn/go-sqlite3"
)

// DBConn global var to connect to database
var (
	DBConn *sql.DB
)
