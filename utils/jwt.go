// utils.go
package utils

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   int    `json:"userid"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var jwtSecret []byte
var once sync.Once

func loadSecret() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET not set in environment")
	}
	jwtSecret = []byte(secret)
}

func GenerateJWT(userID int, username, role string) (string, error) {
	once.Do(loadSecret)

	duration, err := time.ParseDuration(os.Getenv("JWT_EXPIRATION"))
	if err != nil {
		duration = 24 * time.Hour
	}

	claims := Claims{
		UserID:   userID,
		Username: username,
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
	once.Do(loadSecret)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		fmt.Println("JWT parse error:", err)
		return nil, errors.New("Invalid or expired token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		fmt.Println("Token claims invalid or not valid")
		return nil, errors.New("Could not parse claims or token is invalid")
	}

	return claims, nil
}

type contextKey string

const userContextKey = contextKey("user")

func SetClaimsInContext(ctx context.Context, claims *Claims) context.Context {
	return context.WithValue(ctx, userContextKey, claims)

}
func GetClaimsFromContext(ctx context.Context) (*Claims, error) {
	claims, ok := ctx.Value(userContextKey).(*Claims)
	if !ok || claims == nil {
		return nil, errors.New("no claims found in context")
	}
	return claims, nil

}
