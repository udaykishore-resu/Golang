package controllers

import (
	"net/http"
	"strconv"
	"timesheet-app/services"
	"timesheet-app/utils"
)

// GetSubprojectsHandler godoc
// @Summary Retrieve subprojects by project ID
// @Description Fetches the list of subprojects associated with a specified project ID
// @Tags Subprojects
// @Produce json
// @Param project_id query int true "Project ID"
// @Success 200 {array} []models.Subproject "List of subprojects"
// @Failure 400 {object} map[string]string "Missing or invalid project ID"
// @Failure 500 {object} map[string]string "Could not retrieve subprojects"
// @Router /subprojects [get]

func GetSubprojectsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the project ID from the URL parameters
	projectIDStr := r.URL.Query().Get("project_id") // Assuming you send the project ID as a query parameter
	if projectIDStr == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Missing project ID")
		return
	}

	projectID, err := strconv.Atoi(projectIDStr) // Convert string to int
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid project ID")
		return
	}

	// Fetch the subprojects for the specified project ID
	subprojects, err := services.GetSubprojects(projectID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not retrieve subprojects")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, subprojects)
}
