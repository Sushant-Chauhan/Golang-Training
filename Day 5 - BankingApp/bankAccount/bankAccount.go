package bankAccount

import (
	"bankingApp/helper"
	"errors"
	"strconv"
)

type BankAccountInterface interface {
	CloseBankAccount() (float64, error)
	GetBalance() float64
	GetAccountNumber() int
	GetBankId() int
	GetCustomerId() int
	GetPassbook() *Passbook
	DepositMoney(depositAmount float64) error
	WithdrawMoney(withdrawAmount float64) error
	InitiateTransferMoneyTo(transferAmount float64, accountNumber int, bankId int, note string) (int, error)
	RefundUnsuccessfulTransfer(transactionId int) error
	TransferMoneyFrom(transferAmount float64, accountNumberFromWhichTransferInitiated int, bankIdFromWhichTransferInitiated int, note string) error
}

type bankAccount struct {
	accountNumber int
	balance       float64
	bankId        int
	customerId    int
	isActive      bool
	passbook      *Passbook
}

func NewBankAccount(
	accountNumber int,
	customerId int,
	initialBalance float64,
	bankId int) (*bankAccount, error) {

	err := helper.ValidateAll(
		validateBankId(bankId),
		validateAccountNumber(accountNumber),
		validateCustomerId(customerId))
	if err != nil {
		return nil, err
	}
	tempPassBook := newPassBook(accountNumber)
	tempBankAccount := &bankAccount{
		accountNumber: accountNumber,
		balance:       initialBalance,
		customerId:    customerId,
		bankId:        bankId,
		isActive:      true,
		passbook:      tempPassBook,
	}
	tempPassBook.addNewDepositToPassbook(initialBalance, tempBankAccount.GetBalance(), "Initial Balance")
	return tempBankAccount, nil
}

func (bankAccount *bankAccount) CloseBankAccount() (float64, error) {
	err := bankAccount.validateIfActive()
	if err != nil {
		return 0, err
	}
	bankAccount.isActive = false
	tempBalance := bankAccount.balance
	bankAccount.balance = 0
	return tempBalance, nil
}

func (account *bankAccount) GetBalance() float64 {
	return account.balance
}

func (account *bankAccount) GetAccountNumber() int {
	return account.accountNumber
}

func (account *bankAccount) GetCustomerId() int {
	return account.customerId
}

func (account *bankAccount) GetBankId() int {
	return account.bankId
}

func (account *bankAccount) GetPassbook() *Passbook {
	return account.passbook
}

func (account *bankAccount) DepositMoney(depositAmount float64) error {
	err1 := account.validateIfActive()
	if err1 != nil {
		return err1
	}
	err := validateDepositAmount(depositAmount)
	if err != nil {
		return err
	}
	tempNote := "Deposit Money at XYZ"
	account.creditToBalance(depositAmount)
	account.passbook.addNewDepositToPassbook(depositAmount, account.GetBalance(), tempNote)

	return nil

}

func (account *bankAccount) WithdrawMoney(withdrawAmount float64) error {

	err := helper.ValidateAll(
		account.validateIfActive(),
		validateWithdrawAmount(withdrawAmount),
		account.validateIfBalanceToWithdraw(withdrawAmount))
	if err != nil {
		return err
	}
	account.debitFromBalance(withdrawAmount)
	tempNote := "Amount Withdraw at XYZ"

	account.passbook.addNewWithdrawToPassbook(withdrawAmount, account.GetBalance(), tempNote)

	return nil

}

func (account *bankAccount) InitiateTransferMoneyTo(transferAmount float64, accountNumber int, bankId int, note string) (int, error) {

	err := helper.ValidateAll(
		account.validateIfActive(),
		validateWithdrawAmount(transferAmount),
		account.validateIfBalanceToWithdraw(transferAmount))
	if err != nil {
		return -1, err
	}

	account.debitFromBalance(transferAmount)
	transactionId := account.passbook.addNewTransferToPassbook(transferAmount, account.GetBalance(), DEBIT_ENTRY, bankId, accountNumber, note)
	return transactionId, nil

}

func (account *bankAccount) RefundUnsuccessfulTransfer(transactionId int) error {

	err1 := account.validateIfActive()
	if err1 != nil {
		panic(err1)
	}
	tempTransaction := account.passbook.getTransactionById(transactionId)

	account.creditToBalance(tempTransaction.transactionAmount)
	account.passbook.addNewDepositToPassbook(tempTransaction.transactionAmount, account.GetBalance(), "Transaction UnSuccessful : "+strconv.Itoa(transactionId))
	return nil

}

func (account *bankAccount) TransferMoneyFrom(transferAmount float64, accountNumberFromWhichTransferInitiated int, bankIdFromWhichTransferInitiated int, note string) error {

	err1 := account.validateIfActive()
	if err1 != nil {
		return err1
	}
	err := validateDepositAmount(transferAmount)
	if err != nil {
		return err
	}
	account.creditToBalance(transferAmount)
	account.passbook.addNewTransferToPassbook(transferAmount, account.GetBalance(), CREDIT_ENTRY, bankIdFromWhichTransferInitiated, accountNumberFromWhichTransferInitiated, note)

	return nil

}

// private Functions

func (account *bankAccount) creditToBalance(depositAmount float64) {
	newBalance := account.GetBalance() + depositAmount
	account.balance = newBalance
}

func (account *bankAccount) debitFromBalance(withdrawalAmount float64) {
	newBalance := account.GetBalance() - withdrawalAmount
	account.balance = newBalance
}

// validate
func validateAccountNumber(accountNumber int) error {
	if accountNumber < 0 {
		return errors.New("account Number Cannot be Negative")
	}
	return nil
}

func validateCustomerId(customerId int) error {
	if customerId < 0 {
		return errors.New("customerId Cannot be Negative")
	}
	return nil
}
func validateBankId(bankId int) error {
	if bankId < 0 {
		return errors.New("bank Id Cannot be Negative")
	}
	return nil
}

func validateDepositAmount(depositAmount float64) error {
	if depositAmount < 0 {
		return errors.New("deposit Amount Cannot be Negative")
	}
	if depositAmount == 0 {
		return errors.New("deposit Amount Cannot be Zero")
	}
	return nil
}
func validateWithdrawAmount(withdrawAmount float64) error {
	if withdrawAmount < 0 {
		return errors.New("withdraw Amount Cannot be Negative")
	}
	if withdrawAmount == 0 {
		return errors.New("withdraw Amount Cannot be Zero")
	}
	return nil
}

func (account *bankAccount) validateIfBalanceToWithdraw(withdrawAmount float64) error {
	if account.GetBalance() < withdrawAmount {
		return errors.New("Insufficient Balance to withdraw")
	}
	return nil
}

func (account *bankAccount) validateIfActive() error {
	if !account.isActive {
		return errors.New("Account Donot Exist")
	}
	return nil
}
