package models

import (
	"errors"
	"regexp"
)

type AccountType string

const (
	AccountTypeSending   AccountType = "sending"
	AccountTypeReceiving AccountType = "receiving"
)

type Account struct {
	AccountNumber string      `json:"account_number"`
	AccountName   string      `json:"account_name"`
	IBAN          string      `json:"iban"`
	Address       string      `json:"address"`
	Amount        float64     `json:"amount"`
	Type          AccountType `json:"type"`
}

func (a *Account) Validate() error {
	if a.AccountNumber == "" {
		return errors.New("account number is required")
	}
	if a.AccountName == "" {
		return errors.New("account name is required")
	}
	if !isValidIBAN(a.IBAN) {
		return errors.New("invalid IBAN")
	}
	if a.Address == "" {
		return errors.New("address is required")
	}
	if a.Amount < 0 {
		return errors.New("amount must be non-negative")
	}
	if a.Type != AccountTypeSending && a.Type != AccountTypeReceiving {
		return errors.New("invalid account type")
	}
	return nil
}

// This is very basic IBAN validation, just for demo purpose
func isValidIBAN(iban string) bool {
	match, _ := regexp.MatchString(`^[A-Z]{2}\d{2}[A-Z0-9]{11,30}$`, iban)
	return match
}