package utils

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/hacktiv8-ks07-g04/final-project-4/pkg/errs"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

type Claims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(id uint, email string) (string, error) {
	claims := Claims{
		ID:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return ss, nil
}

func ExtractToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errs.Unauthorized("Token is required"))
	}

	// Check token type
	if strings.Split(bearerToken, " ")[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errs.Unauthorized("Invalid token type"))
	}

	// Get bearer token
	token := strings.Split(bearerToken, " ")[1]
	return token
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	tokenString := ExtractToken(c)

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	claims, ok := token.Claims.(*Claims)
	if !ok && !token.Valid {
		return nil, err
	}

	user := map[string]interface{}{
		"id":    claims.ID,
		"email": claims.Email,
	}

	return user, nil
}
