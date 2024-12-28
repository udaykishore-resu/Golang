package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"timesheet-app/models"
	"timesheet-app/services"
	"timesheet-app/utils"
)

// LoginHandler godoc
// @Summary Authenticate user and generate JWT token
// @Description Validates user credentials and returns a JWT token on successful authentication
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param credentials body models.Employee true "User credentials"
// @Success 200 {object} map[string]string "token"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 401 {object} map[string]string "Invalid username or password"
// @Router /login [post]
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
