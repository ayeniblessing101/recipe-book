// Package server provides a method that handles every requests incoming and outcoming the recipe application
package server

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/ayeniblessing101/recipe-book/internal/models"
	"github.com/davecgh/go-spew/spew"
)

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	recipe := &models.Recipe{
		ID: 23,
		Name: "Singaporean Noodles",
		Ingredients: []string{"Rice Noodle", "Red Bell Pepper", "Green Bell Pepper", "Onion", "Garlic"},
		Directions: []string{"Boil Water", "Soak Noodles in Hot Water", "Fry Vegetables"},
		Calories: 23,
	}
	fmt.Fprintf(w, "%s\n", "Hello Dima")
	fmt.Fprintf(w, html.EscapeString(spew.Sdump(recipe)))

}

// Server method handles all requests
func Server(port string) {
	http.HandleFunc("/", handleHelloWorld)
	log.Fatal(http.ListenAndServe(port, nil))
}
