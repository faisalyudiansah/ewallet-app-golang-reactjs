package repository

import (
	"context"

	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/database"
)

type TransactionTypeRepository interface {
	GetAll(ctx context.Context) ([]model.TransactionType, error)
}

type transactionTypeRepositoryPostgreSQL struct {
	db *database.GormWrapper
}

func NewTransactionTypeRepository(db *database.GormWrapper) *transactionTypeRepositoryPostgreSQL {
	return &transactionTypeRepositoryPostgreSQL{
		db: db,
	}
}

func (r *transactionTypeRepositoryPostgreSQL) GetAll(ctx context.Context) ([]model.TransactionType, error) {
	res := []model.TransactionType{}

	if err := r.db.Start(ctx).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
