// Package server provides a method that handles every requests incoming and outcoming the recipe application
package server

import (
	"database/sql"
	"log"

	"github.com/ayeniblessing101/recipe-book/internal/database"
	"github.com/ayeniblessing101/recipe-book/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func setupRoutes(app *fiber.App) {
	app.Get("/categories", handlers.GetCategories)
	app.Post("/categories", handlers.AddCategory)
}

func initialDatabase() {
	var err error
	database.DBConn, err = sql.Open("sqlite3", "./recipe.db")

	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	database.DBConn.SetMaxOpenConns(1)
}

// Server method handles all requests
func Server(port string) {
	engine := html.New("internal/handlers/views", ".html")
	
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	initialDatabase()

	setupRoutes(app)

	log.Fatal(app.Listen(port))
}
