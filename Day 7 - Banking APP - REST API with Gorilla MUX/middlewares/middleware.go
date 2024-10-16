package middleware

import (
	"net/http"
	// "strings"
	// "fmt"
)

// AuthMiddleware checks if the request contains a valid JWT token and whether the user is an admin.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		// // Bearer token should start with "Bearer "
		// if !strings.HasPrefix(authHeader, "Bearer ") {
		// 	http.Error(w, "Invalid authorization token format", http.StatusUnauthorized)
		// 	return
		// }

		// Trim the "Bearer " prefix and extract the token
		token := authHeader

		// Verify the JWT token
		claims, err := VerifyJWT(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Ensure the user is an admin
		if !claims.IsAdmin {
			http.Error(w, "Unauthorized: Admin access required", http.StatusForbidden)
			return
		}

		// User is authenticated and authorized as admin, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
