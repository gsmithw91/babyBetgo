package middleware

import (
	"babybetgo/utils"
	"net/http"
	"strings"
)

func RequireAuth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Authorization header from request
		authHeader := r.Header.Get("Authorization")

		// Check if authorization header is empty or missing 'Bearer '
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
			return
		}
		// Create a context with claims
		ctx := utils.SetClaimsInContext(r.Context(), claims)

		// the next serves the http with the updated context containing the claims
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
