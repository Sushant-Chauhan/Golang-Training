package services

import (
	"BankingApp/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateSuperAdmin handles the creation of the first super admin.
func CreateSuperAdmin(db *gorm.DB, user *models.User) (*models.User, error) {
	// Check if any admin exists in the system
	var adminCount int64
	db.Model(&models.User{}).Where("is_admin = ?", true).Count(&adminCount)
	if adminCount > 0 {
		return nil, errors.New("an admin already exists")
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Set user as Super Admin
	user.Password = string(hashedPassword)
	user.IsAdmin = true
	user.IsActive = true

	// Create the Super Admin in the database
	if result := db.Create(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// CreateUser handles the creation of a new user.
func CreateUser(db *gorm.DB, user *models.User) (*models.User, error) {
	if user.Username == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" {
		return nil, errors.New("required fields are missing")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	// Create the user in the database
	if result := db.Create(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// GetUserByID retrieves a user by ID.
func GetUserByID(db *gorm.DB, id uint) (*models.User, error) {
	var user models.User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetAllUsers retrieves all users from the database.
func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// UpdateUser updates a user's details.
func UpdateUser(db *gorm.DB, user *models.User) (*models.User, error) {
	result := db.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// DeleteUser deletes a user by ID.
func DeleteUser(db *gorm.DB, id uint) error {
	result := db.Delete(&models.User{}, id)
	return result.Error
}

//////////// AuthenticateUser checks the user's credentials and returns the user if valid.
func AuthenticateUser(db *gorm.DB, username, password string) (*models.User, error) {
	var user models.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
