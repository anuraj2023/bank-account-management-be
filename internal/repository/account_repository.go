package repository

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"log"

	"github.com/anuraj2023/bank-account-management-be/internal/models"
	"github.com/anuraj2023/bank-account-management-be/pkg/immudb"
)

var ErrAccountNotFound = errors.New("Account not found")

type AccountRepository interface {
	CreateAccount(ctx context.Context, account *models.Account) (*models.Account, error)
	GetAccount(ctx context.Context, accountNumber string) (*models.Account, error)
	ListAccounts(ctx context.Context) ([]*models.Account, error)
}

type accountRepository struct {
	client *immudb.Client
}

func NewAccountRepository(client *immudb.Client) AccountRepository {
	return &accountRepository{client: client}
}

func (r *accountRepository) CreateAccount(ctx context.Context, account *models.Account) (*models.Account, error) {
	data, err := json.Marshal(account)
	if err != nil {
		return nil, err
	}

	err = r.client.Set(ctx, account.AccountNumber, data)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r *accountRepository) GetAccount(ctx context.Context, accountNumber string) (*models.Account, error) {
	data, err := r.client.Get(ctx, accountNumber)
	log.Printf("GetAccount error: %v", err)
	if err != nil {
		if strings.Contains(err.Error(), "key not found") {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}

	var account models.Account
	err = json.Unmarshal(data, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}


// TO-DO: Pagination
func (r *accountRepository) ListAccounts(ctx context.Context) ([]*models.Account, error) {
	data, err := r.client.Scan(ctx, "")
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.Account, 0, len(data))
	for _, v := range data {
		var account models.Account
		err = json.Unmarshal(v, &account)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}

	return accounts, nil
}