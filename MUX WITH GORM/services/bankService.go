package services

import (
	"BankingApp/models"
)

func CreateBank(bank models.Bank) (models.Bank, error) {
	result := models.DB.Create(&bank)
	if result.Error != nil {
		return models.Bank{}, result.Error
	}
	return bank, nil
}

func GetAllBanks() ([]models.Bank, error) {
	var banks []models.Bank
	result := models.DB.Find(&banks)
	if result.Error != nil {
		return nil, result.Error
	}
	return banks, nil
}

func GetBankByID(id int) (models.Bank, error) {
	var bank models.Bank
	result := models.DB.First(&bank, id)
	if result.Error != nil {
		return models.Bank{}, result.Error
	}
	return bank, nil
}

func UpdateBank(bank models.Bank) (models.Bank, error) {
	result := models.DB.Save(&bank)
	if result.Error != nil {
		return models.Bank{}, result.Error
	}
	return bank, nil
}

func DeleteBank(id int) error {
	result := models.DB.Delete(&models.Bank{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
