// Package server provides a method that handles every requests incoming and outcoming the recipe application
package server

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/ayeniblessing101/recipe-book/internal/handlers"
	"github.com/davecgh/go-spew/spew"
)

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", "Hello Dima")
	fmt.Fprintf(w, html.EscapeString(spew.Sdump()))
}

// Server method handles all requests
func Server(port string) {
	http.HandleFunc("/", handleHelloWorld)
	http.HandleFunc("/categories", handlers.AddCategory)
	http.HandleFunc("/categories", handlers.GetCategories)
	log.Fatal(http.ListenAndServe(port, nil))
}
