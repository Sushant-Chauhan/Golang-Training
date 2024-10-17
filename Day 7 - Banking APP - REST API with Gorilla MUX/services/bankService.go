package services

import "errors"

type Bank struct {
	ID           int
	FullName     string
	Abbreviation string
	IsActive     bool
}

var banks []Bank

func CreateBank(bank Bank) (Bank, error) {
	// Validate bank fields
	if err := validateBank(bank); err != nil {
		return Bank{}, err
	}
	bank.ID = findBankID()
	banks = append(banks, bank)

	return bank, nil
}

func validateBank(bank Bank) error {
	if bank.FullName == "" {
		return errors.New("full name cannot be empty")
	}
	if bank.Abbreviation == "" {
		return errors.New("abbreviation cannot be empty")
	}
	return nil
}

//next unique bank ID.
func findBankID() int {
	if len(banks) == 0 {
		return 1 // Start IDs at 1
	}
	return banks[len(banks)-1].ID + 1
}

func GetBankByID(id int) (Bank, error) {
	for _, bank := range banks {
		if bank.ID == id {
			return bank, nil
		}
	}
	return Bank{}, errors.New("Bank not found")
}

func UpdateBankByID(id int, updatedBank Bank) error {
	for i, bank := range banks {
		if bank.ID == id {
			banks[i] = updatedBank
			return nil
		}
	}
	return errors.New("Bank not found")
}

func DeleteBankByID(id int) error {
	for i, bank := range banks {
		if bank.ID == id {
			banks = append(banks[:i], banks[i+1:]...)
			return nil
		}
	}
	return errors.New("Bank not found")
}

func GetAllBanks() []Bank {
	return banks
}
