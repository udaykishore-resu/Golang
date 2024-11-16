package services

import (
	"database/sql"
	"emp-login-service/database"
	"emp-login-service/utils"
	"errors"
	"fmt"
)

func Authenticate(db *sql.DB, jwt string, username, password string) (token string, err error) {
	var store_pass string
	if err := db.QueryRow(database.LoginQueries.GetUserPassword, username).Scan(&store_pass); err != nil {
		fmt.Println("Invalid User")
		return "", err
	}

	if store_pass != password {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateJWTToken(username, jwt)
}
