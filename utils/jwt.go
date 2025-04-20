package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   int    `json:"userid"`
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userID int, username, role string) (string, error) {

	if len(jwtSecret) == 0 {
		panic("JWT_SECRET not set in environment")
	}

	expStr := os.Getenv("JWT_EXPIRATION")

	if expStr == "" {
		expStr = "24h"

	}

	duration, err := time.ParseDuration(expStr)
	if err != nil {
		duration = 24 * time.Hour

	}

	claims := Claims{
		UserID:   userID,
		UserName: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)

}

func ParseJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		return jwtSecret, nil

	})

	if err != nil {
		return nil, errors.New("Invalid or expired token")

	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("Could not parse claims")
	}

	return claims, nil

}
