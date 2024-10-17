package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("my_secret_key") // Use a secret key for signing JWTs

// Claims defines the structure of the JWT payload
type Claims struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT token for the given user.
func GenerateToken(userID int, username string, isAdmin bool) (string, error) {
	claims := &Claims{
		UserID:   userID,
		Username: username,
		IsAdmin:  isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 1 day expiration
		},
	}

	// Create the JWT token with the claims and the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Middleware to check if a valid JWT token is present in the request header
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		claims := &Claims{}

		// Parse the JWT token and verify the claims
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Store user details in the request context for further use
		r = r.WithContext(SetUserContext(r.Context(), claims.UserID, claims.IsAdmin))
		next.ServeHTTP(w, r)
	})
}

// Ensure this is only declared once
func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Extract token from header
		tokenString := strings.Split(authHeader, " ")[1]
		claims := &jwt.MapClaims{}

		// Parse and validate the token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Provide the secret key used to sign the token
			return []byte("your-secret-key"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Check if the user is an admin
		if isAdmin, ok := (*claims)["isAdmin"].(bool); !ok || !isAdmin {
			http.Error(w, "Forbidden: Admins only", http.StatusForbidden)
			return
		}

		// Proceed to the next handler if admin
		next.ServeHTTP(w, r)
	})
}

// Middleware to check if the user is active
func ActiveUserOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, isAdmin := GetUserFromContext(r.Context())

		if !isAdmin && !CheckUserActive(userID) {
			http.Error(w, "Inactive users cannot perform this action", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
