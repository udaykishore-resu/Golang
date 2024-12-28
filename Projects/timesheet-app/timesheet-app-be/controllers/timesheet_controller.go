package controllers

import (
	"encoding/json"
	"net/http"
	"timesheet-app/models"
	"timesheet-app/services"
	"timesheet-app/utils"
)

// SubmitTimesheetHandler godoc
// @Summary Submit a new timesheet
// @Description Receives and processes a timesheet submission, including validation and storage
// @Tags Timesheets
// @Accept json
// @Produce json
// @Param timesheet body models.TimesheetDetail true "Timesheet details"
// @Success 200 {object} map[string]string "Timesheet submitted successfully"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Could not save timesheet"
// @Router /timesheet [post]
func SubmitTimesheetHandler(w http.ResponseWriter, r *http.Request) {
	var timesheet models.TimesheetDetail
	if err := json.NewDecoder(r.Body).Decode(&timesheet); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := services.SubmitTimesheet(timesheet); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Could not save timesheet")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Timesheet submitted successfully"})
}
