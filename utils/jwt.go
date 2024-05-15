package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "ASDASDSAD"

func GenerateToken(email string, userID int64) (string, error) {
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString([]byte(secret))
}
