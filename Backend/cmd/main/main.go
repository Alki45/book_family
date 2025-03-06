package main

import (
	"book_family/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	// CORS settings
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allows all origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Register the routes for books
	routes.RegisterBookStoreRoutes(r)

	// Register the routes for users
	routes.RegisterUserRoutes(r)

	handler := c.Handler(r)

	// Start the server
	log.Fatal(http.ListenAndServe("localhost:9010", handler))
}
