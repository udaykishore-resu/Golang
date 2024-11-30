package services

import (
	"database/sql"
	"errors"
	"timesheet-app/utils"
)

// Authenticate checks the username and password against the database and returns a JWT token if valid.
func Authenticate(db *sql.DB, username, password string) (string, error) {

	// Query the database to validate user credentials
	var storedPassword string
	query := "SELECT password FROM Employee WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&storedPassword)
	if err != nil {
		return "", errors.New("user not found or invalid credentials")
	}

	// Check if the provided password matches the stored password
	if storedPassword != password {
		return "", errors.New("invalid credentials")
	}

	// Generate a JWT token for the authenticated user
	return utils.GenerateJWT(username)
}
