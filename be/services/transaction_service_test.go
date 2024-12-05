package services_test

import (
	"context"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/apperrors"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/dtos"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransactionUserServiceImplementation_GetListTransactionsUserService(t *testing.T) {
	tests := []struct {
		name           string
		totalCount     int64
		reqUserId      int64
		returnData     []models.TransactionUserAndSourceOfFund
		mockRepository func(*mocks.TransactionUserRepository, int64, []models.TransactionUserAndSourceOfFund, error)
		want           []dtos.ResponseSingleDataTransactionUser
		wantErr        error
	}{
		{
			name:       "should success get all list transactions user who is login",
			totalCount: int64(100),
			reqUserId:  int64(1),
			returnData: []models.TransactionUserAndSourceOfFund(nil),
			mockRepository: func(tur *mocks.TransactionUserRepository, userId int64, tuasof []models.TransactionUserAndSourceOfFund, err error) {
				tur.On("GetListTransactionsRepository", mock.Anything, userId, "frieren", "amount", "asc", int64(5), int64(1), "2024-01-01", "2024-02-02").Return([]models.TransactionUserAndSourceOfFund(nil), int64(100), nil)
			},
			want:    []dtos.ResponseSingleDataTransactionUser(nil),
			wantErr: nil,
		},
		{
			name:       "should failed get all list transactions user who is login",
			totalCount: int64(0),
			reqUserId:  int64(1),
			returnData: []models.TransactionUserAndSourceOfFund(nil),
			mockRepository: func(tur *mocks.TransactionUserRepository, userId int64, tuasof []models.TransactionUserAndSourceOfFund, err error) {
				tur.On("GetListTransactionsRepository", mock.Anything, userId, "frieren", "amount", "asc", int64(5), int64(1), "2024-01-01", "2024-02-02").Return(nil, int64(0), err)
			},
			want:    nil,
			wantErr: apperrors.ErrISE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTx := &mocks.TransactionRepository{}
			mockTxUser := &mocks.TransactionUserRepository{}
			mockUser := &mocks.UserRepository{}
			mockWallet := &mocks.WalletRepository{}
			txuService := services.NewTransactionUserServiceImplementation(mockTxUser, mockUser, mockWallet, mockTx)
			tt.mockRepository(mockTxUser, tt.reqUserId, tt.returnData, tt.wantErr)

			result, totalCount, err := txuService.GetListTransactionsUserService(context.Background(), tt.reqUserId, "frieren", "amount", "asc", int64(5), int64(1), "2024-01-01", "2024-02-02")
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.totalCount, totalCount)
		})
	}
}
