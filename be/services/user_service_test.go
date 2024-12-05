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

var (
	userId          = int64(1)
	walletId        = int64(1)
	emailUser       = "frieren@gmail.com"
	pwdUser         = "12345"
	fullnameUser    = "Mbak Frieren"
	bdtUser         = "1908-10-10"
	walletNumber    = "7770000000001"
	walletNumber2   = "7770000000002"
	walletNumber3   = "7770000000003"
	emptyBalance, _ = decimal.NewFromString("0")
	bigBalance, _   = decimal.NewFromString("10000000")
	newBalance, _   = decimal.NewFromString("10001000")
)

var reqBodyRegister = dtos.RequestRegisterUser{
	Email:     emailUser,
	Password:  pwdUser,
	FullName:  fullnameUser,
	BirthDate: bdtUser,
}

var reqBodyLogin = dtos.RequestLoginUser{
	Email:    emailUser,
	Password: pwdUser,
}

var modelUser = &models.User{
	ID:         userId,
	Email:      emailUser,
	Password:   pwdUser,
	ChanceGame: 0,
	FullName:   fullnameUser,
	BirthDate:  time.Time{},
	CreatedAt:  time.Time{},
	UpdatedAt:  time.Time{},
	DeleteAt:   &time.Time{},
}

var modelUser2 = &models.User{
	ID:         int64(100),
	Email:      emailUser,
	Password:   pwdUser,
	ChanceGame: 1,
	FullName:   fullnameUser,
	BirthDate:  time.Time{},
	CreatedAt:  time.Time{},
	UpdatedAt:  time.Time{},
	DeleteAt:   &time.Time{},
}

var modelWallet = &models.Wallet{
	ID:           walletId,
	UserId:       userId,
	WalletNumber: walletNumber,
	Balance:      emptyBalance,
	CreatedAt:    time.Time{},
	UpdatedAt:    time.Time{},
	DeleteAt:     &time.Time{},
}

var modelWalletWithBigBalance = &models.Wallet{
	ID:           walletId,
	UserId:       userId,
	WalletNumber: walletNumber2,
	Balance:      bigBalance,
	CreatedAt:    time.Time{},
	UpdatedAt:    time.Time{},
	DeleteAt:     &time.Time{},
}

var modelWalletWithBigBalanceChangeBalance = &models.Wallet{
	ID:           walletId,
	UserId:       userId,
	WalletNumber: walletNumber2,
	Balance:      newBalance,
	CreatedAt:    time.Time{},
	UpdatedAt:    time.Time{},
	DeleteAt:     &time.Time{},
}

var modelWalletWithBigBalance2 = &models.Wallet{
	ID:           walletId,
	UserId:       userId,
	WalletNumber: walletNumber3,
	Balance:      bigBalance,
	CreatedAt:    time.Time{},
	UpdatedAt:    time.Time{},
	DeleteAt:     &time.Time{},
}

var resWallet = dtos.ResponseWallet{
	ID:           walletId,
	UserId:       userId,
	WalletNumber: walletNumber,
	Balance:      emptyBalance,
	CreatedAt:    time.Time{},
	UpdatedAt:    time.Time{},
	DeleteAt:     &time.Time{},
}

var resUserAndWallet = &dtos.ResponseUserAndWallet{
	ID:         userId,
	Email:      emailUser,
	ChanceGame: 0,
	FullName:   fullnameUser,
	BirthDate:  time.Time{},
	CreatedAt:  time.Time{},
	UpdatedAt:  time.Time{},
	DeleteAt:   &time.Time{},
	Wallet:     resWallet,
}

var resUserDto = &dtos.ResponseUser{
	ID:         userId,
	Email:      emailUser,
	ChanceGame: 0,
	FullName:   fullnameUser,
	BirthDate:  time.Time{},
	CreatedAt:  time.Time{},
	UpdatedAt:  time.Time{},
	DeleteAt:   &time.Time{},
}

var userAndWallet = &models.UserAndWallet{
	User:   *modelUser,
	Wallet: *modelWallet,
}

func TestUserServiceImplementation_PostRegisterUserService(t *testing.T) {
	tests := []struct {
		name           string
		reqBody        dtos.RequestRegisterUser
		mockRepository func(*mocks.TransactionRepository, *mocks.UserRepository, *mocks.Bcrypt, *mocks.WalletRepository, dtos.RequestRegisterUser, *models.UserAndWallet, error)
		returnAtomic   *models.UserAndWallet
		want           *dtos.ResponseUserAndWallet
		wantErr        error
	}{
		{
			name:    "should success create user when user register new account",
			reqBody: reqBodyRegister,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				b *mocks.Bcrypt,
				wr *mocks.WalletRepository,
				rru dtos.RequestRegisterUser,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("IsEmailAlreadyRegistered", mock.Anything, reqBodyRegister.Email).Return(false)
				b.On("HashPassword", reqBodyRegister.Password, 10).Return([]byte(pwdUser), nil)
				ur.On("PostUser", mock.Anything, reqBodyRegister, pwdUser).Return(modelUser, nil)
				wr.On("PostCreateWalletUser", mock.Anything, modelUser.ID).Return(modelWallet, nil)
			},
			returnAtomic: userAndWallet,
			want:         resUserAndWallet,
			wantErr:      nil,
		},
		{
			name:    "should error something wrong in system when system try to post/create new wallet",
			reqBody: reqBodyRegister,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				b *mocks.Bcrypt,
				wr *mocks.WalletRepository,
				rru dtos.RequestRegisterUser,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("IsEmailAlreadyRegistered", mock.Anything, reqBodyRegister.Email).Return(false)
				b.On("HashPassword", reqBodyRegister.Password, 10).Return([]byte(pwdUser), nil)
				ur.On("PostUser", mock.Anything, reqBodyRegister, pwdUser).Return(modelUser, nil)
				wr.On("PostCreateWalletUser", mock.Anything, modelUser.ID).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrUserFailedRegister,
		},
		{
			name:    "should error something wrong in system when system try to post/create new account",
			reqBody: reqBodyRegister,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				b *mocks.Bcrypt,
				wr *mocks.WalletRepository,
				rru dtos.RequestRegisterUser,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("IsEmailAlreadyRegistered", mock.Anything, reqBodyRegister.Email).Return(false)
				b.On("HashPassword", reqBodyRegister.Password, 10).Return([]byte(pwdUser), nil)
				ur.On("PostUser", mock.Anything, reqBodyRegister, pwdUser).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrUserFailedRegister,
		},
		{
			name:    "should error something wrong in bcrypt when bcrypt try to hash password",
			reqBody: reqBodyRegister,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				b *mocks.Bcrypt,
				wr *mocks.WalletRepository,
				rru dtos.RequestRegisterUser,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("IsEmailAlreadyRegistered", mock.Anything, reqBodyRegister.Email).Return(false)
				b.On("HashPassword", reqBodyRegister.Password, 10).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrUserFailedRegister,
		},
		{
			name:    "should error when user register input an email already exists before",
			reqBody: reqBodyRegister,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				b *mocks.Bcrypt,
				wr *mocks.WalletRepository,
				rru dtos.RequestRegisterUser,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("IsEmailAlreadyRegistered", mock.Anything, reqBodyRegister.Email).Return(true)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrUserEmailAlreadyExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTx := &mocks.TransactionRepository{}
			mockUser := &mocks.UserRepository{}
			mockBc := &mocks.Bcrypt{}
			mockWallet := &mocks.WalletRepository{}
			mockJwt := &mocks.JWTProvider{}
			mockGenerateNumber := &mocks.GenerateNumberInterface{}
			userService := services.NewUserServiceImplementation(mockUser, mockWallet, mockTx, mockBc, mockJwt, mockGenerateNumber)
			tt.mockRepository(mockTx, mockUser, mockBc, mockWallet, tt.reqBody, tt.returnAtomic, tt.wantErr)

			result, err := userService.PostRegisterUserService(context.Background(), tt.reqBody)
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestUserServiceImplementation_PostLoginUserService(t *testing.T) {
	tests := []struct {
		name           string
		reqBody        dtos.RequestLoginUser
		mockRepository func(*mocks.TransactionRepository, *mocks.UserRepository, *mocks.Bcrypt, *mocks.WalletRepository, *mocks.JWTProvider, dtos.RequestLoginUser, string, error)
		returnAtomic   string
		want           *dtos.ResponseAccessToken
		wantErr        error
	}{
		{
			name:    "should success login when user try to login and without error",
			reqBody: reqBodyLogin,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				b *mocks.Bcrypt,
				wr *mocks.WalletRepository,
				jwt *mocks.JWTProvider,
				rlu dtos.RequestLoginUser,
				returnAtomic string,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserByEmail", mock.Anything, rlu.Email).Return(modelUser, nil)
				b.On("CheckPassword", rlu.Password, []byte(modelUser.Password)).Return(true, nil)
				jwt.On("CreateToken", int64(modelUser.ID)).Return(returnAtomic, nil)
			},
			returnAtomic: "ini_jwt_token_dong",
			want: &dtos.ResponseAccessToken{
				AccessToken: "ini_jwt_token_dong",
			},
			wantErr: nil,
		},
		{
			name:    "should error when something wrong in create JWT",
			reqBody: reqBodyLogin,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				b *mocks.Bcrypt,
				wr *mocks.WalletRepository,
				jwt *mocks.JWTProvider,
				rlu dtos.RequestLoginUser,
				returnAtomic string,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, err)
				ur.On("GetUserByEmail", mock.Anything, rlu.Email).Return(modelUser, nil)
				b.On("CheckPassword", rlu.Password, []byte(modelUser.Password)).Return(true, nil)
				jwt.On("CreateToken", int64(modelUser.ID)).Return("", err)
			},
			returnAtomic: "",
			want:         nil,
			wantErr:      apperrors.ErrISE,
		},
		{
			name:    "should error when invalid password input user",
			reqBody: reqBodyLogin,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				b *mocks.Bcrypt,
				wr *mocks.WalletRepository,
				jwt *mocks.JWTProvider,
				rlu dtos.RequestLoginUser,
				returnAtomic string,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, err)
				ur.On("GetUserByEmail", mock.Anything, rlu.Email).Return(modelUser, nil)
				b.On("CheckPassword", rlu.Password, []byte(modelUser.Password)).Return(false, err)
			},
			returnAtomic: "",
			want:         nil,
			wantErr:      apperrors.ErrUserInvalidEmailPassword,
		},
		{
			name:    "should error when user input invalid email",
			reqBody: reqBodyLogin,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				b *mocks.Bcrypt,
				wr *mocks.WalletRepository,
				jwt *mocks.JWTProvider,
				rlu dtos.RequestLoginUser,
				returnAtomic string,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, err)
				ur.On("GetUserByEmail", mock.Anything, rlu.Email).Return(nil, err)
			},
			returnAtomic: "",
			want:         nil,
			wantErr:      apperrors.ErrUserInvalidEmailPassword,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTx := &mocks.TransactionRepository{}
			mockUser := &mocks.UserRepository{}
			mockBc := &mocks.Bcrypt{}
			mockWallet := &mocks.WalletRepository{}
			mockJwt := &mocks.JWTProvider{}
			mockGenerateNumber := &mocks.GenerateNumberInterface{}
			userService := services.NewUserServiceImplementation(mockUser, mockWallet, mockTx, mockBc, mockJwt, mockGenerateNumber)
			tt.mockRepository(mockTx, mockUser, mockBc, mockWallet, mockJwt, tt.reqBody, tt.returnAtomic, tt.wantErr)

			result, err := userService.PostLoginUserService(context.Background(), tt.reqBody)
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestUserServiceImplementation_GetUserDetailService(t *testing.T) {
	tests := []struct {
		name           string
		req            int64
		mockRepository func(*mocks.TransactionRepository, *mocks.UserRepository, *mocks.Bcrypt, *mocks.WalletRepository, int64, *models.UserAndWallet, error)
		returnAtomic   *models.UserAndWallet
		want           *dtos.ResponseUserAndWallet
		wantErr        error
	}{
		{
			name: "should success login when user try to login and without error",
			req:  int64(1),
			mockRepository: func(
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				b *mocks.Bcrypt,
				wr *mocks.WalletRepository,
				reqId int64,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, reqId).Return(modelUser, err)
				wr.On("GetWalletByIdUser", mock.Anything, modelUser.ID).Return(modelWallet, err)
			},
			returnAtomic: userAndWallet,
			want:         resUserAndWallet,
			wantErr:      nil,
		},
		{
			name: "should error when failed get wallet user",
			req:  int64(1),
			mockRepository: func(
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				b *mocks.Bcrypt,
				wr *mocks.WalletRepository,
				reqId int64,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, reqId).Return(modelUser, nil)
				wr.On("GetWalletByIdUser", mock.Anything, modelUser.ID).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrInvalidAccessToken,
		},
		{
			name: "should error when failed get data user by id",
			req:  int64(1),
			mockRepository: func(
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				b *mocks.Bcrypt,
				wr *mocks.WalletRepository,
				reqId int64,
				returnAtomic *models.UserAndWallet,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserById", mock.Anything, reqId).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrInvalidAccessToken,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTx := &mocks.TransactionRepository{}
			mockUser := &mocks.UserRepository{}
			mockBc := &mocks.Bcrypt{}
			mockWallet := &mocks.WalletRepository{}
			mockJwt := &mocks.JWTProvider{}
			mockGenerateNumber := &mocks.GenerateNumberInterface{}
			userService := services.NewUserServiceImplementation(mockUser, mockWallet, mockTx, mockBc, mockJwt, mockGenerateNumber)
			tt.mockRepository(mockTx, mockUser, mockBc, mockWallet, tt.req, tt.returnAtomic, tt.wantErr)

			result, err := userService.GetUserDetailService(context.Background(), tt.req)
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
