package controllers

import (
	"fmt"
	"net/http"
	"timesheet-app/services"
	"timesheet-app/utils"
)

func GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetProjectsHandler:::Entry")
	projects, err := services.GetProjects()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not retrieve projects")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, projects)
}
