package bank

import (
	"errors"
	"fmt"
	"strings"
	// "bankingApp/customer"
)

type LedgerEntry struct {
    BankAbbreviation string
    Balance          int
}

// Bank struct to represent a Bank entity, with a Ledger for interbank transactions
type Bank struct {
	BankID       int
	FullName     string
	Abbreviation string
	IsActive     bool
    Ledger       []LedgerEntry // Ledger as a slice of LedgerEntry
}

// AllBanks will hold all the banks
var AllBanks []Bank

// NewBank factory function to create a new bank
func NewBank(fullName, abbreviation string) (*Bank, error) {
	// Validate the bank details
	if err := BankValidations(fullName, abbreviation); err != nil {
		return nil, err
	}

	// Create and return a new bank
	bank := &Bank{
		BankID:       getBankByID(),
		FullName:     fullName,
		Abbreviation: abbreviation,
		IsActive:     true,
		Ledger:       make(map[string]int), // Initialize an empty ledger
	}
	AllBanks = append(AllBanks, *bank)
	return bank, nil
}

// Getters and Setters for Bank attributes
func (b *Bank) GetBankID() int {
	return b.BankID
}

func (b *Bank) SetBankID(bankID int) {
	b.BankID = bankID
}

func (b *Bank) GetBankName() string {
	return b.FullName
}

func (b *Bank) SetBankName(fullName string) error {
	if err := validateBankName(fullName); err != nil {
		return err
	}
	b.FullName = fullName
	return nil
}

func (b *Bank) GetAbbreviation() string {
	return b.Abbreviation
}

func (b *Bank) SetAbbreviation(abbreviation string) error {
	if err := validateBankAbbreviation(abbreviation); err != nil {
		return err
	}
	b.Abbreviation = abbreviation
	return nil
}

func (b *Bank) GetIsActive() bool {
	return b.IsActive
}

func (b *Bank) SetIsActive(isActive bool) {
	b.IsActive = isActive
}

// Ledger Operations
// UpdateLedger updates the ledger of the bank with the given bank abbreviation and amount
func (b *Bank) UpdateLedger(bankAbbreviation string, amount int) error {
    if bankAbbreviation == b.Abbreviation {
        return errors.New("a bank cannot have a ledger entry for itself")
    }

    // Check if the entry for the other bank already exists
    for i, entry := range b.Ledger {
        if entry.BankAbbreviation == bankAbbreviation {
            b.Ledger[i].Balance += amount
            return nil
        }
    }

    // If entry doesn't exist, create a new one
    newEntry := LedgerEntry{
        BankAbbreviation: bankAbbreviation,
        Balance:          amount,
    }
    b.Ledger = append(b.Ledger, newEntry)
    return nil
}

// GetLedger returns the ledger of the bank
func (b *Bank) GetLedger() []LedgerEntry {
    return b.Ledger
}

// DisplayLedger prints the bank's ledger 
func (b *Bank) DisplayLedger() {
    fmt.Printf("Ledger for %s:\n", b.FullName)
    for _, entry := range b.Ledger {
        fmt.Printf("%s : %+d\n", entry.BankAbbreviation, entry.Balance)
    }
}

// LendTo allows the bank to lend money to a customer's account
func (b *Bank) LendTo(accountID int, amount float64, bankID int) error {
	// Find the customer account using the customer package
	customerAccount, err := customer.GetAccountByID(accountID)
	if err != nil {
		return err
	}

	// Create a new transaction for lending
	transaction, err := customer.NewTransaction("credit", amount, customerAccount.Balance+amount, accountID, bankID, time.Now())
	if err != nil {
		return err
	}

	// Add the transaction to the customer's passbook
	customerAccount.Passbook.AddTransaction(transaction)

	// Update the ledger for this bank
	err = b.UpdateLedger(customerAccount.BankID, int(amount))
	if err != nil {
		return err
	}

	return nil
}

// Admin CRUD Operations

// GetAllActiveBanks retrieves all active banks
func GetAllActiveBanks() []Bank {
	var activeBanks []Bank
	for _, bank := range AllBanks {
		if bank.IsActive {
			activeBanks = append(activeBanks, bank)
		}
	}
	return activeBanks
}

// DeleteBankByID sets a bank as inactive by ID
func DeleteBankByID(bankID int) error {
	for i, bank := range AllBanks {
		if bank.BankID == bankID {
			AllBanks[i].IsActive = false
			return nil
		}
	}
	return errors.New("bank not found")
}

// UpdateBankDetails allows admins to update bank details
func UpdateBankDetails(bankID int, parameter, value string) error {
	for i, bank := range AllBanks {
		if bank.BankID == bankID {
			if !bank.GetIsActive() {
				return errors.New("cannot update inactive bank")
			}

			switch parameter {
			case "FullName":
				return AllBanks[i].SetBankName(value)
			case "Abbreviation":
				return AllBanks[i].SetAbbreviation(value)
			default:
				return errors.New("invalid parameter to update")
			}
		}
	}
	return errors.New("bank not found")
}

// Helper functions and validations

// GetBankByID returns the next available BankID
func getBankByID() int {
	if len(AllBanks) == 0 {
		return 1
	}
	return AllBanks[len(AllBanks)-1].BankID + 1
}

// BankValidations performs validations on the bank's FullName and Abbreviation
func BankValidations(fullName, abbreviation string) error {
	if err := validateBankName(fullName); err != nil {
		return err
	}
	if err := validateBankAbbreviation(abbreviation); err != nil {
		return err
	}
	return nil
}

// Validate bank name (at least 3 characters, no numbers)
func validateBankName(name string) error {
	if len(name) < 3 {
		return errors.New("bank name must be at least 3 characters long")
	}
	if strings.ContainsAny(name, "0123456789") {
		return errors.New("bank name cannot contain numbers")
	}
	return nil
}

// Validate abbreviation (exactly 3 uppercase letters)
func validateBankAbbreviation(abbreviation string) error {
	if len(abbreviation) != 3 {
		return errors.New("abbreviation must be exactly 3 characters")
	}
	if abbreviation != strings.ToUpper(abbreviation) {
		return errors.New("abbreviation must be in uppercase letters")
	}
	return nil
}
