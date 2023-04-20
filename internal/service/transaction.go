package service

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/storage"
	pb "awesomeProject/proto"
	"context"
	uuid "github.com/google/uuid"
	"time"
)

type TransactionService struct {
	repo *storage.Storage
}

func NewTransactionService(repo *storage.Storage) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (s *TransactionService) Create(ctx context.Context, transaction *models.Transaction, in pb.Transaction) (string, error) {
	transaction.ID = uuid.NewString()
	transaction.ItemID = string(in.ItemId)
	transaction.UserID = string(in.UserId)
	transaction.PaymentAmount = float64(in.Price)
	transaction.PaymentTime = time.Now()
	return s.repo.Transaction.Create(ctx, transaction)
}

func (s *TransactionService) Delete(ctx context.Context, ID string) error {
	return s.repo.Transaction.Delete(ctx, ID)
}

func (s *TransactionService) Get(ctx context.Context, ID string) (models.Transaction, error) {
	return s.repo.Transaction.Get(ctx, ID)
}

func (s *TransactionService) List(ctx context.Context) ([]models.Transaction, error) {
	return s.repo.Transaction.List(ctx)
}
