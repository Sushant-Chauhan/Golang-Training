package controllers

import (
	"BankingApp/models"
	"BankingApp/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateLedgerEntry(w http.ResponseWriter, r *http.Request) {
	var entry models.LedgerData
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.CreateLedgerEntry(entry); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetLedgerEntries(w http.ResponseWriter, r *http.Request) {
	bankIDStr := mux.Vars(r)["bankID"]
	bankID, err := strconv.Atoi(bankIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	entries, err := services.GetLedgerEntries(bankID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(entries)
}
