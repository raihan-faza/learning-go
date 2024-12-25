package api

import (
	"fmt"
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

func validateToken(user_token string) error {
	token, err := jwt.Parse(user_token, func(t *jwt.Token) (interface{}, error) {
		return os.Getenv("secret"), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
