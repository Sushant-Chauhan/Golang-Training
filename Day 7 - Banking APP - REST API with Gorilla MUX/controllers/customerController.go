

package controllers

import (
	middleware "bankingApp/middlewares"
	// "bankingApp/models"
	"bankingApp/services"
	"encoding/json"
	// "fmt"
	"net/http"
	"strconv"
	"strings"
)


// CreateCustomerController handles creating a new customer
func CreateCustomerController(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	claims, err := middleware.VerifyJWT(strings.TrimPrefix(token, "Bearer "))
	if err != nil || !claims.IsAdmin {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var requestData struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		FirstName string `json:"firstName"` 
		LastName  string `json:"lastName"`  
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	customer, err := services.CreateCustomer(requestData.Username, requestData.Password, requestData.FirstName, requestData.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(customer)
}

// GetCustomerByIDController retrieves a customer by ID
func GetCustomerByIDController(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("UserId")

	if idStr == "" {
		http.Error(w, "UserId parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid UserId parameter", http.StatusBadRequest)
		return
	}

	// Fetch customer by userID from service
	customer, err := services.GetCustomerByID(userID)
	if err != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(customer)
}

// GetAllCustomersController retrieves all customers
func GetAllCustomersController(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	claims, err := middleware.VerifyJWT(strings.TrimPrefix(token, "Bearer "))
	if err != nil || !claims.IsAdmin {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	customers := services.GetAllCustomers()
	json.NewEncoder(w).Encode(customers)
}

// UpdateCustomerController updates a customer by ID
func UpdateCustomerController(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	claims, err := middleware.VerifyJWT(strings.TrimPrefix(token, "Bearer "))
	if err != nil || !claims.IsAdmin {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := r.URL.Query().Get("UserId")
	if idStr == "" {
		http.Error(w, "UserId parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	var requestData struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		FirstName string `json:"firstName"` // Updated field
		LastName  string `json:"lastName"`  // Updated field
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	customer, err := services.UpdateCustomer(userID, requestData.Username, requestData.Password, requestData.FirstName, requestData.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(customer)
}

// DeleteCustomerController deletes a customer by ID
func DeleteCustomerController(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	claims, err := middleware.VerifyJWT(strings.TrimPrefix(token, "Bearer "))
	if err != nil || !claims.IsAdmin {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := r.URL.Query().Get("UserId")
	if idStr == "" {
		http.Error(w, "UserId parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	if err := services.DeleteCustomer(userID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

