package handlers

import (
	"github.com/ayeniblessing101/recipe-book/internal/database"
	"github.com/ayeniblessing101/recipe-book/internal/models"
	"github.com/gofiber/fiber/v2"
)

// AddCategory method adds category to the categories table
func AddCategory(c *fiber.Ctx) error {
	cat := new(models.Category)
	db := database.DBConn
	stmt, createTableError := db.Prepare("CREATE TABLE IF NOT EXISTS categories (id INTEGER PRIMARY KEY, name TEXT)")

	if createTableError != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error": createTableError,
		})
	}
	stmt.Exec()

	if err := c.BodyParser(cat); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error": err,
		})
	}

	stmt, insertingError := database.DBConn.Prepare("INSERT INTO categories (name) VALUES (?)")

	defer stmt.Close()

	if insertingError != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error": insertingError,
		})
	}	
		if _, err := stmt.Exec(cat.Name); err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error": err,
			})
		}
	return c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "Category created successfully",
		"category": cat,
	})
}

// GetCategories method retrieves all categories from the categories table
func GetCategories(c *fiber.Ctx) error {
	db := database.DBConn
	
	rows, err := db.Query("SELECT * FROM categories")

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error": err,
		})
	}

	var category models.Category
	var categories []models.Category

	for rows.Next() {
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error": err, 
			})
		}
		categories = append(categories, category)
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"categories": categories,
	})
}
