package services

import (
	"database/sql"
	"timesheet-app/database"
	"timesheet-app/models"
)

// getDBFunc is a variable that can be replaced for testing
var getDBFunc = database.GetDB

// GetSubprojects retrieves a list of subprojects for a given project ID from the database
// If no subprojects are found, it returns the project details instead
func GetSubprojects(projectID int) (interface{}, error) {
	db := getDBFunc() // Get the database connection

	// Query to fetch subprojects based on the provided project ID
	subprojectsQuery := "SELECT SubProjectID, SubProjectName, ProjectID FROM SubProjects WHERE ProjectID = ?"
	rows, err := db.Query(subprojectsQuery, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subprojects []models.SubProject
	for rows.Next() {
		var subproject models.SubProject
		if err := rows.Scan(&subproject.SubProjectID, &subproject.SubProjectName, &subproject.ProjectID); err != nil {
			return nil, err
		}
		subprojects = append(subprojects, subproject)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// If subprojects are found, return them
	if len(subprojects) > 0 {
		return subprojects, nil
	}

	// If no subprojects are found, query for project details
	projectQuery := "SELECT ProjectID, ProjectName FROM Projects WHERE ProjectID = ?"
	var project models.Project
	err = db.QueryRow(projectQuery, projectID).Scan(&project.ProjectId, &project.ProjectName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No project found
		}
		return nil, err
	}

	return project, nil
}
