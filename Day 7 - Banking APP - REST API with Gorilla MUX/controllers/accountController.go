package controllers

import (
	"bankapp/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// creates a new account for a user.
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, _ := strconv.Atoi(params["userID"])
	bankID, _ := strconv.Atoi(params["bankID"])

	// Fetch the user (customer)
	user, err := services.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Pass user as a pointer
	account, err := services.CreateAccount(&user, bankID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(account)
}

// GetAccount retrieves account details.
func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID, _ := strconv.Atoi(params["id"])

	account, err := services.GetAccountByID(accountID)
	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(account)
}

// // UpdateAccount updates account details.
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID, _ := strconv.Atoi(params["id"])

	var updatedAccountInfo struct {
		Balance  float64 `json:"balance"`
		IsActive bool    `json:"isActive"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updatedAccountInfo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedAccount, err := services.UpdateAccount(accountID, updatedAccountInfo.Balance, updatedAccountInfo.IsActive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(updatedAccount)
}

// DeleteAccount deletes an account.
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID, _ := strconv.Atoi(params["id"])

	err := services.DeleteAccount(accountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent) // No content response for successful deletion
}

// Deposit adds money to an account.
func Deposit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID, _ := strconv.Atoi(params["id"])

	var depositInfo struct {
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&depositInfo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	account, err := services.Deposit(accountID, depositInfo.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(account)
}

// Withdraw removes money from an account.
func Withdraw(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID, _ := strconv.Atoi(params["id"])

	var withdrawInfo struct {
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&withdrawInfo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	account, err := services.Withdraw(accountID, withdrawInfo.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(account)
}

// Transfer moves money from one account to another.
func Transfer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fromID, _ := strconv.Atoi(params["fromID"])
	toID, _ := strconv.Atoi(params["toID"])

	var transferInfo struct {
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&transferInfo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := services.Transfer(fromID, toID, transferInfo.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
