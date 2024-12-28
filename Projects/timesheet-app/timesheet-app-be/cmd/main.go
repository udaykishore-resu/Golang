package main

import (
	"log"
	"net/http"
	"time"
	"timesheet-app/config"
	"timesheet-app/database"
	"timesheet-app/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// @title Timesheet API
// @version 1.0
// @description API for managing timesheets
// @contact.name Anusha Cheerla
// @contact.email anu123.cheerla@gmail.com
// @license.name MIT
// @host localhost:8080
// @BasePath /
func main() {

	// Load application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize the database connection
	db, err := database.ConnectDB(&cfg.Database) // Make sure this function is correctly defined
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing the database connection: %v", err)
		}
	}()

	// Set up the router
	r := mux.NewRouter()
	routes.RegisterRoutes(r, db) // Ensure this function correctly registers your routes

	// Create a CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Add your React app's URL
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Wrap the router with the CORS handler
	handler := c.Handler(r)

	// Set up the HTTP server
	server := &http.Server{
		Handler:      handler,
		Addr:         cfg.ServerPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start the server
	log.Printf("Starting server on port %s...", cfg.ServerPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
