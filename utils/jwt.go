package utils

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

type Claims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(id uint, email, role string) (string, error) {
	claims := Claims{
		ID:    id,
		Email: email,
		Role:  role,
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

func ExtractToken(c *gin.Context) (string, error) {
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		return "", errors.New("token is required")
	}

	if strings.Split(bearerToken, " ")[0] != "Bearer" {
		return "", errors.New("invalid token type")
	}

	token := strings.Split(bearerToken, " ")[1]
	return token, nil
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	tokenString, err := ExtractToken(c)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok && !token.Valid {
		return nil, err
	}

	user := map[string]interface{}{
		"id":    claims.ID,
		"email": claims.Email,
		"role":  claims.Role,
	}

	return user, nil
}
