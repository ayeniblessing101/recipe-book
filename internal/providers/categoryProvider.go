package providers

import (
	"database/sql"
	"log"

	"github.com/ayeniblessing101/recipe-book/internal/models"
)

// CategoryProvider Interface show different behaviours that can be implemented by any concrete type
type CategoryProvider interface {
	CategoryGet(id int) (*models.Category, error)
	CategoryUpdate(*models.Category) error
	CategoryDelete(id int) error
}

type provider struct {
	db *sql.DB
}

// NewProvider function create a new instance of the provider struct
func NewProvider(db *sql.DB) CategoryProvider {
	return &provider{
		db: db,
	}
}

// CategoryGet is a provider method that get a category from the database and returns it
func (p *provider) CategoryGet(id int) (*models.Category, error) {
	if p.db == nil {
		panic("Blessing, db is nil, it means it did not initialize")
	}

	row := p.db.QueryRow("SELECT id, name FROM categories WHERE id=?", id)

	// Blessing, don't write to uninitialized objects
	category := &models.Category{}
	if err := row.Scan(&category.ID, &category.Name); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return category, nil
}

// CategoryDelete is a provider method that delete a category from the database and returns an error if any
func (p *provider) CategoryDelete(id int) error {
	if p.db == nil {
		panic("db is nil")
	}

	stmt, err := p.db.Prepare("DELETE FROM categories WHERE id=?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
    
	// Hi Dima do I need to check for the rows affected
	// affect, err := res.RowsAffected()

	return nil
}

// CategoryUpdate is a provider method that update a category from the database and return it
func (p *provider) CategoryUpdate(cat *models.Category) error {
	if p.db == nil {
		panic("db is nil")
	}

	stmt,err := p.db.Prepare("UPDATE categories SET name=? WHERE id=?")

	return nil
}
