package services

import (
	"errors"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Username  string // Add this field for user validation
	Password  string // Password field to check against
	IsAdmin   bool
	IsActive  bool
}

var users []User

func CreateUser(user User) (User, error) {

	if err := validateNewUser(user); err != nil {
		return User{}, err
	}

	// unique ID to the new user
	user.ID = findUserID()
	users = append(users, user)

	return user, nil
}

// checks if the user fields are valid. /// PRIVATE
func validateNewUser(user User) error {
	if user.FirstName == "" {
		return errors.New("first name cannot be empty")
	}
	if user.LastName == "" {
		return errors.New("last name cannot be empty")
	}
	if user.Username == "" {
		return errors.New("username cannot be empty")
	}
	if user.Password == "" {
		return errors.New("password cannot be empty")
	}
	return nil
}

// next unique user ID.
func findUserID() int {
	if len(users) == 0 {
		return 1
	}
	return users[len(users)-1].ID + 1
}

func GetUserByID(userID int) (*User, error) {
	for _, user := range users {
		if user.ID == userID {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

// checks username and password match an existing user. // PUBLIC   --- CONTROLLER VALIDATION
func ValidateUser(username, password string) (User, error) {
	for _, user := range users {
		if user.Username == username && user.Password == password && user.IsActive {
			return user, nil
		}
	}
	return User{}, errors.New("invalid username or password")
}

func HasAdmins() bool {
	for _, user := range users {
		if user.IsAdmin {
			return true
		}
	}
	return false
}

func UpdateUserByID(id int, updatedUser User) error {
	for i, user := range users {
		if user.ID == id {
			users[i] = updatedUser
			return nil
		}
	}
	return errors.New("User not found")
}

func DeleteUserByID(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("User not found")
}

func GetAllUsers() []User {
	return users
}
