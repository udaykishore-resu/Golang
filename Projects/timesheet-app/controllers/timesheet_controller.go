package controllers

import (
	"encoding/json"
	"net/http"
	"timesheet-app/models"
	"timesheet-app/services"
	"timesheet-app/utils"
)

func SubmitTimesheetHandler(w http.ResponseWriter, r *http.Request) {
	var timesheet models.TimesheetDetail
	if err := json.NewDecoder(r.Body).Decode(&timesheet); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Process the timesheet, including the TimesheetJSON
	// You might want to parse the TimesheetJSON and validate it here

	if err := services.SubmitTimesheet(timesheet); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not save timesheet")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Timesheet submitted successfully"})
}
