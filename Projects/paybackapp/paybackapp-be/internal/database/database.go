package database

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() error {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        return fmt.Errorf("error loading .env file: %v", err)
    }

    // Get database connection details from environment variables
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbHost := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")

    // Create connection string
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbName)

    // Open database connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return fmt.Errorf("error connecting to database: %v", err)
    }

    // Test connection
    if err := db.Ping(); err != nil {
        return fmt.Errorf("error pinging database: %v", err)
    }

    DB = db
    return nil
}