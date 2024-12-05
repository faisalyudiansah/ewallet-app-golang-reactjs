package services_test

import (
	"context"
	"testing"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/apperrors"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var modelGameBox = &models.GameBox{
	ID:        int64(1),
	UserID:    int64(1),
	IsOpen:    true,
	Box1:      float64(1),
	Box2:      float64(2),
	Box3:      float64(3),
	Box4:      float64(4),
	Box5:      float64(5),
	Box6:      float64(6),
	Box7:      float64(7),
	Box8:      float64(8),
	Box9:      float64(9),
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
	DeleteAt:  &time.Time{},
}

func TestGameServiceImplementation_PostStartGameService(t *testing.T) {
	tests := []struct {
		name           string
		reqUserId      int64
		mockRepository func(
			*mocks.GameRepositoryInterface,
			*mocks.TransactionRepository,
			*mocks.TransactionUserRepository,
			*mocks.UserRepository,
			int64,
			[]int,
			error,
		)
		returnAtomic []int
		want         []int
		wantErr      error
	}{
		{
			name:      "should success start game",
			reqUserId: int64(1),
			mockRepository: func(
				gri *mocks.GameRepositoryInterface,
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				ur *mocks.UserRepository,
				reqUserId int64,
				returnAtomic []int,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, reqUserId).Return(modelUser2, nil)
				gri.On("CreateGameBox", mock.Anything, reqUserId, mock.Anything).Return(nil)
				ur.On("PutAttemptGame", mock.Anything, 0, modelUser2.ID).Return(modelUser2, nil)
			},
			returnAtomic: []int{},
			want:         []int{},
			wantErr:      nil,
		},
		{
			name:      "should error Put AttemptGame",
			reqUserId: int64(1),
			mockRepository: func(
				gri *mocks.GameRepositoryInterface,
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				ur *mocks.UserRepository,
				reqUserId int64,
				returnAtomic []int,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, reqUserId).Return(modelUser2, nil)
				gri.On("CreateGameBox", mock.Anything, reqUserId, mock.Anything).Return(nil)
				ur.On("PutAttemptGame", mock.Anything, 0, modelUser2.ID).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrISE,
		},
		{
			name:      "should error Create GameBox",
			reqUserId: int64(1),
			mockRepository: func(
				gri *mocks.GameRepositoryInterface,
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				ur *mocks.UserRepository,
				reqUserId int64,
				returnAtomic []int,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, reqUserId).Return(modelUser2, nil)
				gri.On("CreateGameBox", mock.Anything, reqUserId, mock.Anything).Return(err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrISE,
		},
		{
			name:      "should error when user ttry to play but does not have attempt game",
			reqUserId: int64(1),
			mockRepository: func(
				gri *mocks.GameRepositoryInterface,
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				ur *mocks.UserRepository,
				reqUserId int64,
				returnAtomic []int,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, reqUserId).Return(modelUser, nil)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrGameZeroChance,
		},
		{
			name:      "should error when Get User By Id",
			reqUserId: int64(1),
			mockRepository: func(
				gri *mocks.GameRepositoryInterface,
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				ur *mocks.UserRepository,
				reqUserId int64,
				returnAtomic []int,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, reqUserId).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrInvalidAccessToken,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGame := &mocks.GameRepositoryInterface{}
			mockTx := &mocks.TransactionRepository{}
			mockTxU := &mocks.TransactionUserRepository{}
			mockUser := &mocks.UserRepository{}
			gameService := services.NewGameServiceImplementation(mockGame, mockTxU, mockTx, mockUser)
			tt.mockRepository(
				mockGame,
				mockTx,
				mockTxU,
				mockUser,
				tt.reqUserId,
				tt.returnAtomic,
				tt.wantErr,
			)

			result, err := gameService.PostStartGameService(context.Background(), tt.reqUserId)
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGameServiceImplementation_PostChooseGachaBoxService(t *testing.T) {
	tests := []struct {
		name             string
		notValidBoxIndex bool
		reqUserId        int64
		mockRepository   func(
			*mocks.GameRepositoryInterface,
			*mocks.TransactionRepository,
			*mocks.TransactionUserRepository,
			*mocks.UserRepository,
			int64,
			float64,
			error,
		)
		returnAtomic float64
		want         float64
		wantErr      error
	}{
		{
			name:             "should success get reward from gacha",
			notValidBoxIndex: true,
			reqUserId:        int64(1),
			mockRepository: func(
				gri *mocks.GameRepositoryInterface,
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				ur *mocks.UserRepository,
				userId int64,
				returnAtomic float64,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser2, nil)
				gri.On("GetGameBox", mock.Anything, userId).Return(modelGameBox, nil)
				gri.On("ChangeStatusBox", mock.Anything, modelGameBox.ID).Return(nil)
				tur.On("PostNewTransaction", mock.Anything, modelUser2.ID, int64(4), float64(1), "Reward", "").Return(modelTx, nil)
				tur.On("PostTransactionUserPivot", mock.Anything, modelUser2.ID, modelTx.ID).Return(nil)
			},
			returnAtomic: float64(1),
			want:         float64(1),
			wantErr:      nil,
		},
		{
			name:             "should error when Post Transaction User on table Pivot",
			notValidBoxIndex: true,
			reqUserId:        int64(1),
			mockRepository: func(
				gri *mocks.GameRepositoryInterface,
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				ur *mocks.UserRepository,
				userId int64,
				returnAtomic float64,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser2, nil)
				gri.On("GetGameBox", mock.Anything, userId).Return(modelGameBox, nil)
				gri.On("ChangeStatusBox", mock.Anything, modelGameBox.ID).Return(nil)
				tur.On("PostNewTransaction", mock.Anything, modelUser2.ID, int64(4), float64(1), "Reward", "").Return(modelTx, nil)
				tur.On("PostTransactionUserPivot", mock.Anything, modelUser2.ID, modelTx.ID).Return(err)
			},
			returnAtomic: float64(0),
			want:         float64(0),
			wantErr:      apperrors.ErrISE,
		},
		{
			name:             "should error when Post New Transaction",
			notValidBoxIndex: true,
			reqUserId:        int64(1),
			mockRepository: func(
				gri *mocks.GameRepositoryInterface,
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				ur *mocks.UserRepository,
				userId int64,
				returnAtomic float64,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser2, nil)
				gri.On("GetGameBox", mock.Anything, userId).Return(modelGameBox, nil)
				gri.On("ChangeStatusBox", mock.Anything, modelGameBox.ID).Return(nil)
				tur.On("PostNewTransaction", mock.Anything, modelUser2.ID, int64(4), float64(1), "Reward", "").Return(nil, err)
			},
			returnAtomic: float64(0),
			want:         float64(0),
			wantErr:      apperrors.ErrISE,
		},
		{
			name:             "should error when Change Status Box",
			notValidBoxIndex: true,
			reqUserId:        int64(1),
			mockRepository: func(
				gri *mocks.GameRepositoryInterface,
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				ur *mocks.UserRepository,
				userId int64,
				returnAtomic float64,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser2, nil)
				gri.On("GetGameBox", mock.Anything, userId).Return(modelGameBox, nil)
				gri.On("ChangeStatusBox", mock.Anything, modelGameBox.ID).Return(err)
			},
			returnAtomic: float64(0),
			want:         float64(0),
			wantErr:      apperrors.ErrISE,
		},
		{
			name:             "should error when choose invalid box index",
			notValidBoxIndex: false,
			reqUserId:        int64(1),
			mockRepository: func(
				gri *mocks.GameRepositoryInterface,
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				ur *mocks.UserRepository,
				userId int64,
				returnAtomic float64,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser2, nil)
				gri.On("GetGameBox", mock.Anything, userId).Return(modelGameBox, nil)
			},
			returnAtomic: float64(0),
			want:         float64(0),
			wantErr:      apperrors.ErrGameInvalidBoxIndex,
		},
		{
			name:             "should error when Get Game Box",
			notValidBoxIndex: true,
			reqUserId:        int64(1),
			mockRepository: func(
				gri *mocks.GameRepositoryInterface,
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				ur *mocks.UserRepository,
				userId int64,
				returnAtomic float64,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, err)
				ur.On("GetUserById", mock.Anything, userId).Return(modelUser2, nil)
				gri.On("GetGameBox", mock.Anything, userId).Return(nil, err)
			},
			returnAtomic: float64(0),
			want:         float64(0),
			wantErr:      apperrors.ErrGameZeroChance,
		},
		{
			name:             "should error when Get User By Id",
			notValidBoxIndex: true,
			reqUserId:        int64(1),
			mockRepository: func(
				gri *mocks.GameRepositoryInterface,
				tr *mocks.TransactionRepository,
				tur *mocks.TransactionUserRepository,
				ur *mocks.UserRepository,
				userId int64,
				returnAtomic float64,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, err)
				ur.On("GetUserById", mock.Anything, userId).Return(nil, err)
			},
			returnAtomic: float64(0),
			want:         float64(0),
			wantErr:      apperrors.ErrInvalidAccessToken,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGame := &mocks.GameRepositoryInterface{}
			mockTx := &mocks.TransactionRepository{}
			mockTxU := &mocks.TransactionUserRepository{}
			mockUser := &mocks.UserRepository{}
			gameService := services.NewGameServiceImplementation(mockGame, mockTxU, mockTx, mockUser)
			tt.mockRepository(
				mockGame,
				mockTx,
				mockTxU,
				mockUser,
				tt.reqUserId,
				tt.returnAtomic,
				tt.wantErr,
			)

			if !tt.notValidBoxIndex {
				result, err := gameService.PostChooseGachaBoxService(context.Background(), tt.reqUserId, 100)
				assert.Equal(t, tt.want, result)
				assert.Equal(t, tt.wantErr, err)
			} else {
				result, err := gameService.PostChooseGachaBoxService(context.Background(), tt.reqUserId, 1)
				assert.Equal(t, tt.want, result)
				assert.Equal(t, tt.wantErr, err)
			}
		})
	}
}
