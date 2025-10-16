package auth

import (
	"time"
	"fmt"

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

func ValidateToken(tokenString, secretKey string) (*jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(secretKey), nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return &claims, nil
    }

    return nil, fmt.Errorf("invalid token")
}