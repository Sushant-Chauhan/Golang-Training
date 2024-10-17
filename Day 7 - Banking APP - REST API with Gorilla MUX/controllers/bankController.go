package controllers

import (
	"bankapp/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Secure the CreateBank endpoint for Admin Only
func CreateBank(w http.ResponseWriter, r *http.Request) {
	var bank services.Bank
	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newBank, err := services.CreateBank(bank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newBank)
}

func GetAllBanks(w http.ResponseWriter, r *http.Request) {
	banks := services.GetAllBanks()
	json.NewEncoder(w).Encode(banks)
}

func GetBankByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	bank, err := services.GetBankByID(id)
	if err != nil {
		http.Error(w, "Bank not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(bank)
}

func UpdateBankByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var bank services.Bank
	err := json.NewDecoder(r.Body).Decode(&bank)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err = services.UpdateBankByID(id, bank)
	if err != nil {
		http.Error(w, "Bank not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(bank)
}

func DeleteBankByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	err := services.DeleteBankByID(id)
	if err != nil {
		http.Error(w, "Bank not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Bank deleted")
}
