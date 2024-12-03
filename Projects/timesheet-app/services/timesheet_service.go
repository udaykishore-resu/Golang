package services

import (
	"fmt"
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
	fmt.Println("query: ", query)
	fmt.Println("args: ", args)
	return err
}
