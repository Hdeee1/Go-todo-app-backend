package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateToken(secretKey string, userID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id":	userID,
		"exp":		time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
} 