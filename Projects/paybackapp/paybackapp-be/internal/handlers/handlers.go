package handlers

import (
	"net/http"
	"paybackapp/internal/database"
	"paybackapp/internal/models"

	"github.com/gin-gonic/gin"
)

func GetRestaurants(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, name, description, image_url, created_at FROM restaurants")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var restaurants []models.Restaurant
	for rows.Next() {
		var r models.Restaurant
		if err := rows.Scan(&r.ID, &r.Name, &r.Description, &r.ImageURL, &r.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		restaurants = append(restaurants, r)
	}

	c.JSON(http.StatusOK, restaurants)
}

func CreateRestaurant(c *gin.Context) {
	var restaurant models.Restaurant
	if err := c.ShouldBindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := database.DB.Exec(
		"INSERT INTO restaurants (name, description, image_url) VALUES (?, ?, ?)",
		restaurant.Name, restaurant.Description, restaurant.ImageURL,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	restaurant.ID = uint(id)
	c.JSON(http.StatusCreated, restaurant)
}

func GetRestaurant(c *gin.Context) {
	id := c.Param("id")
	var restaurant models.Restaurant
	err := database.DB.QueryRow(
		"SELECT id, name, description, image_url, created_at FROM restaurants WHERE id = ?",
		id,
	).Scan(&restaurant.ID, &restaurant.Name, &restaurant.Description, &restaurant.ImageURL, &restaurant.CreatedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	c.JSON(http.StatusOK, restaurant)
}

func GetUserPoints(c *gin.Context) {
	userId := c.Param("userId")
	var points int
	err := database.DB.QueryRow("SELECT points FROM users WHERE id = ?", userId).Scan(&points)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}

func AddPoints(c *gin.Context) {
	var transaction models.PointsTransaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := database.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Insert transaction
	_, err = tx.Exec(
		"INSERT INTO points_transactions (user_id, restaurant_id, points) VALUES (?, ?, ?)",
		transaction.UserID, transaction.RestaurantID, transaction.Points,
	)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update user points
	_, err = tx.Exec(
		"UPDATE users SET points = points + ? WHERE id = ?",
		transaction.Points, transaction.UserID,
	)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In a real application, you would hash the password here
	result, err := database.DB.Exec(
		"INSERT INTO users (email, password) VALUES (?, ?)",
		user.Email, user.Password,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	user.ID = uint(id)
	user.Password = "" // Don't send password back
	c.JSON(http.StatusCreated, user)
}

func LoginUser(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := database.DB.QueryRow(
		"SELECT id, email FROM users WHERE email = ? AND password = ?",
		credentials.Email, credentials.Password,
	).Scan(&user.ID, &user.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// In a real application, you would generate and return a JWT token here
	c.JSON(http.StatusOK, user)
}
