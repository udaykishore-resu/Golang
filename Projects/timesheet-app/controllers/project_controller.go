package controllers

import (
	"net/http"
	"timesheet-app/services"
	"timesheet-app/utils"
)

func GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := services.GetProjects()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not retrieve projects")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, projects)
}
