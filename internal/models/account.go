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
    AccountNumber string      `json:"acc_number" example:"1234567890"`
    AccountName   string      `json:"acc_name" example:"Tom Cruise"`
    IBAN          string      `json:"iban" example:"DE89370400440532013000"`
    Address       string      `json:"address" example:"123 Becker Str, Berlin, DE 12345"`
    Amount        float64     `json:"amount" example:"1000.50"`
    Type          AccountType `json:"type" example:"sending" enums:"sending,receiving"`
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