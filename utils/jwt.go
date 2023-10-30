package utils

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parsedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parsedToken.SignedString([]byte(SECRET_KEY))

	return signedToken
}
