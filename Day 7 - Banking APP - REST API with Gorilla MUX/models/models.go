
package models

import "errors"

// User represents a user in the system
type User struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	FirstName string `json:"firstName"` 
	LastName  string `json:"lastName"`  
	IsAdmin   bool   `json:"isAdmin"`
}

// Initialize some predefined users
var Users = []User{
	{UserID: 1, Username: "admin", Password: "admin123", FirstName: "Admin", LastName: "User", IsAdmin: true},
	{UserID: 2, Username: "user", Password: "user123", FirstName: "Regular", LastName: "User", IsAdmin: false},
}

// AuthenticateUser validates the username and password
func AuthenticateUser(username, password string) (*User, error) {
	for _, user := range Users {
		if user.Username == username && user.Password == password {
			return &user, nil
		}
	}
	return nil, errors.New("invalid credentials")
}

