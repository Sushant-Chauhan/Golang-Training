package account

import (
    "bankingApp/customer" //  import for transactions.go - transactions
	"errors"
	"time"
)

//represent customer account
type Account struct {
	AccountNo int
	BankID    int
	Balance   float64
	IsActive  bool
	Passbook  *Passbook  // Use the Passbook struct here
}

var allAccounts []*Account
var accountID int

//////////  Passbook code ///////////////////
// Passbook struct that stores all transactions related to an account
type Passbook struct {
    transactions []*customer.Transaction


// NewPassbook factory function
func NewPassbook(initialBalance float64, accountID int, bankID int) (*Passbook, error) {
    transaction, err := customer.NewTransaction("credit", initialBalance, initialBalance, accountID, bankID, time.Now())
    if err != nil {
        return nil, err
    }
    passbook := &Passbook{
        transactions: []*customer.Transaction{transaction},
    }
    return passbook, nil
}

// PrintPassbook prints all transactions in the passbook
func (p *Passbook) PrintPassbook() {
    if len(p.transactions) == 0 {
        println("No transactions available in the passbook.")
        return
    }

    println("Passbook Transactions:")
    for _, transaction := range p.transactions {
        transaction.PrintTransaction()         //check correcty implemented
    }
}


/////////////  Account code ////////////////
// NewAccount creates a new account with validation for the initial balance
func NewAccount(initialBalance float64, bankID int) (*Account, error) {
	if initialBalance < 1000 {
		return nil, errors.New("initial balance must be at least Rs. 1000")
	}

	accountID++
	passbook, err := NewPassbook(initialBalance, accountID, bankID)
	if err != nil {
		return nil, err
	}
	account := &Account{
		AccountNo: accountID,
		BankID:    bankID,
		Balance:   initialBalance,
		IsActive:  true,
		Passbook:  passbook, // Initialize with a new Passbook from customer package
	}

	allAccounts = append(allAccounts, account)
	return account, nil
}

// Getter for transactions
func (p *Passbook) GetTransactions() []*customer.Transaction {
	return p.transactions
}

// Setter for transactions
func (p *Passbook) SetTransactions(transactions []*customer.Transaction) {
	p.transactions = transactions
}

// AddTransaction method to add a new transaction to the passbook
func (p *Passbook) AddTransaction(t *Transaction) {
	p.transactions = append(p.transactions, t)
}

// GetAccountByID retrieves an account by its ID
func GetAccountByID(accountID int) (*Account, error) {
	for _, account := range allAccounts {
		if account.AccountNo == accountID {
			return account, nil
		}
	}
	return nil, errors.New("account not found")
}

// GetAllAccounts returns a slice of all accounts
func GetAllAccounts() []*Account {
	return allAccounts
}

// UpdateAccount updates an account's attribute
func (a *Account) UpdateAccount(attribute string, newValue interface{}) error {
	switch attribute {
	case "Balance":
		if val, ok := newValue.(float64); ok {
			a.Balance = val
		} else {
			return errors.New("invalid value type for Balance")
		}
	case "IsActive":
		if val, ok := newValue.(bool); ok {
			a.IsActive = val
		} else {
			return errors.New("invalid value type for IsActive")
		}
	default:
		return errors.New("attribute not recognized")
	}
	return nil
}

// DeactivateAccount deactivates the account
func (a *Account) DeactivateAccount() {
	a.SetIsActive(false)
}

// Deposit adds money to the account
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	a.Balance += amount

	// Create and add transaction to passbook
	transaction, err := transactions.NewTransaction("credit", amount, a.Balance, a.AccountNo, a.BankID, time.Now())
	if err != nil {
		return err
	}
	a.Passbook.AddTransaction(transaction)

	return nil
}

// Withdraw subtracts money from the account
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	if amount > a.Balance {
		return errors.New("insufficient balance")
	}

	a.Balance -= amount

	// Create and add transaction to passbook
	transaction, err := transactions.NewTransaction("debit", amount, a.Balance, a.AccountNo, a.BankID, time.Now())
	if err != nil {
		return err
	}
	a.Passbook.AddTransaction(transaction)

	return nil
}

// Getter and Setter functions
func (a *Account) GetAccountNumber() int {
	return a.AccountNo
}

func (a *Account) GetBankID() int {
	return a.BankID
}

func (a *Account) GetIsActive() bool {
	return a.IsActive
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func (a *Account) SetIsActive(active bool) {
	a.IsActive = active
}

func (a *Account) GetPassbook() *Passbook {
	return a.Passbook
}

func (a *Account) SetPassbook(passbook *Passbook) {
	a.Passbook = passbook
}
