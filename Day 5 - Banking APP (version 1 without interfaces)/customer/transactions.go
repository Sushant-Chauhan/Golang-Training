package customer

import (
	"fmt"
	"time"
)

// Transaction struct to represent a financial transaction
type Transaction struct {
	Type            string    // Type of transaction: debit or credit
	Amount          float64   // Amount of the transaction
	BalanceAfter    float64   // Balance after the transaction
	AccountID       int       // Account ID of the corresponding account
	BankID          int       // Bank ID of the corresponding account
	Timestamp       time.Time // Timestamp of the transaction
}

// NewTransaction factory function to create a new Transaction
func NewTransaction(transactionType string, amount float64, balanceAfter float64, accountID int, bankID int, timestamp time.Time) (*Transaction, error) {
	if transactionType != "debit" && transactionType != "credit" {
		return nil, fmt.Errorf("invalid transaction type: %s", transactionType)
	}
	return &Transaction{
		Type:         transactionType,
		Amount:       amount,
		BalanceAfter: balanceAfter,
		AccountID:    accountID,
		BankID:       bankID,
		Timestamp:    timestamp,
	}, nil
}

// Getters
func (t *Transaction) GetType() string {
	return t.Type
}

func (t *Transaction) GetAmount() float64 {
	return t.Amount
}

func (t *Transaction) GetBalanceAfter() float64 {
	return t.BalanceAfter
}

func (t *Transaction) GetAccountID() int {
	return t.AccountID
}

func (t *Transaction) GetBankID() int {
	return t.BankID
}

func (t *Transaction) GetTimestamp() time.Time {
	return t.Timestamp
}

// Setters
func (t *Transaction) SetType(transactionType string) error {
	if transactionType != "debit" && transactionType != "credit" {
		return fmt.Errorf("invalid transaction type: %s", transactionType)
	}
	t.Type = transactionType
	return nil
}

func (t *Transaction) SetAmount(amount float64) {
	t.Amount = amount
}

func (t *Transaction) SetBalanceAfter(balance float64) {
	t.BalanceAfter = balance
}

func (t *Transaction) SetAccountID(accountID int) {
	t.AccountID = accountID
}

func (t *Transaction) SetBankID(bankID int) {
	t.BankID = bankID
}

func (t *Transaction) SetTimestamp(timestamp time.Time) {
	t.Timestamp = timestamp
}

// PrintTransaction prints the details of the transaction
func (t *Transaction) PrintTransaction() {
	fmt.Printf("Transaction Type: %s\nAmount: %.2f\nBalance After: %.2f\nAccount ID: %d\nBank ID: %d\nTimestamp: %s\n",
		t.Type, t.Amount, t.BalanceAfter, t.AccountID, t.BankID, t.Timestamp.Format(time.RFC1123))
}
