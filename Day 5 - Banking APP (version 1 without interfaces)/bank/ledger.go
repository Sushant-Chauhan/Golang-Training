package bank

import (
    "errors"
    "fmt"
)

// LedgerEntry represents an interbank transaction record.
type LedgerEntry struct {
    BankAbbreviation string
    Balance          int // +ve means this bank has to receive, -ve means this bank owes
}

// UpdateLedgerEntry updates or adds an entry in the bank's ledger.
func UpdateLedgerEntry(ledger *[]LedgerEntry, bankAbbreviation string, amount int) error {
    for i, entry := range *ledger {
        if entry.BankAbbreviation == bankAbbreviation {
            (*ledger)[i].Balance += amount
            return nil
        }
    }

    // If the entry for the bank doesn't exist, create a new one
    newEntry := LedgerEntry{
        BankAbbreviation: bankAbbreviation,
        Balance:          amount,
    }
    *ledger = append(*ledger, newEntry)
    return nil
}

// GetLedgerEntry retrieves a ledger entry for a specific bank.
func GetLedgerEntry(ledger []LedgerEntry, bankAbbreviation string) (*LedgerEntry, error) {
    for _, entry := range ledger {
        if entry.BankAbbreviation == bankAbbreviation {
            return &entry, nil
        }
    }
    return nil, errors.New("ledger entry not found")
}

// DisplayLedger prints all ledger entries.
func DisplayLedger(ledger []LedgerEntry) {
    if len(ledger) == 0 {
        fmt.Println("No ledger entries found.")
        return
    }

    fmt.Println("Ledger:")
    for _, entry := range ledger {
        balanceStatus := "owes"
        if entry.Balance > 0 {
            balanceStatus = "receives"
        }
        fmt.Printf("%s %s %+d\n", entry.BankAbbreviation, balanceStatus, entry.Balance)
    }
}
