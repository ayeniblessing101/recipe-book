package handlers

import (
	"database/sql"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/ayeniblessing101/recipe-book/internal/database"
	"github.com/ayeniblessing101/recipe-book/internal/models"
	"github.com/davecgh/go-spew/spew"
)

var db *sql.DB
var err error

func init() {
	db, err = database.ConnectToDatabase()

	if err != nil {
		log.Fatal("An error occured when connecting to database", err)
	}
}

// AddCategory method adds category to the categories table
func AddCategory(w http.ResponseWriter, r *http.Request) {
	stmt, createTableError := db.Prepare("CREATE TABLE IF NOT EXISTS categories (id INTEGER PRIMARY KEY, name TEXT)")

	if createTableError != nil {
		log.Fatal("An error occured creating the table: ", createTableError)
	}
	stmt.Exec()

	stmt, insertingError := db.Prepare("INSERT INTO categories (name) VALUES (?)")

	defer stmt.Close()

	if insertingError != nil {
		log.Fatal("An error occured inserting into the table : ", insertingError)
	}

	if _, err := stmt.Exec("Vegetables"); err != nil {
		log.Fatal("An error occured", err)
	}
}

// GetCategories method retrieves all categories from the categories table
func GetCategories(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM categories")

	if err != nil {
		log.Fatal("An error occured querying query database", err)
	}

	var category models.Category
	for rows.Next() {
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, html.EscapeString(spew.Sdump(category.ID, category.Name)))
	}
}
