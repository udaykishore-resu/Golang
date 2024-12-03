package services

import (
	"timesheet-app/database"
	"timesheet-app/models"
)

func SubmitTimesheet(timesheet models.TimesheetDetail) error {
	db := database.GetDB()
	query := "INSERT INTO Timesheets (ProjectID,SubProjectID, JiraSnowID, TaskDescription, HoursSpent, Comments) VALUES (?,?, ?,?,?,?)"
	args := []interface{}{
		timesheet.ProjectID,
		timesheet.SubProjectID,
		timesheet.JiraSnowID,
		timesheet.TaskDescription,
		timesheet.HoursSpent,
		timesheet.Comments,
	}
	_, err := db.Exec(query, args...)
	return err
}
