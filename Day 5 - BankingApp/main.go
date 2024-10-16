package main

import (
	"bankingApp/user"
	"fmt"
)

type Shape interface {
	abc()
}
type Circle struct {
}

func (c *Circle) abc() {

}

func main() {

	admin, err := user.NewAdminUser("Amdin", "Amind")
	if err != nil {
		panic(err)
	}
	c1, err2 := admin.NewCustomerUser("Tanmie", "Kamat")
	if err != nil {
		panic(err2)
	}
	c2, err2 := admin.NewCustomerUser("Ash", "Kamat")
	if err != nil {
		panic(err2)
	}
	bk, _ := admin.NewBank("Ktk", "KTK")

	acc, err4 := c1.OpenNewBankAccount(bk.GetId())
	fmt.Println(c1.GetTotalBalance())
	err6 := c1.DepositMoney(acc.GetAccountNumber(), acc.GetBankId(), 152)
	if err6 != nil {
		panic(err6)
	}

	fmt.Println(c1.GetTotalBalance())
	err7 := c1.WithdrawMoney(acc.GetAccountNumber(), acc.GetBankId(), 120)
	if err7 != nil {
		panic(err7)
	}
	fmt.Println(c1.GetTotalBalance())

	passbook1 := acc.GetPassbook()
	fmt.Println(passbook1.GetAllTransactionsAsString())

	acc2, err5 := c2.OpenNewBankAccount(bk.GetId())
	fmt.Println(c2.GetTotalBalance())

	errx := c1.TransferMoneyTo(acc.GetAccountNumber(), acc.GetBankId(), 100, acc2.GetAccountNumber(), acc2.GetBankId(), "Transfer from First account to second")
	if errx != nil {
		panic(errx)
	}

	passbook2 := acc2.GetPassbook()
	fmt.Println(passbook2.GetAllTransactionsAsString())
	passbook3 := acc.GetPassbook()
	fmt.Println(passbook3.GetAllTransactionsAsString())
	if err4 != nil || err5 != nil {
		panic(err4)
	}
}
