package bank

import (
	"bankingApp/bankAccount"
	"bankingApp/helper"
	"errors"
)

var allBanks []BankInterface
var bankIdCounter = 0

// Interface For Bank with Other Class as Bank Implementation -> This Helps to Utilises DIP
// TODO: bank with ledger was to be composed

type BankInterface interface {
	OpenNewBankAccount(customerId int) (bankAccount.BankAccountInterface, error)
	CloseBankAccount(accountNumber int, customerId int) (float64, error)
	SendMoneyToAnotherBank(receiverBank *Bank, amount float64) error
	GetBankBalance() float64
	GetBalanceEntryForBankId(bankId int) (float64, error)
	TransferMoneyFrom(accountNumberTo int, bankIdTo int, amount float64, accountNumberFrom int, bankIdFrom int, note string) error
	GetId() int
	IsBankActive() bool
}

type Bank struct {
	bankId       int
	isActive     bool
	bankName     string
	abbreviation string
	accounts     []bankAccount.BankAccountInterface
	ledger       *Ledger
}

type BankOnline struct {
	*Bank
	onlineData string
}

func NewBank(bankName string, bankAbbreviation string) (*Bank, error) {
	bankAbbreviation = helper.RemoveAllLeadingAndTrailingSpaces(bankAbbreviation)
	bankName = helper.RemoveAllLeadingAndTrailingSpaces(bankName)
	err := helper.ValidateAll(
		validateBankAbbreviation(bankAbbreviation),
		validateBankName(bankName))
	if err != nil {
		return nil, err
	}
	tempLedgerObject := newLedger()
	var tempEmptyAccounts []bankAccount.BankAccountInterface
	tempBankObj := &Bank{
		bankId:       bankIdCounter,
		bankName:     bankName,
		abbreviation: bankAbbreviation,
		isActive:     true,
		accounts:     tempEmptyAccounts,
		ledger:       tempLedgerObject,
	}
	allBanks = append(allBanks, tempBankObj)
	bankIdCounter++
	return tempBankObj, nil
}

// functions
func (bank *Bank) OpenNewBankAccount(customerId int) (bankAccount.BankAccountInterface, error) {

	nextAccountNumber := bank.getNextAccountNumber()
	var tempBankAccount bankAccount.BankAccountInterface

	tempBankAccount, err := bankAccount.NewBankAccount(nextAccountNumber, customerId, 1000, bank.bankId)
	if err != nil {
		return nil, err
	}
	bank.accounts = append(bank.accounts, tempBankAccount)
	return tempBankAccount, nil

}

func (bank *Bank) CloseBankAccount(accountNumber int, customerId int) (float64, error) {
	bankAccount, err := bank.getBankAccountByNumber(accountNumber)
	if err != nil {
		return 0, err
	}
	err1 := validateIfCustomerHasAccessToBankAccount(bankAccount, customerId)
	if err1 != nil {
		return 0, err
	}

	balance, err2 := bankAccount.CloseBankAccount()

	return balance, err2
}

func (bank *Bank) SendMoneyToAnotherBank(receiverBank *Bank, amount float64) error {
	err := validateIfAmountIsGreaterThanZero(amount)
	if err != nil {
		return err
	}
	bank.ledger.debitBalanceTo(receiverBank.bankId, amount)
	receiverBank.ledger.creditBalanceFrom(bank.bankId, amount)
	return nil

}

func (bank *Bank) GetBankBalance() float64 {
	return bank.ledger.getCurrentBalance()
}

func (bank *Bank) GetBalanceEntryForBankId(bankId int) (float64, error) {
	err := validateBankId(bankId)
	if err != nil {
		return 0, err
	}
	return bank.ledger.getBalanceEntryForBankId(bankId), nil
}

// getters
func GetAllBanks() []BankInterface {
	var tempBankArray []BankInterface
	for _, bank := range allBanks {
		if !bank.IsBankActive() {
			continue
		}
		tempBankArray = append(tempBankArray, bank)
	}
	return tempBankArray
}

func GetBankById(bankId int) (BankInterface, error) {
	err := validateBankId(bankId)
	if err != nil {
		return nil, err
	}
	for _, bank := range allBanks {
		if !bank.IsBankActive() {
			continue
		}
		if bank.GetId() == bankId {
			return bank, nil
		}
	}
	return nil, errors.New("Bank Not Found")
}

func (bank *Bank) TransferMoneyFrom(accountNumberTo int, bankIdTo int, amount float64, accountNumberFrom int, bankIdFrom int, note string) error {
	bankAccount, err := bank.getBankAccountByNumber(accountNumberTo)
	if err != nil {
		return err
	}

	err2 := (bankAccount).TransferMoneyFrom(amount, accountNumberFrom, bankIdFrom, note)

	return err2
}

func (bank *Bank) GetId() int {
	return bank.bankId
}

func (bank *Bank) getBankAccountByNumber(accountNumber int) (bankAccount.BankAccountInterface, error) {
	for _, bankAccount := range bank.getAllBankAccountsForTheBank() {
		if bankAccount.GetAccountNumber() == accountNumber {
			return bankAccount, nil
		}
	}
	return nil, errors.New("Bank Account not found")

}

func (bank *Bank) getAllBankAccountsForTheBank() []bankAccount.BankAccountInterface {
	return bank.accounts
}
func (b *Bank) getNextAccountNumber() int {
	if len(b.accounts) == 0 {
		return 1
	}
	return b.accounts[len(b.accounts)-1].GetAccountNumber() + 1
}

func (b *Bank) IsBankActive() bool {
	return b.isActive
}

// validations
func validateBankId(bankId int) error {
	if bankId < 0 || bankId >= bankIdCounter {
		return errors.New("Invalid BankId")
	}
	return nil
}
func validateBankName(bankName string) error {
	if len(bankName) < 2 {
		return errors.New("length of BankName should atleast be 2")
	}
	return nil
}
func validateBankAbbreviation(abbreviation string) error {
	if len(abbreviation) == 0 {
		return errors.New("Abbreviation Cannot be Empty")
	}
	return nil
}

func validateIfCustomerHasAccessToBankAccount(account bankAccount.BankAccountInterface, customerId int) error {
	if account.GetCustomerId() != customerId {
		return errors.New("Your are Not Authorized For this Action")
	}
	return nil
}

func validateIfAmountIsGreaterThanZero(amount float64) error {
	if amount < 0 {
		return errors.New("Amount Cannot be Negative")
	}
	return nil
}
