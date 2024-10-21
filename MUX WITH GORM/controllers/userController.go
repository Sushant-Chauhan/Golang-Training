package controllers

import (
	"BankingApp/models"
	"BankingApp/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// CreateSuperAdminController handles the creation of the first super admin.
func CreateSuperAdminController(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		createdUser, err := services.CreateSuperAdmin(db, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(createdUser)
	}
}

// CreateUserController handles the creation of a new user.
func CreateUserController(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		createdUser, err := services.CreateUser(db, &user)
		if err != nil {
			http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(createdUser)
	}
}

// GetCustomerByIDController retrieves a customer by ID.
func GetCustomerByIDController(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		user, err := services.GetUserByID(db, uint(id))
		if err != nil {
			http.Error(w, "User not found: "+err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

// GetAllCustomersController retrieves all customers.
func GetAllCustomersController(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := services.GetAllUsers(db)
		if err != nil {
			http.Error(w, "Error retrieving users: "+err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(users)
	}
}

// UpdateCustomerController updates a customer by ID.
func UpdateCustomerController(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		user.ID = uint(id)
		updatedUser, err := services.UpdateUser(db, &user)
		if err != nil {
			http.Error(w, "Error updating user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(updatedUser)
	}
}

// DeleteCustomerController deletes a customer by ID.
func DeleteCustomerController(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		if err := services.DeleteUser(db, uint(id)); err != nil {
			http.Error(w, "Error deleting user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
