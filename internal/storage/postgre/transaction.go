package postgre

import (
	"awesomeProject/internal/models"
	"context"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(ctx context.Context, transaction *models.Transaction) (string, error) {
	result := r.db.WithContext(ctx).Omit("deleted_at").Create(&transaction)
	return transaction.ID, result.Error
}

func (r *TransactionRepository) Get(ctx context.Context, ID string) (models.Transaction, error) {
	var trans models.Transaction
	err := r.db.WithContext(ctx).Where("id = ?", ID).First(&trans).Error
	if err != nil {
		return models.Transaction{}, err
	}

	return trans, nil
}

func (r *TransactionRepository) Delete(ctx context.Context, ID string) error {
	var res models.Transaction
	return r.db.WithContext(ctx).Where("id = ?", ID).Delete(&res).Error
}

func (r *TransactionRepository) List(ctx context.Context) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.WithContext(ctx).Find(&transactions)
	return transactions, err.Error
}
