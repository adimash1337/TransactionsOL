package service

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/storage"
	"context"
	"errors"
)

type ITransactionService interface {
	Create(ctx context.Context, transaction *models.Transaction) (string, error)
	Delete(ctx context.Context, ID string) error
	Get(ctx context.Context, ID string) (models.Transaction, error)
	List(ctx context.Context) ([]models.Transaction, error)
}

type Manager struct {
	Transaction ITransactionService
}

func NewManager(storage *storage.Storage) (*Manager, error) {
	tSrv := NewTransactionService(storage)
	if storage == nil {
		return nil, errors.New("no storage provided")
	}

	return &Manager{
		Transaction: tSrv,
	}, nil
}
