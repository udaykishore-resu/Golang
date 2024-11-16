package main

import (
	"emp-login-service/config"
	"emp-login-service/database"
	"emp-login-service/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.ConnectDB(&cfg.Database)
	if err != nil {
		log.Println("Database connection failed")
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Println("db connection close failed")
		}
	}()

	r := mux.NewRouter()
	routes.RegisterLoginRoutes(r, db)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(r)

	server := &http.Server{
		Handler:      handler,
		Addr:         cfg.Serverport,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Starting User Login Service on port %s...", cfg.Serverport)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("User Login Service failed to start: %v", err)
	}
}
