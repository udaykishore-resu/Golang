package routes

import (
	"database/sql"
	"net/http"
	"timesheet-app/controllers"
	"timesheet-app/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, db *sql.DB) {

	r.HandleFunc("/login", controllers.LoginHandler(db)).Methods("POST")
	r.Handle("/projects", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetProjectsHandler))).Methods("GET")
	r.Handle("/subprojects", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetSubprojectsHandler))).Methods("GET")
	r.Handle("/timesheet", middleware.AuthMiddleware(http.HandlerFunc(controllers.SubmitTimesheetHandler))).Methods("POST")
}
