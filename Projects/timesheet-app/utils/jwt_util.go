package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key") // Ensure this is securely managed in production

// CheckPasswordHash compares a hashed password with a plaintext password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT creates a new JWT token for the given username
func GenerateJWT(username string) (string, error) {

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Subject:   username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT checks if the provided token is valid
func ValidateJWT(tokenStr string) (bool, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}
