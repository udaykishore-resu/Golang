package services

import (
	"database/sql"
	"errors"
	"testing"
	"timesheet-app/models"

	"github.com/DATA-DOG/go-sqlmock"
)

// MockDBConnector is an interface to abstract database connection
type MockDBConnector interface {
	GetDB() *sql.DB
}

// mockDBContainer implements MockDBConnector
type mockDBContainer struct {
	db *sql.DB
}

func (m *mockDBContainer) GetDB() *sql.DB {
	return m.db
}

func TestGetSubprojects(t *testing.T) {
	// Create a mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	// Create a mock database connector
	mockConnector := &mockDBContainer{db: mockDB}

	testCases := []struct {
		name                 string
		projectID            int
		mockSubprojectsQuery func(*sqlmock.Rows)
		mockProjectQuery     func(*sqlmock.Rows)
		expectedResult       interface{}
		expectedError        bool
		dbConnector          MockDBConnector
	}{
		{
			name:      "Successful Subprojects Retrieval",
			projectID: 1,
			mockSubprojectsQuery: func(rows *sqlmock.Rows) {
				rows.AddRow(101, "Subproject Alpha", 1)
				rows.AddRow(102, "Subproject Beta", 1)
			},
			dbConnector: mockConnector,
			expectedResult: []models.SubProject{
				{SubProjectID: 101, SubProjectName: "Subproject Alpha", ProjectID: 1},
				{SubProjectID: 102, SubProjectName: "Subproject Beta", ProjectID: 1},
			},
			expectedError: false,
		},
		{
			name:      "No Subprojects - Project Found",
			projectID: 2,
			mockSubprojectsQuery: func(rows *sqlmock.Rows) {
				// Empty subprojects result
			},
			mockProjectQuery: func(rows *sqlmock.Rows) {
				rows.AddRow(2, "Main Project")
			},
			dbConnector: mockConnector,
			expectedResult: models.Project{
				ProjectId:   2,
				ProjectName: "Main Project",
			},
			expectedError: false,
		},
		{
			name:      "No Project Found",
			projectID: 3,
			mockSubprojectsQuery: func(rows *sqlmock.Rows) {
				// Empty subprojects result
			},
			mockProjectQuery: func(rows *sqlmock.Rows) {
				// Simulate no rows found
			},
			dbConnector:    mockConnector,
			expectedResult: nil,
			expectedError:  false,
		},
		{
			name:      "Subprojects Query Error",
			projectID: 4,
			mockSubprojectsQuery: func(rows *sqlmock.Rows) {
				mock.ExpectQuery("SELECT SubProjectID, SubProjectName, ProjectID FROM SubProjects").
					WithArgs(4).
					WillReturnError(errors.New("database error"))
			},
			dbConnector:    mockConnector,
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:      "Project Query Error",
			projectID: 5,
			mockSubprojectsQuery: func(rows *sqlmock.Rows) {
				// Empty subprojects result
			},
			mockProjectQuery: func(rows *sqlmock.Rows) {
				mock.ExpectQuery("SELECT ProjectID, ProjectName FROM Projects").
					WithArgs(5).
					WillReturnError(errors.New("database error"))
			},
			dbConnector:    mockConnector,
			expectedResult: nil,
			expectedError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Override GetDB for this test
			originalGetDB := getDBFunc
			getDBFunc = tc.dbConnector.GetDB
			defer func() { getDBFunc = originalGetDB }()

			// Prepare subprojects query expectation
			subprojectsRows := sqlmock.NewRows([]string{"SubProjectID", "SubProjectName", "ProjectID"})
			if tc.mockSubprojectsQuery != nil {
				tc.mockSubprojectsQuery(subprojectsRows)
			}

			mock.ExpectQuery("SELECT SubProjectID, SubProjectName, ProjectID FROM SubProjects").
				WithArgs(tc.projectID).
				WillReturnRows(subprojectsRows)

			// Prepare project query expectation if needed
			if tc.mockProjectQuery != nil {
				projectRows := sqlmock.NewRows([]string{"ProjectID", "ProjectName"})
				tc.mockProjectQuery(projectRows)

				mock.ExpectQuery("SELECT ProjectID, ProjectName FROM Projects").
					WithArgs(tc.projectID).
					WillReturnRows(projectRows)
			}

			// Call the function
			result, err := GetSubprojects(tc.projectID)

			// Verify error expectation
			if tc.expectedError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
				return
			}

			// Verify no unexpected errors
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Compare results
			switch expected := tc.expectedResult.(type) {
			case []models.SubProject:
				subprojects, ok := result.([]models.SubProject)
				if !ok {
					t.Errorf("Expected []models.SubProject, got %T", result)
					return
				}

				if len(subprojects) != len(expected) {
					t.Errorf("Expected %d subprojects, got %d", len(expected), len(subprojects))
					return
				}

				for i, sp := range subprojects {
					if sp.SubProjectID != expected[i].SubProjectID ||
						sp.SubProjectName != expected[i].SubProjectName ||
						sp.ProjectID != expected[i].ProjectID {
						t.Errorf("Subproject mismatch at index %d", i)
					}
				}

			case models.Project:
				project, ok := result.(models.Project)
				if !ok {
					t.Errorf("Expected models.Project, got %T", result)
					return
				}

				if project.ProjectId != expected.ProjectId ||
					project.ProjectName != expected.ProjectName {
					t.Errorf("Project details do not match")
				}

			case nil:
				if result != nil {
					t.Errorf("Expected nil result, got %v", result)
				}

			default:
				t.Errorf("Unexpected result type: %T", result)
			}

			// Ensure all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
