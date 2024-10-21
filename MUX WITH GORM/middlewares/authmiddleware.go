package middlewares

var secretKey = []byte("secret_key")

/*
// ///////////// TokenAuthMiddleware checks the JWT token for validity and extracts claims.
func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := VerifyJWT(tokenStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Insert the claims into the context for use in the handler
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ///////////// VerifyJWT parses and validates a JWT token and extracts the claims.
func VerifyJWT(tokenStr string) (*models.Claims, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func VerifyAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value("claims").(*models.Claims)

		if claims == nil || !claims.IsAdmin {
			http.Error(w, "Unauthorized: admin access required", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func VerifyCustomer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value("claims").(*models.Claims)

		if claims == nil || claims.IsAdmin {
			http.Error(w, "Unauthorized: customer access required", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
*/
