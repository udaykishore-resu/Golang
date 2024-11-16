package routes

import (
	"database/sql"
	"emp-login-service/controllers"

	"github.com/gorilla/mux"
)

func RegisterLoginRoutes(r *mux.Router, db *sql.DB) {
	r.Handle("/login", controllers.LoginHandler(db)).Methods("POST")
}
