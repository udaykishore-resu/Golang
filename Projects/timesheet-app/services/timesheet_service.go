package services

import (
	"timesheet-app/database"
	"timesheet-app/models"
)

func SubmitTimesheet(timesheet models.TimesheetDetail) error {
	db := database.GetDB()
	_, err := db.Exec("INSERT INTO Timesheets (EmployeeID,ProjectID,SubProjectID, JiraSnowID, TaskDescription, HoursSpent, Comments) VALUES (?,?, ?,?,?,?,?)", timesheet.EmployeeID, timesheet.ProjectID, timesheet.SubProjectID, timesheet.JiraSnowID, timesheet.TaskDescription, timesheet.HoursSpent, timesheet.Comments)
	return err
}
