package jwt

import (
	"errors"
	"log"

	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/nibir30/go-microservices/auth/internal/model/data"
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

func VerifyToken(tokenString string) (*data.JwtUser, error) {
	errorENV := godotenv.Load(".env")
	if errorENV != nil {
		panic("Failed to load env file")

	}
	secretKey := os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		log.Printf("invalid token")
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return &data.JwtUser{
		Username: claims["username"].(string),
		Expires:  time.Unix(int64(claims["exp"].(float64)), 0),
	}, nil
}
