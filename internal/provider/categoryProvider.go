package provider

import (
	"database/sql"
	"log"

	"github.com/ayeniblessing101/recipe-book/internal/handlers"
	"github.com/ayeniblessing101/recipe-book/internal/models"
)

// CategoryProvider Interface show different behaviours that can be implemented by any concrete type
type CategoryProvider interface {
	CategoryGet(id int) (*models.Category, error)
	CategoryUpdate(*models.Category) error
	CategoryDelete(id int) error
}

type Provider struct {
	db *sql.DB
}

// NewProvider function create a new instance of the provider struct
func NewProvider(db *sql.DB) handlers.CategoryProvider {
	return &Provider{
		db: db,
	}
}

// CategoryGet is a provider method that get a category from the database and returns it
func (p *Provider) CategoryGet(id int) (*models.Category, error) {
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

func (p *Provider) CategoryDelete(id int) error {
	// Blessing, I had to add this in order to satisfy the interface
	return nil
}

func (p *Provider) CategoryUpdate(cat *models.Category) error {
	// Blessing, I had to add this in order to satisfy the interface
	return nil
}
