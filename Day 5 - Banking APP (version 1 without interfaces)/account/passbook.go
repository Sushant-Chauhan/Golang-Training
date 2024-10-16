package account

// I have included this is account.go (earlier in customer) - to avoid complexity and cycle not allowed error
// import (
// 	"fmt"
// 	"time"
// )

// type Passbook struct {
// 	transactions []*Transaction
// }

// // Factory function to create a new Passbook with an initial transaction
// func NewPassbook(initialBalance float64, accountID int, bankID int) (*Passbook, error) {
// 	transaction, err := NewTransaction("credit", initialBalance, initialBalance, accountID, bankID, time.Now())
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Initialize the Passbook with the initial transaction
// 	return &Passbook{
// 		transactions: []*Transaction{transaction},
// 	}, nil
// }

// // Getter for transactions
// func (p *Passbook) GetTransactions() []*Transaction {
// 	return p.transactions
// }

// // Setter for transactions
// func (p *Passbook) SetTransactions(transactions []*Transaction) {
// 	p.transactions = transactions
// }

// // AddTransaction method to add a new transaction to the passbook
// func (p *Passbook) AddTransaction(t *Transaction) {
// 	p.transactions = append(p.transactions, t)
// }

// // PrintPassbook method to print all transactions in the passbook
// func (p *Passbook) PrintPassbook() {
// 	if len(p.GetTransactions()) == 0 {
// 		fmt.Println("No transactions available in the passbook.")
// 		return
// 	}

// 	fmt.Println("Passbook Transactions:")
// 	for _, transaction := range p.transactions {
// 		transaction.PrintTransaction()
// 	}
// }
