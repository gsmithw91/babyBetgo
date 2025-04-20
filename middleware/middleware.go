package middleware

import (
	"babybetgo/utils"
	"context"
	"log"
	"net/http"
	"strings"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		end := time.Since(start)
		log.Printf("%s %s %s", r.Method, r.URL.Path, end)

	})

}

func RequireAuth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Authorization header from request
		authHeader := r.Header.Get("Authorization")

		// Check if authroization header is empty or missing 'Bearer '
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid or missing authorization header", http.StatusUnauthorized)
			return
		}

		// Remove the Bearer Prefix
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse the JWT for the claims
		claims, err := utils.ParseJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)

		}
		// Create a context with claims
		ctx := context.WithValue(r.Context(), "user", claims)

		// the next Servers the http with the updated contexxt containing the claims
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
