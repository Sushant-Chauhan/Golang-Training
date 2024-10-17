package services

import (
	"errors"
)

type Account struct {
	AccountID int
	Balance   float64
	IsActive  bool
	BankID    int
}

var allAccounts []*Account
var accountID int = 1

// CreateAccount creates a new account for an active user with an initial balance of Rs. 1000.
func CreateAccount(user *User, bankID int) (*Account, error) {
	if !user.IsActive {
		return nil, errors.New("inactive users cannot create accounts")
	}

	if user.IsAdmin {
		return nil, errors.New("admins cannot create accounts")
	}

	initialBalance := 1000.0

	newAccount := &Account{
		AccountID: accountID,
		Balance:   initialBalance,
		IsActive:  true,
		BankID:    bankID,
	}

	allAccounts = append(allAccounts, newAccount)
	accountID++

	return newAccount, nil
}

// GetAccountByID retrieves an account by its ID.
func GetAccountByID(accountID int) (*Account, error) {
	for _, account := range allAccounts {
		if account.AccountID == accountID {
			return account, nil
		}
	}
	return nil, errors.New("account not found")
}

// UpdateAccount updates account details.
func UpdateAccount(accountID int, balance float64, isActive bool) (*Account, error) {
	account, err := GetAccountByID(accountID)
	if err != nil {
		return nil, err
	}

	account.Balance = balance
	account.IsActive = isActive
	return account, nil
}

// DeleteAccount deletes an account.
func DeleteAccount(accountID int) error {
	for i, account := range allAccounts {
		if account.AccountID == accountID {
			allAccounts = append(allAccounts[:i], allAccounts[i+1:]...) // Remove the account
			return nil
		}
	}
	return errors.New("account not found")
}

// Deposit adds money to the account balance.
func Deposit(accountID int, amount float64) (*Account, error) {
	account, err := GetAccountByID(accountID)
	if err != nil {
		return nil, err
	}

	if amount <= 0 {
		return nil, errors.New("deposit amount must be greater than zero")
	}

	account.Balance += amount
	return account, nil
}

// Withdraw subtracts an amount from the account balance.
func Withdraw(accountID int, amount float64) (*Account, error) {
	account, err := GetAccountByID(accountID)
	if err != nil {
		return nil, err
	}

	if amount <= 0 {
		return nil, errors.New("withdrawal amount must be greater than zero")
	}

	if account.Balance < amount {
		return nil, errors.New("insufficient funds")
	}

	account.Balance -= amount
	return account, nil
}

// Transfer moves money from one account to another.
func Transfer(fromID, toID int, amount float64) error {
	fromAccount, err := GetAccountByID(fromID)
	if err != nil {
		return err
	}

	toAccount, err := GetAccountByID(toID)
	if err != nil {
		return err
	}

	if amount <= 0 {
		return errors.New("transfer amount must be greater than zero")
	}

	if fromAccount.Balance < amount {
		return errors.New("insufficient funds for transfer")
	}

	fromAccount.Balance -= amount
	toAccount.Balance += amount
	return nil
}
