package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"timesheet-app/models"
	"timesheet-app/services"
	"timesheet-app/utils"
)

// LoginHandler authenticates a user and returns a JWT token if successful
func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials models.Employee
		if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		token, err := services.Authenticate(db, credentials.Username, credentials.Password)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid username or password")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
	}
}
