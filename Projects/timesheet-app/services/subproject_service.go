package services

import (
	"timesheet-app/database"
	"timesheet-app/models"
)

// GetSubprojects retrieves a list of subprojects for a given project ID from the database
func GetSubprojects(projectID int) ([]models.SubProject, error) {
	db := database.GetDB() // Get the database connection
	var subprojects []models.SubProject

	// Query to fetch subprojects based on the provided project ID
	query := "SELECT SubProjectID, SubProjectName, ProjectID FROM SubProjects WHERE ProjectID = ?"
	rows, err := db.Query(query, projectID)
	if err != nil {
		return nil, err // Return nil and the error if the query fails
	}
	defer rows.Close() // Ensure rows are closed after processing

	for rows.Next() {
		var subproject models.SubProject
		// Scan the row into the subproject struct
		if err := rows.Scan(&subproject.SubProjectID, &subproject.SubProjectName, &subproject.ProjectID); err != nil {
			return nil, err // Return nil and the error if scanning fails
		}
		subprojects = append(subprojects, subproject) // Append the subproject to the slice
	}

	// Check for any error encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err // Return nil and the error if there was an issue
	}

	return subprojects, nil // Return the list of subprojects and a nil error
}
