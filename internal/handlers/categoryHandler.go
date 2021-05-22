package handlers

import (
	"strconv"

	"github.com/ayeniblessing101/recipe-book/internal/database"
	"github.com/ayeniblessing101/recipe-book/internal/models"
	"github.com/ayeniblessing101/recipe-book/internal/providers"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// AddCategory method adds category to the categories table
func AddCategory(c *fiber.Ctx) error {
	cat := new(models.Category)
	db := database.DBConn
	stmt, createTableError := db.Prepare("CREATE TABLE IF NOT EXISTS categories (id INTEGER PRIMARY KEY, name TEXT)")
  zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if createTableError != nil {
		log.Error().Err(createTableError).Msg("")
	}
	stmt.Exec()

	if err := c.BodyParser(cat); err != nil {
		return err
	}

	stmt, insertingError := database.DBConn.Prepare("INSERT INTO categories (name) VALUES (?)")

	defer stmt.Close()

	if insertingError != nil {
		log.Error().Err(insertingError).Msg("")
	}
	if _, err := stmt.Exec(cat.Name); err != nil {
		log.Error().Err(err).Msg("")
	}
	return c.Status(201).SendString("Category created successfully")
}

// GetCategories method retrieves all categories from the categories table
func GetCategories(c *fiber.Ctx) error {
	db := database.DBConn

	rows, err := db.Query("SELECT * FROM categories")

	if err != nil {
		log.Error().Err(err).Msg("")
	}

	var category models.Category
	var categories []models.Category

	for rows.Next() {
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			log.Error().Err(err).Msg("")
		}
		categories = append(categories, category)
	}
	return c.Render("category", categories)
}

// GetCategory method retrieves a category from the categories table
func GetCategory(p providers.CategoryProvider) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		categoryID, _ := strconv.Atoi(c.Params("id"))
		category, err := p.CategoryGet(categoryID)

		if err != nil {
			log.Error().Err(err).Msg("")
		}

		// wrapped into array to it works in the template
		result := make([]*models.Category, 0)
		result = append(result, category)
		return c.Render("category", result)
	}
}

// UpdateCategory method update a category from the categories table and returns an error
func UpdateCategory(p providers.CategoryProvider) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {  
		cat := new(models.Category)

		if err := c.BodyParser(cat); err != nil {
			log.Error().Err(err).Msg("")
		}
		categoryID, _ := strconv.Atoi(c.Params("id"))
		err := p.CategoryUpdate(cat, categoryID)

		if err != nil {
			log.Error().Err(err).Msg("")
		}
		result := make([]*models.Category, 0)
		return c.Render("category", result)
	}
}

// DeleteCategory method delete a category from the categories table and returns an error
func DeleteCategory(p providers.CategoryProvider) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		categoryID, _ := strconv.Atoi(c.Params("id"))
		err := p.CategoryDelete(categoryID)

		if err != nil {
			log.Error().Err(err).Msg("")
		}

		result := make([]*models.Category, 0)
		return c.Render("category", result)
	}
}
