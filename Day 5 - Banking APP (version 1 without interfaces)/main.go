package main

import (
	"fmt"
	"bankingApp/account"
	"bankingApp/customer" // Import the customer package
	// "bankingApp/transactions"
	"time"
)

func main() {
	// Creating Admins (customers with admin rights)
	adminCustomer1, _ := customer.NewAdmin("Rajesh", "Sharma")
	adminCustomer2, _ := customer.NewAdmin("Amit", "Verma")

	// Creating Customers via Admins
	customer1, _ := adminCustomer1.NewCustomer("Sushant", "Chauhan")
	customer2, _ := adminCustomer1.NewCustomer("Varun", "Singh")
	customer3, _ := adminCustomer1.NewCustomer("Arun", "Sharma")

	// Creating Banks
	sbi, _ := bank.NewBank("State Bank of India", "SBI")
	icici, _ := bank.NewBank("ICICI Bank", "ICICI")

	// Creating Accounts for Customers
	account1, _ := account.NewAccount(5000, sbi.GetBankID()) // Account for Sushant
	account2, _ := account.NewAccount(7000, icici.GetBankID()) // Account for Varun
	account3, _ := account.NewAccount(10000, sbi.GetBankID()) // Account for Arun

	// Adding Accounts to Customers
	customer1.SetAccount(account1)
	customer2.SetAccount(account2)
	customer3.SetAccount(account3)

	// Deposit Balance for Customer1 (Sushant)
	fmt.Println("\nDepositing Rs. 2000 to Sushant's Account")
	err := account1.Deposit(2000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("New Balance for Sushant: Rs. %.2f\n", account1.GetBalance())
	}

	// Withdraw Balance for Customer2 (Varun)
	fmt.Println("\nWithdrawing Rs. 1000 from Varun's Account")
	err = account2.Withdraw(1000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("New Balance for Varun: Rs. %.2f\n", account2.GetBalance())
	}

	// Printing Passbook for Customer1 (Sushant)
	fmt.Println("\nPrinting Passbook for Sushant:")
	account1.GetPassbook().PrintPassbook()

	// Printing Passbook for Customer2 (Varun)
	fmt.Println("\nPrinting Passbook for Varun:")
	account2.GetPassbook().PrintPassbook()

	// Printing All Active Customers
	fmt.Println("\nList of All Active Customers:")
	for _, cust := range customer.GetAllCustomers() {
		if cust.GetIsActive() {
			fmt.Printf("Customer: %s %s\n", cust.GetFirstName(), cust.GetLastName())
		}
	}
	// Printing All Active Admins
	fmt.Println("\nList of All Active Admins:")
	for _, adminCust := range customer.GetAllAdmins() {
		fmt.Printf("Admin: %s %s\n", adminCust.GetFirstName(), adminCust.GetLastName())
	}

	// Ledger - interbank transactions
	// Update ledgers
	sbi.UpdateLedger("ICICI", 100)  // SBI expects to receive 100 from ICICI
	sbi.UpdateLedger("BOI", -5000)  // SBI owes 5000 to BOI

	// Printing Ledger for SBI
	fmt.Println("\nLedger for SBI:")
	sbi.DisplayLedger()

	// Printing Ledger for ICICI
	fmt.Println("\nLedger for ICICI:")
	icici.DisplayLedger()

	// Retrieve specific ledger entry
	entry, _ := sbi.GetLedgerEntry("ICICI")
	fmt.Printf("SBI balance with ICICI: %+d\n", entry.Balance)

	// Printing Total Balance for all Customers
	totalBalance := 0.0
	fmt.Println("\nTotal Balance of All Accounts:")
	for _, acc := range account.GetAllAccounts() {
		fmt.Printf("Account %d (Customer ID %d): Rs. %.2f\n", acc.GetAccountNumber(), acc.GetAccountNumber(), acc.GetBalance())
		totalBalance += acc.GetBalance()
	}
	fmt.Printf("Total Balance of All Customers: Rs. %.2f\n", totalBalance)
}
