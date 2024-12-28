package controllers

import (
	"net/http"
	"timesheet-app/services"
	"timesheet-app/utils"
)

// GetProjectsHandler godoc
// @Summary Retrieve all projects
// @Description Fetches the list of projects from the service
// @Tags Projects
// @Produce json
// @Success 200 {array} []models.Project "List of projects"
// @Failure 500 {object} map[string]string "Could not retrieve projects"
// @Router /projects [get]
func GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := services.GetProjects()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not retrieve projects")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, projects)
}
