package database

import (
	"database/sql"
	"emp-login-service/config"
	"fmt"
	"log"
)

func ConnectDB(cfg *config.DatabaseConfig) (db *sql.DB, err error) {

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Schema)

	db, err = sql.Open("mysql", datasource)

	if err != nil {
		log.Printf("Error opening database: %v\n", err)
		return nil, err
	}

	fmt.Println("DB Connection String: ", datasource)

	if err := db.Ping(); err != nil {
		log.Println("Error pinging database....")
		return nil, err
	} else {
		log.Println("Successfully pinging the database")
	}

	fmt.Println("Database connection has been successfully established")

	return db, nil
}
