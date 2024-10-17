package controllers

import (
	"bankapp/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// the first admin if no admins exist
func CreateSuperAdmin(w http.ResponseWriter, r *http.Request) {
	var user services.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if an admin already exists
	if services.HasAdmins() {
		http.Error(w, "Admin already exists", http.StatusConflict)
		return
	}

	user.IsAdmin = true
	user.IsActive = true

	// Create the user and capture both return values
	newUser, err := services.CreateUser(user)
	if err != nil {
		http.Error(w, "Could not create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created user as JSON
	json.NewEncoder(w).Encode(newUser)
}

// CreateUser handles the user creation request
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user services.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call CreateUser service function
	newUser, err := services.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(newUser)
}

// GetUser retrieves a user by ID.
func GetUserByID(w http.ResponseWriter, r *http.Request) { // CORE LOGIC IN SERVICES - here just error checking
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	user, err := services.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var user services.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err = services.UpdateUserByID(id, user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	err := services.DeleteUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User deleted")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := services.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}
