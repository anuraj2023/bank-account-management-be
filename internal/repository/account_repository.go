package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/anuraj2023/bank-account-management-be/internal/models"
	"github.com/anuraj2023/bank-account-management-be/pkg/immudb"
)

var ErrAccountNotFound = errors.New("Account not found")

type AccountRepository interface {
	CreateAccount(ctx context.Context, account *models.Account) (*models.Account, error)
	GetAllAccounts(ctx context.Context) ([]*models.Account, error)
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

	err = r.client.Save(ctx, data)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r *accountRepository) GetAllAccounts(ctx context.Context) ([]*models.Account, error) {
	documents, err := r.client.GetAll(ctx, 1, 100) 
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.Account, 0, len(documents))
	for _, doc := range documents {
		data, err := json.Marshal(doc)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal document: %v", err)
		}

		var account models.Account
		err = json.Unmarshal(data, &account)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal document: %v", err)
		}
		accounts = append(accounts, &account)
	}

	return accounts, nil
}