package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"timesheet-app/models"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestLoginHandler(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Create a request body
	credentials := models.Employee{
		Username: "testuser",
		Password: "testpass",
	}
	body, _ := json.Marshal(credentials)

	// Create a new request
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Set up the expected query
	mock.ExpectQuery("SELECT password FROM Employee WHERE username = ?").
		WithArgs("testuser").
		WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow("testpass"))

	// Call the handler function
	handler := LoginHandler(db)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := `{"token":"[a-zA-Z0-9-_.]+"}` // Regex to match JWT token format
	if match, _ := regexp.MatchString(expected, rr.Body.String()); !match {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestLoginHandlerInvalidCredentials(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Create a request body with invalid credentials
	credentials := models.Employee{
		Username: "testuser",
		Password: "wrongpass",
	}
	body, _ := json.Marshal(credentials)

	// Create a new request
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Set up the expected query
	mock.ExpectQuery("SELECT password FROM Employee WHERE username = ?").
		WithArgs("testuser").
		WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow("testpass"))

	// Call the handler function
	handler := LoginHandler(db)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	// Check the response body
	expected := `{"error":"Invalid username or password"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
