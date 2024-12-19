package api

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/raihan-faza/learning-go/initializers"
)

func init() {
	initializers.LoadEnv()
}

func generateToken(username string, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 12).Unix(),
	})
	secretKey := os.Getenv("secret")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func validateToken(token string) {
	return
}

func refreshToken(token string) {
	return
}
