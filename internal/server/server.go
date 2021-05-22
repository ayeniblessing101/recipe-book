// Package server provides a method that handles every requests incoming and outcoming the recipe application
package server

import (
	"database/sql"
	"errors"
	golog "log"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

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
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	database.DBConn.SetMaxOpenConns(1)

}

// Server method handles all requests
func Server(port string) {
	engine := html.New("internal/handlers/views", ".html")
	var errMessage models.Error

	app := fiber.New(fiber.Config{
		Prefork:              false,
		ServerHeader:         "",
		StrictRouting:        false,
		CaseSensitive:        false,
		Immutable:            false,
		UnescapePath:         false,
		ETag:                 false,
		BodyLimit:            0,
		Concurrency:          0,
		Views:                engine,
		ReadTimeout:          0,
		WriteTimeout:         0,
		IdleTimeout:          0,
		ReadBufferSize:       0,
		WriteBufferSize:      0,
		CompressedFileSuffix: "",
		ProxyHeader:          "",
		GETOnly:              false,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if errors.Is(err, sql.ErrNoRows) {
				errMessage = models.Error{Message: "Page Not Found"}
				ctx.Status(404)
				return ctx.Render("404", errMessage)
			} else if err != nil {
				errMessage = models.Error{Message: "An ERROR occured please try again later"}
				zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
				log.Print("error: ", err.Error())
				return ctx.Render("404", errMessage)
			}
			return nil
		},
		DisableKeepalive:          false,
		DisableDefaultDate:        false,
		DisableDefaultContentType: false,
		DisableHeaderNormalizing:  false,
		DisableStartupMessage:     false,
		ReduceMemoryUsage:         false,
	})

	initialDatabase()
	setupRoutes(app)

	golog.Fatal(app.Listen(port))
}
