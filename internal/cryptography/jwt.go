package cryptography

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GetJWTSecretKey() (string, error) {
	if os.Getenv("CI") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
		}
	}
	password := os.Getenv("JWT_SECRET_KEY")
	if password == "" {
		return "", errors.New("no JWT_SECRET_KEY found")
	}
	return password, nil
}

func ExtractToken(authorizationHeader string) (string, error) {
	const prefix = "Bearer "
	if authorizationHeader == "" {
		return "", errors.New("header Authorization is not available")
	}

	if !strings.HasPrefix(authorizationHeader, prefix) {
		return "", errors.New("not correct form of header Authorization")
	}

	return strings.TrimPrefix(authorizationHeader, prefix), nil
}

func GenerateAccessToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	password, err := GetJWTSecretKey()
	if err != nil {
		return "", err
	}

	return token.SignedString([]byte(password))
}

func GenerateRefreshToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	password, err := GetJWTSecretKey()
	if err != nil {
		return "", err
	}

	return token.SignedString([]byte(password))
}
