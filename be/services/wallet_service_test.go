package services_test

import (
	"context"
	"testing"
	"time"

	"ewallet-server-v1/apperrors"
	"ewallet-server-v1/dtos"
	"ewallet-server-v1/mocks"
	"ewallet-server-v1/models"
	"ewallet-server-v1/services"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var reqTopUp = dtos.RequestTopUpWallet{
	Amount:       float64(10000000),
	SourceOfFund: int64(1),
}

var reqTransfer = dtos.RequestTransferFund{
	ToWalletNumber: "7770000000001",
	Amount:         float64(12345),
	Description:    "mantap",
}

var modelSof = &models.SourceOfFund{
	ID:        int64(1),
	Name:      "Bank Transfer",
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
	DeleteAt:  &time.Time{},
}

var modelTx = &models.Transaction{
	ID:              int64(1),
	SourceId:        int64(1),
	RecipientId:     int64(1),
	TransactionTime: time.Time{},
	Amount:          decimal.NewFromFloat(float64(10000000)),
	Description:     "note",
	CreatedAt:       time.Time{},
	UpdatedAt:       time.Time{},
	DeleteAt:        &time.Time{},
}

func TestWalletServiceImplementation_PutTopWalletService(t *testing.T) {
	tests := []struct {
		name           string
		reqUserId      int64
		reqBody        dtos.RequestTopUpWallet
		mockRepository func(*mocks.TransactionRepository, *mocks.TransactionUserRepository, *mocks.SourceOfFundRepository, *mocks.UserRepository, *mocks.WalletRepository, int64, *models.UserAndWallet, error)
		returnAtomic   *models.UserAndWallet
		want           *dtos.ResponseUserAndWallet
		wantErr        error
	}{
		{
			name:      "should return error user try to use source of fund id 4",
			reqUserId: int64(1),
			reqBody: dtos.RequestTopUpWallet{
				Amount:       float64(12345),
				SourceOfFund: int64(4),
			},
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrForbiddenAccess,
		},
		{
			name:      "should return error when failed on Get User By Id",
			reqUserId: int64(1),
			reqBody:   reqTopUp,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrInvalidAccessToken,
		},
		{
			name:      "should return error user input invalid id of source of fund",
			reqUserId: int64(1),
			reqBody: dtos.RequestTopUpWallet{
				Amount:       float64(12345),
				SourceOfFund: int64(2024),
			},
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser, nil)
				sofr.On("GetSourceOfFundById", mock.Anything, int64(2024)).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrSOFIdIsNotExists,
		},
		{
			name:      "should return error when failed get wallet in system",
			reqUserId: int64(1),
			reqBody:   reqTopUp,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser, nil)
				sofr.On("GetSourceOfFundById", mock.Anything, reqTopUp.SourceOfFund).Return(modelSof, nil)
				wr.On("GetWalletByIdUser", mock.Anything, modelUser.ID).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrInvalidAccessToken,
		},
		{
			name:      "should return error when failed change balance wallet in system",
			reqUserId: int64(1),
			reqBody:   reqTopUp,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser, nil)
				sofr.On("GetSourceOfFundById", mock.Anything, reqTopUp.SourceOfFund).Return(modelSof, nil)
				wr.On("GetWalletByIdUser", mock.Anything, modelUser.ID).Return(modelWallet, nil)
				wr.On("PutChangeBalanceWallet", mock.Anything, reqTopUp.Amount, modelUser.ID).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrISE,
		},
		{
			name:      "should return error when failed record new transaction in system",
			reqUserId: int64(1),
			reqBody:   reqTopUp,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser, nil)
				sofr.On("GetSourceOfFundById", mock.Anything, reqTopUp.SourceOfFund).Return(modelSof, nil)
				wr.On("GetWalletByIdUser", mock.Anything, modelUser.ID).Return(modelWallet, nil)
				wr.On("PutChangeBalanceWallet", mock.Anything, reqTopUp.Amount, modelUser.ID).Return(modelWallet, nil)
				tur.On("PostNewTransaction", mock.Anything, modelUser.ID, modelSof.ID, reqTopUp.Amount, modelSof.Name, "").Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrISE,
		},
		{
			name:      "should return error when failed create transaction user (pivot) in system",
			reqUserId: int64(1),
			reqBody:   reqTopUp,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser, nil)
				sofr.On("GetSourceOfFundById", mock.Anything, reqTopUp.SourceOfFund).Return(modelSof, nil)
				wr.On("GetWalletByIdUser", mock.Anything, modelUser.ID).Return(modelWallet, nil)
				wr.On("PutChangeBalanceWallet", mock.Anything, reqTopUp.Amount, modelUser.ID).Return(modelWallet, nil)
				tur.On("PostNewTransaction", mock.Anything, modelUser.ID, modelSof.ID, reqTopUp.Amount, modelSof.Name, "").Return(modelTx, nil)
				tur.On("PostTransactionUserPivot", mock.Anything, modelUser.ID, modelTx.ID).Return(err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrISE,
		},
		{
			name:      "should success add chance attempt game for the user if user top up 10000000",
			reqUserId: int64(1),
			reqBody: dtos.RequestTopUpWallet{
				Amount:       float64(10000000),
				SourceOfFund: int64(1),
			},
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser, nil)
				sofr.On("GetSourceOfFundById", mock.Anything, reqTopUp.SourceOfFund).Return(modelSof, nil)
				wr.On("GetWalletByIdUser", mock.Anything, modelUser.ID).Return(modelWallet, nil)
				wr.On("PutChangeBalanceWallet", mock.Anything, reqTopUp.Amount, modelUser.ID).Return(modelWallet, nil)
				tur.On("PostNewTransaction", mock.Anything, modelUser.ID, modelSof.ID, reqTopUp.Amount, modelSof.Name, "").Return(modelTx, nil)
				tur.On("PostTransactionUserPivot", mock.Anything, modelUser.ID, modelTx.ID).Return(nil)
				ur.On("PutAttemptGame", mock.Anything, 1, modelUser.ID).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrISE,
		},
		{
			name:      "should success all process top up when user top up fund",
			reqUserId: int64(1),
			reqBody: dtos.RequestTopUpWallet{
				Amount:       float64(10000000),
				SourceOfFund: int64(1),
			},
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser, nil)
				sofr.On("GetSourceOfFundById", mock.Anything, reqTopUp.SourceOfFund).Return(modelSof, nil)
				wr.On("GetWalletByIdUser", mock.Anything, modelUser.ID).Return(modelWallet, nil)
				wr.On("PutChangeBalanceWallet", mock.Anything, reqTopUp.Amount, modelUser.ID).Return(modelWallet, nil)
				tur.On("PostNewTransaction", mock.Anything, modelUser.ID, modelSof.ID, reqTopUp.Amount, modelSof.Name, "").Return(modelTx, nil)
				tur.On("PostTransactionUserPivot", mock.Anything, modelUser.ID, modelTx.ID).Return(nil)
				ur.On("PutAttemptGame", mock.Anything, 1, modelUser.ID).Return(modelUser, nil)
			},
			returnAtomic: userAndWallet,
			want:         resUserAndWallet,
			wantErr:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTx := &mocks.TransactionRepository{}
			mockTxUser := &mocks.TransactionUserRepository{}
			mockSof := &mocks.SourceOfFundRepository{}
			mockUser := &mocks.UserRepository{}
			mockWallet := &mocks.WalletRepository{}
			userService := services.NewWalletServiceImplementation(mockUser, mockWallet, mockTxUser, mockTx, mockSof)
			tt.mockRepository(mockTx, mockTxUser, mockSof, mockUser, mockWallet, tt.reqUserId, tt.returnAtomic, tt.wantErr)

			result, err := userService.PutTopWalletService(context.Background(), tt.reqBody, tt.reqUserId)
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestWalletServiceImplementation_PostTransferFundService(t *testing.T) {
	tests := []struct {
		name           string
		reqUserId      int64
		reqBody        dtos.RequestTransferFund
		mockRepository func(*mocks.TransactionRepository, *mocks.TransactionUserRepository, *mocks.SourceOfFundRepository, *mocks.UserRepository, *mocks.WalletRepository, int64, *dtos.ResponseSuccessTransfer, error)
		returnAtomic   *dtos.ResponseSuccessTransfer
		want           *dtos.ResponseSuccessTransfer
		wantErr        error
	}{
		{
			name:      "should return error when failed get user by id",
			reqUserId: int64(1),
			reqBody:   reqTransfer,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *dtos.ResponseSuccessTransfer,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrInvalidAccessToken,
		},
		{
			name:      "should return error when failed get wallet user by id",
			reqUserId: int64(1),
			reqBody:   reqTransfer,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *dtos.ResponseSuccessTransfer,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser, nil)
				wr.On("GetWalletByIdUser", mock.Anything, userId).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrInvalidAccessToken,
		},
		{
			name:      "should return error when user transfer to their account",
			reqUserId: int64(1),
			reqBody:   reqTransfer,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *dtos.ResponseSuccessTransfer,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser, nil)
				wr.On("GetWalletByIdUser", mock.Anything, userId).Return(modelWallet, nil)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrWalletTransferToTheirOwnWallet,
		},
		{
			name:      "should return error when user does not have enough fund",
			reqUserId: int64(1),
			reqBody: dtos.RequestTransferFund{
				ToWalletNumber: "7770000000002",
				Amount:         float64(10000000),
				Description:    "mantap",
			},
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *dtos.ResponseSuccessTransfer,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser, nil)
				wr.On("GetWalletByIdUser", mock.Anything, userId).Return(modelWallet, nil)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrWalletBalanceIsInsufficient,
		},
		{
			name:      "should return error when wallet number not valid",
			reqUserId: int64(1),
			reqBody: dtos.RequestTransferFund{
				ToWalletNumber: "7770000000002",
				Amount:         float64(1000),
				Description:    "mantap",
			},
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *dtos.ResponseSuccessTransfer,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser, nil)
				wr.On("GetWalletByIdUser", mock.Anything, userId).Return(modelWalletWithBigBalance2, nil)
				wr.On("GetWalletByWalletNumber", mock.Anything, "7770000000002").Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrWalletNumberIsNotExists,
		},
		{
			name:      "should return error when failed put change balance wallet user",
			reqUserId: int64(1),
			reqBody: dtos.RequestTransferFund{
				ToWalletNumber: "7770000000002",
				Amount:         float64(1000),
				Description:    "mantap",
			},
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *dtos.ResponseSuccessTransfer,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser, nil)
				wr.On("GetWalletByIdUser", mock.Anything, userId).Return(modelWalletWithBigBalance2, nil)
				wr.On("GetWalletByWalletNumber", mock.Anything, "7770000000002").Return(modelWalletWithBigBalance, nil)
				wr.On("PutChangeBalanceWallet", mock.Anything, (modelWalletWithBigBalance.Balance.InexactFloat64()+float64(1000)), modelWalletWithBigBalance.UserId).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrISE,
		},
		{
			name:      "should return error when failed post record transactions",
			reqUserId: int64(1),
			reqBody: dtos.RequestTransferFund{
				ToWalletNumber: "7770000000002",
				Amount:         float64(1000),
				Description:    "mantap",
			},
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *dtos.ResponseSuccessTransfer,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser, nil)
				wr.On("GetWalletByIdUser", mock.Anything, userId).Return(modelWalletWithBigBalance2, nil)
				wr.On("GetWalletByWalletNumber", mock.Anything, "7770000000002").Return(modelWalletWithBigBalance, nil)
				wr.On("PutChangeBalanceWallet", mock.Anything, (modelWalletWithBigBalance.Balance.InexactFloat64()+float64(1000)), modelWalletWithBigBalance.UserId).Return(modelWalletWithBigBalanceChangeBalance, nil)
				tur.On("PostNewTransaction", mock.Anything, modelWalletWithBigBalanceChangeBalance.ID, int64(1), float64(1000), "Bank Transfer", "mantap").Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrISE,
		},
		{
			name:      "should return error when Post Transaction User Pivot",
			reqUserId: int64(1),
			reqBody: dtos.RequestTransferFund{
				ToWalletNumber: "7770000000002",
				Amount:         float64(1000),
				Description:    "mantap",
			},
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *dtos.ResponseSuccessTransfer,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser, nil)
				wr.On("GetWalletByIdUser", mock.Anything, userId).Return(modelWalletWithBigBalance2, nil)
				wr.On("GetWalletByWalletNumber", mock.Anything, "7770000000002").Return(modelWalletWithBigBalance, nil)
				wr.On("PutChangeBalanceWallet", mock.Anything, (modelWalletWithBigBalance.Balance.InexactFloat64()+float64(1000)), modelWalletWithBigBalance.UserId).Return(modelWalletWithBigBalanceChangeBalance, nil)
				tur.On("PostNewTransaction", mock.Anything, modelWalletWithBigBalanceChangeBalance.ID, int64(1), float64(1000), "Bank Transfer", "mantap").Return(modelTx, nil)
				tur.On("PostTransactionUserPivot", mock.Anything, modelUser.ID, modelTx.ID).Return(err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrISE,
		},
		{
			name:      "should return success when user want to transfer fund to another user",
			reqUserId: int64(1),
			reqBody: dtos.RequestTransferFund{
				ToWalletNumber: "7770000000002",
				Amount:         float64(1000),
				Description:    "mantap",
			},
			mockRepository: func(
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				sofr *mocks.SourceOfFundRepository,
				ur *mocks.UserRepository,
				wr *mocks.WalletRepository,
				userId int64,
				returnAtomic *dtos.ResponseSuccessTransfer,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser, nil)
				wr.On("GetWalletByIdUser", mock.Anything, userId).Return(modelWalletWithBigBalance2, nil)
				wr.On("GetWalletByWalletNumber", mock.Anything, "7770000000002").Return(modelWalletWithBigBalance, nil)
				wr.On("PutChangeBalanceWallet", mock.Anything, (modelWalletWithBigBalance.Balance.InexactFloat64()+float64(1000)), modelWalletWithBigBalance.UserId).Return(modelWalletWithBigBalanceChangeBalance, nil)
				tur.On("PostNewTransaction", mock.Anything, modelWalletWithBigBalanceChangeBalance.ID, int64(1), float64(1000), "Bank Transfer", "mantap").Return(modelTx, nil)
				tur.On("PostTransactionUserPivot", mock.Anything, modelUser.ID, modelTx.ID).Return(nil)
			},
			returnAtomic: &dtos.ResponseSuccessTransfer{},
			want:         &dtos.ResponseSuccessTransfer{},
			wantErr:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTx := &mocks.TransactionRepository{}
			mockTxUser := &mocks.TransactionUserRepository{}
			mockSof := &mocks.SourceOfFundRepository{}
			mockUser := &mocks.UserRepository{}
			mockWallet := &mocks.WalletRepository{}
			userService := services.NewWalletServiceImplementation(mockUser, mockWallet, mockTxUser, mockTx, mockSof)
			tt.mockRepository(mockTx, mockTxUser, mockSof, mockUser, mockWallet, tt.reqUserId, tt.returnAtomic, tt.wantErr)

			result, err := userService.PostTransferFundService(context.Background(), tt.reqBody, tt.reqUserId)
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
