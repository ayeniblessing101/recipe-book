// Package server provides a method that handles every requests incoming and outcoming the recipe application
package server

import (
	"fmt"
	"log"
	"net/http"
)

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Hello Dima")
}

// Server method handles all requests
func Server(port string) error {
	http.HandleFunc("/", handleHelloWorld)
	log.Fatal(http.ListenAndServe(port, nil))
	
	return nil
}
