package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWTToken(user string, jwtToken string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Subject:   user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtToken)
}

func ValidateJWT(tokenStr, jwtSecret string) (bool, error) {
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
