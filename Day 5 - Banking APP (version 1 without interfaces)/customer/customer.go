package customer

import (
	"errors"
	"fmt"
	"strings"
	// "bankingApp/account" 
)

type Customer struct {
	CustomerID  int
	FirstName   string
	LastName    string
	TotalBalance float64
	IsAdmin     bool
	IsActive    bool
	Accounts    []account.Account // Customer has multiple accounts
}

// list of all customers
var AllCustomers []Customer

// Factory Functions -

// NewAdmin creates a new Admin 
func NewAdmin(firstName, lastName string) (*Customer, error) {
	if err := validateName(firstName, lastName); err != nil {
		return nil, err
	}
	admin := &Customer{
		CustomerID:  findCustomerID(),
		FirstName:   firstName,
		LastName:    lastName,
		TotalBalance: 0,
		IsAdmin:     true,
		IsActive:    true,
		Accounts:    []account.Account{},
	}
	AllCustomers = append(AllCustomers, *admin)
	return admin, nil
}

// NewCustomer creates a new customer (non-admin)
func NewCustomer(firstName, lastName string) (*Customer, error) {
	if err := validateName(firstName, lastName); err != nil {
		return nil, err
	}
	customer := &Customer{
		CustomerID:  findCustomerID(),
		FirstName:   firstName,
		LastName:    lastName,
		TotalBalance: 0,
		IsAdmin:     false,
		IsActive:    true,
		Accounts:    []account.Account{},
	}
	AllCustomers = append(AllCustomers, *customer)
	return customer, nil
}

// validations and finding next ID
func findCustomerID() int {
	if len(AllCustomers) == 0 {
		return 1
	}
	return AllCustomers[len(AllCustomers)-1].CustomerID + 1
}

func validateName(firstName, lastName string) error {
	if len(firstName) < 2 || len(lastName) < 2 {
		return errors.New("name must be at least 2 characters long")
	}
	if strings.ContainsAny(firstName, "0123456789") || strings.ContainsAny(lastName, "0123456789") {
		return errors.New("name cannot contain numbers")
	}
	return nil
}

// Read Functions

// readAllBanks retrieves all active banks (Admin only)
func (c *Customer) readAllBanks() ([]bank.Bank, error) {
	if !c.IsAdmin {
		return nil, errors.New("only admin can read all banks")
	}
	return bank.GetAllActiveBanks(), nil
}

// readAllCustomers retrieves all customers (Admin only)
func (c *Customer) readAllCustomers() ([]Customer, error) {
	if !c.IsAdmin {
		return nil, errors.New("only admin can read all customers")
	}
	return AllCustomers, nil
}

// readAllAccounts retrieves all accounts of a customer
func (c *Customer) readAllAccounts() ([]account.Account, error) {
	if !c.IsActive {
		return nil, errors.New("inactive customer cannot read accounts")
	}
	return c.Accounts, nil
}

// readAccountByID retrieves account by its ID
func (c *Customer) readAccountByID(accountID int) (*account.Account, error) {
	for _, account := range c.Accounts {
		if account.AccountID == accountID {
			return &account, nil
		}
	}
	return nil, errors.New("account not found")
}

// Update Functions

// UpdateCustomer allows Admin to update a customer’s details
func UpdateCustomer(customerID int, attribute string, newValue interface{}) error {
	for i, customer := range AllCustomers {
		if customer.CustomerID == customerID {
			if !customer.IsAdmin {
				return errors.New("only admin can update customer")
			}
			switch attribute {
			case "FirstName":
				if name, ok := newValue.(string); ok {
					AllCustomers[i].FirstName = name
				} else {
					return errors.New("invalid value for FirstName")
				}
			case "LastName":
				if name, ok := newValue.(string); ok {
					AllCustomers[i].LastName = name
				} else {
					return errors.New("invalid value for LastName")
				}
			case "IsActive":
				if isActive, ok := newValue.(bool); ok {
					AllCustomers[i].IsActive = isActive
				} else {
					return errors.New("invalid value for IsActive")
				}
			default:
				return errors.New("invalid attribute to update")
			}
			return nil
		}
	}
	return errors.New("customer not found")
}

// UpdateAccount allows Admin or the Customer to update an account’s details
func UpdateAccount(customerID, accountID int, attribute string, newValue interface{}) error {
	for _, customer := range AllCustomers {
		if customer.CustomerID == customerID {
			for i, account := range customer.Accounts {
				if account.AccountID == accountID {
					switch attribute {
					case "Balance":
						if balance, ok := newValue.(float64); ok {
							customer.Accounts[i].Balance = balance
							UpdateTotalBalance(customerID) // Update total balance
						} else {
							return errors.New("invalid value for Balance")
						}
					default:
						return errors.New("invalid attribute to update")
					}
					return nil
				}
			}
		}
	}
	return errors.New("account not found")
}


// UpdateTotalBalance recalculates the total balance of the customer
func UpdateTotalBalance(customerID int) {
	for i, customer := range AllCustomers {
		if customer.CustomerID == customerID {
			total := 0.0
			for _, account := range customer.Accounts {
				total += account.Balance
			}
			AllCustomers[i].TotalBalance = total
			return
		}
	}
}

// Delete Functions

// DeleteCustomer allows Admin to deactivate a customer
func DeleteCustomer(customerID int) error {
	for i, customer := range AllCustomers {
		if customer.CustomerID == customerID {
			if !customer.IsAdmin {
				return errors.New("only admin can delete customers")
			}
			AllCustomers[i].IsActive = false
			return nil
		}
	}
	return errors.New("customer not found")
}

// DeleteAccount deactivates an account by setting its balance to zero
func DeleteAccount(customerID, accountID int) error {
	for i, customer := range AllCustomers {
		if customer.CustomerID == customerID {
			if !customer.IsActive {
				return errors.New("inactive customer cannot delete account")
			}
			for j, account := range customer.Accounts {
				if account.AccountID == accountID {
					customer.Accounts[j].Balance = 0
					UpdateTotalBalance(customerID) // Recalculate total balance
					return nil
				}
			}
		}
	}
	return errors.New("account not found")
}

// Banking Methods (Deposit, Withdraw, Transfer) 

// deposit adds money to a specific customer by accountID
func (c *Customer) Deposit(accountID int, amount float64) error {
	if !c.IsActive {
		return errors.New("inactive customer cannot deposit")
	}
	for i, account := range c.Accounts {
		if account.AccountID == accountID {
			c.Accounts[i].Balance += amount
			UpdateTotalBalance(c.CustomerID)
			return nil
		}
	}
	return errors.New("account not found")
}

// withdraw removes money from a specific account
func (c *Customer) Withdraw(accountID int, amount float64) error {
	if !c.IsActive {
		return errors.New("inactive customer cannot withdraw")
	}
	for i, account := range c.Accounts {
		if account.AccountID == accountID {
			if account.Balance < amount {
				return errors.New("insufficient funds")
			}
			c.Accounts[i].Balance -= amount
			UpdateTotalBalance(c.CustomerID)
			return nil
		}
	}
	return errors.New("account not found")
}

// TransferBetweenOwnAccount transfers money between the customer's own accounts
func (c *Customer) TransferBetweenOwnAccount(fromAccountNo, toAccountNo int, amount float64) error {
	var fromAcc, toAcc *account.Account
	for i := range c.Accounts {
		if c.Accounts[i].AccountID == fromAccountNo {
			fromAcc = &c.Accounts[i]
		}
		if c.Accounts[i].AccountID == toAccountNo {
			toAcc = &c.Accounts[i]
		}
	}
	if fromAcc == nil || toAcc == nil {
		return errors.New("one or both accounts not found")
	}
	if fromAcc.Balance < amount {
		return errors.New("insufficient balance")
	}
	fromAcc.Balance -= amount
	toAcc.Balance += amount
	UpdateTotalBalance(c.CustomerID)
	return nil
}

// TransferToOtherAccount transfers money to another customer’s account
func (c *Customer) TransferToOtherAccount(fromAccountNo, toCustID, toAccountNo int, amount float64) error {
	var fromAcc *account.Account
	for i := range c.Accounts {
		if c.Accounts[i].AccountID == fromAccountNo {
			fromAcc = &c.Accounts[i]
		}
	}
	if fromAcc == nil {
		return errors.New("source account not found")
	}
	// Find the target customer and account
	for _, customer := range AllCustomers {
		if customer.CustomerID == toCustID {
			for _, acc := range customer.Accounts {
				if acc.AccountID == toAccountNo {
					if fromAcc.Balance < amount {
						return errors.New("insufficient balance")
					}
					fromAcc.Balance -= amount
					acc.Balance += amount
					UpdateTotalBalance(c.CustomerID)
					UpdateTotalBalance(toCustID)
					return nil
				}
			}
		}
	}
	return errors.New("target account not found")
}

	// // GetAllCustomers returns all customers (active and inactive)
	// func GetAllCustomers() []Customer {
	// 	return AllCustomers
	// }

	// GetAllAdmins returns all active admins
	func GetAllAdmins() []Customer {
		var admins []Customer
		for _, cust := range AllCustomers {
			if cust.IsAdmin && cust.IsActive {
				admins = append(admins, cust)
			}
		}
		return admins
	} 
}