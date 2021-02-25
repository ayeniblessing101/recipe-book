package provider

import (
	"database/sql"

	"github.com/ayeniblessing101/recipe-book/internal/models"
)

type Provider struct {
	db *sql.DB
}

// NewProvider function create a new instance of the provider struct
func NewProvider(db *sql.DB) *Provider {
	return &Provider{
		db: db,
	}
}

// CategoryGet is a provider method that get a category from the database and returns it
func (p *Provider) CategoryGet(id int) (*models.Category, error) {
	row := p.db.QueryRow("SELECT * FROM categories WHERE id = $id")

	var category *models.Category
	if err := row.Scan(&category.ID, &category.Name); err != nil {
		return nil, err
	}

	return category, nil
}
