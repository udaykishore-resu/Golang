package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAuthenticate(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Set up the expected query
	mock.ExpectQuery("SELECT password FROM Employee WHERE username = ?").
		WithArgs("testuser").
		WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow("testpass"))

	// Call the Authenticate function
	token, err := Authenticate(db, "testuser", "testpass")

	// Check for errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if a token was generated
	if token == "" {
		t.Error("Expected a token, got an empty string")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestAuthenticateInvalidCredentials(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Set up the expected query
	mock.ExpectQuery("SELECT password FROM Employee WHERE username = ?").
		WithArgs("testuser").
		WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow("testpass"))

	// Call the Authenticate function with wrong password
	_, err = Authenticate(db, "testuser", "wrongpass")

	// Check for expected error
	if err == nil {
		t.Error("Expected an error, got nil")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestAuthenticateUserNotFound(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Set up the expected query to return no rows
	mock.ExpectQuery("SELECT password FROM Employee WHERE username = ?").
		WithArgs("nonexistentuser").
		WillReturnRows(sqlmock.NewRows([]string{"password"}))

	// Call the Authenticate function with a non-existent user
	_, err = Authenticate(db, "nonexistentuser", "anypassword")

	// Check for expected error
	if err == nil {
		t.Error("Expected an error, got nil")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
