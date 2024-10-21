// folder :
// components > user> service > userService.go

package service

import (
	"errors"
	"fmt"
)

var allUsers []*User

type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	Age      float32 `json:"age"`
	IsAdmin  bool    `json:"isAdmin"`
}

func GetUserByID(userID int) (*User, error) {
	fmt.Println("GetUserByID service called for user ID:", userID)
	for _, user := range allUsers {
		if user.ID == userID {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func NewUser(username, password, name string, age float32, isAdmin bool) (*User, error) {
	// Add validation logic here if needed
	user := &User{
		ID:       len(allUsers) + 1, // simple ID generation logic
		Username: username,
		Password: password,
		Name:     name,
		Age:      age,
		IsAdmin:  isAdmin,
	}
	allUsers = append(allUsers, user)
	return user, nil
}
