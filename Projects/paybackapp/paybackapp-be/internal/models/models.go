package models

import "time"

type User struct {
    ID        uint      `json:"id"`
    Email     string    `json:"email"`
    Password  string    `json:"-"`
    Points    int       `json:"points"`
    CreatedAt time.Time `json:"created_at"`
}

type Restaurant struct {
    ID          uint      `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    ImageURL    string    `json:"image_url"`
    CreatedAt   time.Time `json:"created_at"`
}

type PointsTransaction struct {
    ID           uint      `json:"id"`
    UserID       uint      `json:"user_id"`
    RestaurantID uint      `json:"restaurant_id"`
    Points       int       `json:"points"`
    CreatedAt    time.Time `json:"created_at"`
}