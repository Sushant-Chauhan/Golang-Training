package services

import (
	"BankingApp/models"
)

func CreateLedgerEntry(entry models.LedgerData) error {
	result := db.Create(&entry)
	return result.Error
}

func GetLedgerEntries(bankID int) ([]models.LedgerData, error) {
	var entries []models.LedgerData
	result := db.Where("bank_id = ?", bankID).Find(&entries)
	return entries, result.Error
}
