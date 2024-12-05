package services

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/apperrors"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/dtos"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/helpers"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/repositories"
)

type TransactionServiceInterface interface {
	GetListTransactionsUserService(context.Context, int64, string, string, string, int64, int64, string, string) ([]dtos.ResponseSingleDataTransactionUser, int64, error)
}

type TransactionUserServiceImplementation struct {
	TransactionUserRepository repositories.TransactionUserRepository
	UserRepository            repositories.UserRepository
	WalletRepository          repositories.WalletRepository
	TransactionsRepository    repositories.TransactionRepository
}

func NewTransactionUserServiceImplementation(
	txu repositories.TransactionUserRepository,
	us repositories.UserRepository,
	w repositories.WalletRepository,
	tx repositories.TransactionRepository,
) *TransactionUserServiceImplementation {
	return &TransactionUserServiceImplementation{
		TransactionUserRepository: txu,
		UserRepository:            us,
		WalletRepository:          w,
		TransactionsRepository:    tx,
	}
}

func (ts *TransactionUserServiceImplementation) GetListTransactionsUserService(ctx context.Context, userId int64, query string, sortBy string, sort string, limit int64, offset int64, startDate string, endDate string) ([]dtos.ResponseSingleDataTransactionUser, int64, error) {
	transactions, totalCount, err := ts.TransactionUserRepository.GetListTransactionsRepository(ctx, userId, query, sortBy, sort, limit, offset, startDate, endDate)
	if err != nil {
		return nil, 0, apperrors.ErrISE
	}
	return helpers.FormatterTransactionsList(transactions), totalCount, nil
}
