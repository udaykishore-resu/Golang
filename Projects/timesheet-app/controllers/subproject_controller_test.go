package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock types and functions to avoid import cycles
type MockSubprojectService struct {
	GetSubprojectsFunc func(projectID int) ([]interface{}, error)
}

func (m *MockSubprojectService) GetSubprojects(projectID int) ([]interface{}, error) {
	if m.GetSubprojectsFunc != nil {
		return m.GetSubprojectsFunc(projectID)
	}
	return nil, nil
}

// Mock response utilities
func mockRespondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}

func mockRespondWithJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func TestGetSubprojectsHandler(t *testing.T) {
	testCases := []struct {
		name               string
		projectIDParam     string
		mockGetSubprojects func(projectID int) ([]interface{}, error)
		expectedStatusCode int
		expectedBody       []interface{}
		expectError        bool
	}{
		{
			name:           "Successful Retrieval",
			projectIDParam: "1",
			mockGetSubprojects: func(projectID int) ([]interface{}, error) {
				return []interface{}{
					map[string]interface{}{"id": 1, "name": "Subproject 1"},
				}, nil
			},
			expectedStatusCode: http.StatusOK,
			expectedBody: []interface{}{
				map[string]interface{}{"id": 1, "name": "Subproject 1"},
			},
		},
		{
			name:               "Missing Project ID",
			projectIDParam:     "",
			expectedStatusCode: http.StatusBadRequest,
			expectError:        true,
		},
		{
			name:               "Invalid Project ID",
			projectIDParam:     "invalid",
			expectedStatusCode: http.StatusBadRequest,
			expectError:        true,
		},
		{
			name:           "Service Error",
			projectIDParam: "2",
			mockGetSubprojects: func(projectID int) ([]interface{}, error) {
				return nil, errors.New("service error")
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectError:        true,
		},
		{
			name:           "Empty Subprojects List",
			projectIDParam: "3",
			mockGetSubprojects: func(projectID int) ([]interface{}, error) {
				return []interface{}{}, nil
			},
			expectedStatusCode: http.StatusOK,
			expectedBody:       []interface{}{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create request
			req, err := http.NewRequest("GET", "/subprojects", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Add query parameter
			if tc.projectIDParam != "" {
				q := req.URL.Query()
				q.Add("project_id", tc.projectIDParam)
				req.URL.RawQuery = q.Encode()
			}

			// Create response recorder
			w := httptest.NewRecorder()

			// Create mock service
			mockService := &MockSubprojectService{
				GetSubprojectsFunc: tc.mockGetSubprojects,
			}

			// Create handler
			handler := &SubprojectHandler{
				GetSubprojectsFunc: mockService.GetSubprojects,
				RespondWithError:   mockRespondWithError,
				RespondWithJSON:    mockRespondWithJSON,
			}

			// Call handler
			handler.GetSubprojectsHandler(w, req)

			// Get response
			response := w.Result()

			// Check status code
			if response.StatusCode != tc.expectedStatusCode {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatusCode, response.StatusCode)
			}

			// If expecting successful response, check body
			if !tc.expectError && tc.expectedStatusCode == http.StatusOK {
				var responseBody []interface{}
				err := json.NewDecoder(response.Body).Decode(&responseBody)
				if err != nil {
					t.Fatalf("Failed to decode response body: %v", err)
				}

				// Compare response body
				if len(responseBody) != len(tc.expectedBody) {
					t.Errorf("Expected body length %d, got %d", len(tc.expectedBody), len(responseBody))
				}
			}
		})
	}
}
