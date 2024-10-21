package services

import (
	"BankingApp/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitService(database *gorm.DB) {
	db = database
}

func CreateAccount(account models.Account) (models.Account, error) {
	if account.Balance < 1000 {
		return models.Account{}, errors.New("minimum balance should be 1000")
	}
	result := db.Create(&account)
	if result.Error != nil {
		return models.Account{}, result.Error
	}
	return account, nil
}

func GetAllAccounts() ([]models.Account, error) {
	var accounts []models.Account
	result := db.Find(&accounts)
	if result.Error != nil {
		return nil, result.Error
	}
	return accounts, nil
}

func GetAccountByID(id int) (models.Account, error) {
	var account models.Account
	result := db.First(&account, id)
	if result.Error != nil {
		return models.Account{}, result.Error
	}
	return account, nil
}

func UpdateAccount(account models.Account) (models.Account, error) {
	if account.Balance < 0 {
		return models.Account{}, errors.New("account balance cannot be negative")
	}
	result := db.Save(&account)
	if result.Error != nil {
		return models.Account{}, result.Error
	}
	return account, nil
}

func DeleteAccount(id int) error {
	result := db.Delete(&models.Account{}, id)
	return result.Error
}

func Withdraw(accountID int, amount float64) (models.Account, error) {
	var account models.Account
	result := db.First(&account, accountID)
	if result.Error != nil {
		return account, result.Error
	}

	if account.Balance < amount {
		return account, errors.New("insufficient funds")
	}

	account.Balance -= amount
	result = db.Save(&account)
	return account, result.Error
}

func Deposit(accountID int, amount float64) (models.Account, error) {
	var account models.Account
	result := db.First(&account, accountID)
	if result.Error != nil {
		return account, result.Error
	}

	account.Balance += amount
	result = db.Save(&account)
	return account, result.Error
}

func Transfer(fromAccountID, toAccountID int, amount float64) error {
	var fromAccount, toAccount models.Account
	if err := db.First(&fromAccount, fromAccountID).Error; err != nil {
		return err
	}
	if err := db.First(&toAccount, toAccountID).Error; err != nil {
		return err
	}

	if fromAccount.Balance < amount {
		return errors.New("insufficient funds")
	}

	fromAccount.Balance -= amount
	toAccount.Balance += amount

	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&fromAccount).Error; err != nil {
			return err
		}
		if err := tx.Save(&toAccount).Error; err != nil {
			return err
		}
		return nil
	})

	return nil
}

func AddTransaction(accountID, correspondingAccountID, correspondingBankID int, transactionType string, amount float64, db *gorm.DB) error {
	var account models.Account
	if err := db.First(&account, accountID).Error; err != nil {
		return err
	}

	// Calculate new balance based on transaction type
	newBalance := account.Balance
	if transactionType == "credit" {
		newBalance += amount
	} else if transactionType == "debit" {
		newBalance -= amount
		if newBalance < 0 {
			return errors.New("insufficient funds")
		}
	}

	// Create transaction entry
	transaction := models.TransactionEntry{
		AccountID:              accountID,
		CorrespondingAccountID: correspondingAccountID,
		CorrespondingBankID:    correspondingBankID,
		Type:                   transactionType,
		Amount:                 amount,
		BalanceAfter:           newBalance,
		Timestamp:              time.Now(),
	}

	// Save transaction
	if err := db.Create(&transaction).Error; err != nil {
		return err
	}

	// Update account balance
	account.Balance = newBalance
	return db.Save(&account).Error
}

func GetPassbook(accountID int, db *gorm.DB) ([]models.TransactionEntry, error) {
	var transactions []models.TransactionEntry
	if err := db.Where("account_id = ?", accountID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
