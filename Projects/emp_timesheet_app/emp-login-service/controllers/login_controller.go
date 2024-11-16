package controllers

import (
	"database/sql"
	"emp-login-service/config"
	"emp-login-service/models"
	"emp-login-service/services"
	"emp-login-service/utils"
	"encoding/json"
	"net/http"
)

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var emp models.Employee

		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			utils.RespondWithError(rw, http.StatusBadRequest, "Invalid Request Payload")
			return
		}
		cfg, err := config.LoadConfig()

		token, err := services.Authenticate(db, cfg.JWTSecret, emp.Username, emp.Password)

		if err != nil {
			utils.RespondWithError(rw, http.StatusUnauthorized, "Invalid Username and Password")
			return
		}
		utils.RespondWithJson(rw, http.StatusOK, map[string]string{"token": token})
	}
}
