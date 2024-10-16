
package services

import (
	"bankingApp/models"
	"errors"
)

var allCustomers []models.User //non-admin customers only
var customerIDCounter = 1

// CreateCustomer creates a new customer
func CreateCustomer(username, password, firstName, lastName string) (*models.User, error) {
	// Check if username already exists
	for _, customer := range models.Users {
		if customer.Username == username {
			return nil, errors.New("username already exists")
		}
	}

	customer := &models.User{
		UserID:    customerIDCounter,
		Username:  username,
		Password:  password, // Remember to hash passwords in real apps
		FirstName: firstName,
		LastName:  lastName,
	}

	customerIDCounter++
	allCustomers = append(allCustomers, *customer)
	models.Users=append(models.Users, *customer)
	return customer, nil
}

// GetCustomerByID fetches a customer by their user ID
func GetCustomerByID(userID int) (*models.User, error) {
	for _, customer := range allCustomers {
		if customer.UserID == userID {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}

// GetAllCustomers returns all customers
func GetAllCustomers() []models.User {
	return allCustomers
}

// UpdateCustomer updates the information of an existing customer
func UpdateCustomer(userID int, newUsername, newPassword, firstName, lastName string) (*models.User, error) {
	//validations???
	for i, customer := range allCustomers {
		if customer.UserID == userID {
			allCustomers[i].Username = newUsername
			allCustomers[i].Password = newPassword
			allCustomers[i].FirstName = firstName
			allCustomers[i].LastName = lastName
			return &allCustomers[i], nil
		}
	}
	return nil, errors.New("customer not found")
}

// DeleteCustomer removes a customer by their ID
func DeleteCustomer(userID int) error {
	for i, customer := range allCustomers {
		if customer.UserID == userID {
			allCustomers = append(allCustomers[:i], allCustomers[i+1:]...)
			models.Users = append(models.Users[:i], models.Users[i+1:]...)

			return nil
		}
	}
	return errors.New("customer not found")
}
