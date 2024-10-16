package bank

type Ledger struct {
	currentBalanceOfBank float64
	balances             map[int]float64
}

func newLedger() *Ledger {
	tempBalanceSheet := make(map[int]float64)
	return &Ledger{
		currentBalanceOfBank: 0,
		balances:             tempBalanceSheet,
	}
}

func (ledger *Ledger) creditBalanceFrom(bankId int, amount float64) float64 {
	balance, alreadyPresent := ledger.balances[bankId]
	if !alreadyPresent {
		ledger.balances[bankId] = 0
		balance = 0
	}
	finalBalance := balance + amount
	ledger.balances[bankId] = finalBalance
	ledger.currentBalanceOfBank += amount
	return finalBalance
}

func (ledger *Ledger) debitBalanceTo(bankId int, amount float64) float64 {
	balance, alreadyPresent := ledger.balances[bankId]
	if !alreadyPresent {
		ledger.balances[bankId] = 0
		balance = 0
	}
	finalBalance := balance - amount
	ledger.balances[bankId] = finalBalance
	ledger.currentBalanceOfBank -= amount
	return finalBalance
}

func (ledger *Ledger) getCurrentBalance() float64 {
	return ledger.currentBalanceOfBank
}

func (ledger *Ledger) getBalanceEntryForBankId(bankId int) float64 {
	balance, alreadyPresent := ledger.balances[bankId]
	if !alreadyPresent {
		return 0
	}
	return balance
}
