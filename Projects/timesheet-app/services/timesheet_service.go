package services

import (
	"fmt"
	"timesheet-app/database"
	"timesheet-app/models"
)

func SubmitTimesheet(timesheet models.TimesheetDetail) error {
	db := database.GetDB()
	fmt.Println("args: ", timesheet)
	_, err := db.Exec("INSERT INTO Timesheets (ProjectID,SubProjectID, JiraSnowID, TaskDescription, HoursSpent, Comments) VALUES (?, ?,?,?,?,?)", timesheet.ProjectID, timesheet.SubProjectID, timesheet.JiraSnowID, timesheet.TaskDescription, timesheet.HoursSpent, timesheet.Comments)
	return err
}
