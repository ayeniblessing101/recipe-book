// Package server provides a method that handles every requests incoming and outcoming the recipe application
package server

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ayeniblessing101/recipe-book/internal/database"
	"github.com/ayeniblessing101/recipe-book/internal/handlers"
	"github.com/ayeniblessing101/recipe-book/internal/models"
	"github.com/ayeniblessing101/recipe-book/internal/providers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func setupRoutes(app *fiber.App) {
	p := providers.NewProvider(database.DBConn)

	app.Get("/categories", handlers.GetCategories)
	app.Get("/categories/:id", handlers.GetCategory(p))
	app.Post("/categories", handlers.AddCategory)
	app.Patch("/categories/:id", handlers.UpdateCategory(p))
	app.Delete("/categories/:id", handlers.DeleteCategory(p))
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
	var errMessage models.Error

	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if errors.Is(err, sql.ErrNoRows) {
				errMessage = models.Error{Message: "Page Not Found"}
				ctx.Status(404)
				return ctx.Render("404", errMessage)
			} else if err != nil {
				errMessage = models.Error{Message: "An ERROR occured please try again later"}
				log.Println("error: ", err.Error())
				return ctx.Render("404", errMessage)
			}
			return nil
		},
	})

	initialDatabase()
	setupRoutes(app)

	log.Fatal(app.Listen(port))
}
