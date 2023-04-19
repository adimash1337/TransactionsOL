package storage

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/models"
	"awesomeProject/internal/storage/postgre"
	"context"
)

type ITransactionRepository interface {
	Create(ctx context.Context, transaction *models.Transaction) (string, error)
	Get(ctx context.Context, ID string) (models.Transaction, error)
	List(ctx context.Context) ([]models.Transaction, error)
	Delete(ctx context.Context, ID string) error
}

type Storage struct {
	Transaction ITransactionRepository
}

func New(ctx context.Context, cfg *config.Config) (*Storage, error) {
	pgDb, err := postgre.Dial(*cfg)
	if err != nil {
		return nil, err
	}

	tRepo := postgre.NewTransactionRepository(pgDb)

	storage := Storage{
		Transaction: tRepo,
	}
	return &storage, nil
}
