package user

import (
	"bankingApp/bank"
	"bankingApp/bankAccount"
	"bankingApp/helper"
	"errors"
)

var allUsers []*User
var userIdCounter = 0

type AdminInterface interface {
	NewCustomerUser(firstName, lastName string, customerParameters ...interface{}) (*User, error)
	DeleteCustomer(userId int) (float64, error)
	GetUserById(userId int)
	NewBank(bankName string, abbrevation string) (*bank.BankInterface, error)
	GetAllBanks() []*bank.BankInterface
}

type StaffInterface interface {
	OpenNewBankAccount(bankId int) (bankAccount.BankAccountInterface, error)
	CloseBankAccount(bankId int, accountNumber int) error
	DepositMoney(accountNumber int, bankId int, amount float64) error
	WithdrawMoney(accountNumber int, bankId int, amount float64) error
	TransferMoneyTo(accountNumberFrom int, bankIdFrom int, amount float64, accountNumberTo int, bankIdTo int, note string) error
	GetTotalBalance() (float64, error)
	GetFullName() string
	GetUserId() int
}
type User struct {
	userId    int
	isAdmin   bool
	firstName string
	lastName  string
	customer  *Customer
}

func NewAdminUser(firstName, lastName string) (*User, error) {
	firstName = helper.RemoveAllLeadingAndTrailingSpaces(firstName)
	lastName = helper.RemoveAllLeadingAndTrailingSpaces(lastName)
	err := helper.ValidateAll(
		validateFirstName(firstName),
		validateLastName(lastName))
	if err != nil {
		return nil, err
	}
	tempUserObject := &User{
		userId:    userIdCounter,
		isAdmin:   true,
		firstName: firstName,
		lastName:  lastName,
		customer:  nil,
	}
	userIdCounter++
	allUsers = append(allUsers, tempUserObject)
	return tempUserObject, nil
}

// Admin Function
func (user *User) NewCustomerUser(firstName, lastName string, customerParameters ...interface{}) (*User, error) {

	firstName = helper.RemoveAllLeadingAndTrailingSpaces(firstName)
	lastName = helper.RemoveAllLeadingAndTrailingSpaces(lastName)
	err := helper.ValidateAll(
		validateIfAdmin(user),
		validateFirstName(firstName),
		validateLastName(lastName))
	if err != nil {
		return nil, err
	}
	tempUserObject := &User{
		userId:    userIdCounter,
		isAdmin:   false,
		firstName: firstName,
		lastName:  lastName,
		customer:  newCustomer(customerParameters),
	}
	userIdCounter++
	return tempUserObject, nil
}

//admin functions

// Admin Function
func (user *User) NewBank(bankName string, abbrevation string) (bank.BankInterface, error) {
	err := validateIfAdmin(user)
	if err != nil {
		return nil, err
	}
	return bank.NewBank(bankName, abbrevation)
}

// Admin Function
func (user *User) DeleteCustomer(userId int) (float64, error) {
	err := validateIfAdmin(user)
	if err != nil {
		return 0, err
	}
	customerToDelete := allUsers[userId]
	err2 := validateIfCustomer(customerToDelete)
	if err2 != nil {
		return 0, err2
	}

	return customerToDelete.customer.deleteCustomer()
}

// general
func GetAllBanks() []bank.BankInterface {
	return bank.GetAllBanks()
}

// customer Functions
// staff function
func (user *User) OpenNewBankAccount(bankId int) (bankAccount.BankAccountInterface, error) {
	err := validateIfCustomer(user)
	if err != nil {
		return nil, err
	}
	return user.customer.openNewBankAccount(bankId, user.userId)
}

// staff function
func (user *User) CloseBankAccount(bankId int, accountNumber int) error {
	err := validateIfCustomer(user)
	if err != nil {
		return err
	}
	return user.customer.closeBankAccount(bankId, user.userId, accountNumber)
}

func (user *User) WithdrawMoney(accountNumber int, bankId int, amount float64) error {
	err := validateIfCustomer(user)
	if err != nil {
		return err
	}
	return user.customer.withdrawMoney(accountNumber, bankId, amount)
}
func (user *User) DepositMoney(accountNumber int, bankId int, amount float64) error {
	err := validateIfCustomer(user)
	if err != nil {
		return err
	}
	return user.customer.depositMoney(accountNumber, bankId, amount)

}

func (user *User) TransferMoneyTo(accountNumberFrom int, bankIdFrom int, amount float64, accountNumberTo int, bankIdTo int, note string) error {
	err := validateIfCustomer(user)

	if err != nil {
		return err
	}
	return user.customer.transferMoney(accountNumberFrom, bankIdFrom, amount, accountNumberTo, bankIdTo, note)
}

func (user *User) GetTotalBalance() (float64, error) {
	err := validateIfCustomer(user)
	if err != nil {
		return 0, err
	}
	return user.customer.getBalance(), nil
}

// validations

func validateInputCustomerId(userId int) error {
	if userId < 0 || userIdCounter >= userId {
		return errors.New("Invalid User Id")
	}
	return nil
}
func validateFirstName(firstName string) error {
	if len(firstName) < 2 {
		return errors.New("length of First Name should atleast be 2")
	}
	return nil
}
func validateLastName(lastName string) error {
	if len(lastName) == 0 {
		return errors.New("last Name Cannot be Empty")
	}
	return nil
}

func validateIfAdmin(user *User) error {

	if !user.isUserAdmin() {
		return errors.New("unauthoried Access")
	}
	return nil
}

// TODO : isStaff ->SRP
func validateIfCustomer(user *User) error {

	if user.isUserAdmin() {
		return errors.New("Only Customers can access this")
	}
	return nil
}

// getters
func (user *User) GetUserId() int {
	return user.userId
}
func (user *User) isUserAdmin() bool {
	return user.isAdmin
}
func (user *User) GetFullName() string {
	return user.firstName + " " + user.lastName
}

func (user *User) GetUserById(userId int) (*User, error) {
	err := helper.ValidateAll(
		validateIfAdmin(user),
		validateInputCustomerId(userId))
	if err != nil {
		return nil, err
	}

	return allUsers[userId], nil

}
