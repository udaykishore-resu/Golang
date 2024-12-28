package services

import (
	"timesheet-app/database"
	"timesheet-app/models"
)

// GetProjects retrieves a list of projects from the database
func GetProjects() ([]models.Project, error) {
	db := database.GetDB() // Get the database connection
	var projects []models.Project

	// Query to fetch projects from the database
	query := "SELECT ProjectID, ProjectName FROM Projects"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err // Return nil and the error if the query fails
	}
	defer rows.Close() // Ensure rows are closed after processing

	for rows.Next() {
		var project models.Project
		// Scan the row into the project struct
		if err := rows.Scan(&project.ProjectId, &project.ProjectName); err != nil {
			return nil, err // Return nil and the error if scanning fails
		}
		projects = append(projects, project) // Append the project to the slice
	}

	// Check for any error encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err // Return nil and the error if there was an issue
	}

	return projects, nil // Return the list of projects and a nil error
}
