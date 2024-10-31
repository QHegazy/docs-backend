package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	currentToken    string
	tokenExpiration time.Time
)

func GenerateJWTToken(userId string) (string, error) {
	claims := &jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().UTC().Add(time.Hour * 24).Unix(),
	}
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func GetJWTToken(userID string) (string, error) {
	if currentToken == "" || time.Now().UTC().After(tokenExpiration) {
		token, err := GenerateJWTToken(userID)
		if err != nil {
			return "", err
		}
		currentToken = token
		tokenExpiration = time.Now().UTC().Add(72 * time.Hour)
	}
	return currentToken, nil
}
