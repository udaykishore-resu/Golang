package database

import (
	"database/sql"
	"fmt"
	"log"
	"timesheet-app/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// ConnectDB initializes the database connection with the given config
func ConnectDB(cfg *config.DatabaseConfig) (*sql.DB, error) {

	// Use values from the config
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Printf("Error opening database: %v\n", err)
		return nil, err
	}
	// Test the connection
	if err := db.Ping(); err != nil {
		log.Printf("Error pinging database: %v\n", err)
		return nil, err
	} else {
		log.Printf("pinging database:...")
	}

	log.Println("Database connection successfully established.")
	return db, nil
}

// GetDB returns a reference to the database object
func GetDB() *sql.DB {
	return db
}

// CloseDB closes the database connection
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
