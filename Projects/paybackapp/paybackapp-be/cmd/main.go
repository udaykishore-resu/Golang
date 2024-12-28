package main

import (
	"log"
	"paybackapp/internal/database"
	"paybackapp/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	if err := database.InitDB(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Create router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Routes
	api := r.Group("/api")
	{
		// Restaurant routes
		api.GET("/restaurants", handlers.GetRestaurants)
		api.POST("/restaurants", handlers.CreateRestaurant)
		api.GET("/restaurants/:id", handlers.GetRestaurant)

		// User routes
		api.POST("/register", handlers.RegisterUser)
		api.POST("/login", handlers.LoginUser)

		// Points routes
		api.GET("/points/:userId", handlers.GetUserPoints)
		api.POST("/points", handlers.AddPoints)
	}

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
