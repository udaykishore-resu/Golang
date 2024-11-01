package middleware

import (
	"net/http"
	"strings"
	"timesheet-app/utils"
)

// AuthMiddleware checks if the JWT token is valid before proceeding
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the request header
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Authorization token is required", http.StatusUnauthorized)
			return
		}

		// Check if the token starts with "Bearer "
		if !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Authorization token must be Bearer token", http.StatusUnauthorized)
			return
		}

		// Extract the actual token from the Bearer token
		token = strings.TrimPrefix(token, "Bearer ")

		// Validate the token using utils
		valid, _ := utils.ValidateJWT(token)
		if !valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
