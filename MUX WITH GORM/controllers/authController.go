package controllers

///////////////////////////////////////

var secretKey = []byte("your_secret_key")

/*
// LoginController authenticates the user and generates a JWT token.
func LoginController(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginDetails struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&loginDetails); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Authenticate user
		user, err := services.AuthenticateUser(db, loginDetails.Username, loginDetails.Password)
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		// Create JWT claims
		expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours
		claims := &models.Claims{
			UserID:   user.ID,
			Username: user.Username,
			IsAdmin:  user.IsAdmin,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		// Generate JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		// Send the token in the response
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}
*/
