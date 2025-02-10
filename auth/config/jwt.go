package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func CreateToken(username string) (string, error) {
	errorENV := godotenv.Load(".env")
	if errorENV != nil {
		panic("Failed to load env file")
	}
	secretKey := os.Getenv("SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	// tokenString, err := token.SignedString(secretKey)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	errorENV := godotenv.Load("../.env")
	if errorENV != nil {
		panic("Failed to load env file")
	}
	secretKey := os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		log.Printf("invalid token")
		return errors.New("invalid token")
	}

	return nil
}
