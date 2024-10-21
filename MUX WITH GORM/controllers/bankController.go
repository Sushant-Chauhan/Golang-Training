package controllers

import (
	"BankingApp/models"
	"BankingApp/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateBank(w http.ResponseWriter, r *http.Request) {
	var bank models.Bank
	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bank, err := services.CreateBank(bank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bank)
}

func GetAllBanks(w http.ResponseWriter, r *http.Request) {
	banks, err := services.GetAllBanks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(banks)
}

func GetBankByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid bank ID", http.StatusBadRequest)
		return
	}

	bank, err := services.GetBankByID(id)
	if err != nil {
		http.Error(w, "Bank not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(bank)
}

func UpdateBankByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid bank ID", http.StatusBadRequest)
		return
	}

	var bank models.Bank
	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	bank.ID = id
	bank, err = services.UpdateBank(bank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bank)
}

func DeleteBankByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid bank ID", http.StatusBadRequest)
		return
	}

	if err := services.DeleteBank(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
