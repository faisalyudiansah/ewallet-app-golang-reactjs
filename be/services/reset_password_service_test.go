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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var requestForgetPassword = dtos.RequestForgetPassword{
	Email: "frieren@gmail.com",
}

var resForgetPassword = &dtos.ResponseTokenResetPassword{
	TokenResetPassword: "ini_token",
	LinkResetPassword:  "/reset-password/ini_token",
}

var modelResetPassword = &models.ResetPassword{
	ID:        int64(1),
	UserId:    int64(1),
	Token:     "ini_token",
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
	DeleteAt:  &time.Time{},
}

var reqResetPwd = dtos.RequestResetPassword{
	Email:       "frieren@gmail.com",
	NewPassword: "12345",
}

func TestResetPasswordServiceImplementation_PostForgetPasswordService(t *testing.T) {
	tests := []struct {
		name           string
		reqBody        dtos.RequestForgetPassword
		mockRepository func(*mocks.TransactionRepository, *mocks.UserRepository, *mocks.ResetPasswordRepository, *mocks.GenerateNumberInterface, dtos.RequestForgetPassword, string, error)
		returnAtomic   string
		want           *dtos.ResponseTokenResetPassword
		wantErr        error
	}{
		{
			name:    "should success get token for reset password",
			reqBody: requestForgetPassword,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				rpr *mocks.ResetPasswordRepository,
				gni *mocks.GenerateNumberInterface,
				body dtos.RequestForgetPassword,
				returnAtomic string,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				ur.On("GetUserByEmail", mock.Anything, body.Email).Return(modelUser, nil)
				gni.On("GenerateTokenResetPassword", modelUser.ID).Return("ini_token")
				rpr.On("PostNewDataResetPassword", mock.Anything, modelUser.ID, "ini_token").Return(modelResetPassword, nil)
			},
			returnAtomic: "ini_token",
			want:         resForgetPassword,
			wantErr:      nil,
		},
		{
			name:    "should error post new data forget password",
			reqBody: requestForgetPassword,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				rpr *mocks.ResetPasswordRepository,
				gni *mocks.GenerateNumberInterface,
				body dtos.RequestForgetPassword,
				returnAtomic string,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, err)
				ur.On("GetUserByEmail", mock.Anything, body.Email).Return(modelUser, nil)
				gni.On("GenerateTokenResetPassword", modelUser.ID).Return("ini_token")
				rpr.On("PostNewDataResetPassword", mock.Anything, modelUser.ID, "ini_token").Return(nil, err)
			},
			returnAtomic: "ini_token",
			want:         nil,
			wantErr:      apperrors.ErrISE,
		},
		{
			name:    "should error get user by email",
			reqBody: requestForgetPassword,
			mockRepository: func(
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				rpr *mocks.ResetPasswordRepository,
				gni *mocks.GenerateNumberInterface,
				body dtos.RequestForgetPassword,
				returnAtomic string,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, err)
				ur.On("GetUserByEmail", mock.Anything, body.Email).Return(nil, err)
			},
			returnAtomic: "",
			want:         nil,
			wantErr:      apperrors.ErrUserEmailNotExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockReset := &mocks.ResetPasswordRepository{}
			mockTx := &mocks.TransactionRepository{}
			mockUser := &mocks.UserRepository{}
			mockBc := &mocks.Bcrypt{}
			mockGenerateNumber := &mocks.GenerateNumberInterface{}
			resetPasswordService := services.NewResetPasswordServiceImplementation(mockReset, mockUser, mockTx, mockGenerateNumber, mockBc)
			tt.mockRepository(mockTx, mockUser, mockReset, mockGenerateNumber, tt.reqBody, tt.returnAtomic, tt.wantErr)

			result, err := resetPasswordService.PostForgetPasswordService(context.Background(), tt.reqBody)
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestResetPasswordServiceImplementation_PutResetPasswordService(t *testing.T) {
	tests := []struct {
		name           string
		token          string
		userId         int64
		reqBody        dtos.RequestResetPassword
		mockRepository func(*mocks.Bcrypt, *mocks.TransactionRepository, *mocks.UserRepository, *mocks.ResetPasswordRepository, *mocks.GenerateNumberInterface, string, dtos.RequestResetPassword, *models.User, error)
		returnAtomic   *models.User
		want           *dtos.ResponseUser
		wantErr        error
	}{
		{
			name:    "should success reset password",
			token:   "ini_token",
			userId:  int64(1),
			reqBody: reqResetPwd,
			mockRepository: func(
				bc *mocks.Bcrypt,
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				rpr *mocks.ResetPasswordRepository,
				gni *mocks.GenerateNumberInterface,
				token string,
				body dtos.RequestResetPassword,
				returnAtomic *models.User,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				rpr.On("IsTokenResetValid", mock.Anything, token).Return(true)
				ur.On("GetUserByEmail", mock.Anything, body.Email).Return(modelUser, nil)
				bc.On("HashPassword", body.NewPassword, 10).Return([]byte(pwdUser), nil)
				ur.On("PutResetPassword", mock.Anything, pwdUser, userId).Return(modelUser, nil)
				rpr.On("DeleteToken", mock.Anything, token).Return(nil)
			},
			returnAtomic: modelUser,
			want:         resUserDto,
			wantErr:      nil,
		},
		{
			name:    "should failed delete token",
			token:   "ini_token",
			userId:  int64(1),
			reqBody: reqResetPwd,
			mockRepository: func(
				bc *mocks.Bcrypt,
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				rpr *mocks.ResetPasswordRepository,
				gni *mocks.GenerateNumberInterface,
				token string,
				body dtos.RequestResetPassword,
				returnAtomic *models.User,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				rpr.On("IsTokenResetValid", mock.Anything, token).Return(true)
				ur.On("GetUserByEmail", mock.Anything, body.Email).Return(modelUser, nil)
				bc.On("HashPassword", body.NewPassword, 10).Return([]byte(pwdUser), nil)
				ur.On("PutResetPassword", mock.Anything, pwdUser, userId).Return(modelUser, nil)
				rpr.On("DeleteToken", mock.Anything, token).Return(err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrISE,
		},
		{
			name:    "should failed put reset password",
			token:   "ini_token",
			userId:  int64(1),
			reqBody: reqResetPwd,
			mockRepository: func(
				bc *mocks.Bcrypt,
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				rpr *mocks.ResetPasswordRepository,
				gni *mocks.GenerateNumberInterface,
				token string,
				body dtos.RequestResetPassword,
				returnAtomic *models.User,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				rpr.On("IsTokenResetValid", mock.Anything, token).Return(true)
				ur.On("GetUserByEmail", mock.Anything, body.Email).Return(modelUser, nil)
				bc.On("HashPassword", body.NewPassword, 10).Return([]byte(pwdUser), nil)
				ur.On("PutResetPassword", mock.Anything, pwdUser, userId).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrISE,
		},
		{
			name:    "should failed hash password",
			token:   "ini_token",
			userId:  int64(1),
			reqBody: reqResetPwd,
			mockRepository: func(
				bc *mocks.Bcrypt,
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				rpr *mocks.ResetPasswordRepository,
				gni *mocks.GenerateNumberInterface,
				token string,
				body dtos.RequestResetPassword,
				returnAtomic *models.User,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				rpr.On("IsTokenResetValid", mock.Anything, token).Return(true)
				ur.On("GetUserByEmail", mock.Anything, body.Email).Return(modelUser, nil)
				bc.On("HashPassword", body.NewPassword, 10).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrUserFailedRegister,
		},
		{
			name:    "should failed hash password",
			token:   "ini_token",
			userId:  int64(1),
			reqBody: reqResetPwd,
			mockRepository: func(
				bc *mocks.Bcrypt,
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				rpr *mocks.ResetPasswordRepository,
				gni *mocks.GenerateNumberInterface,
				token string,
				body dtos.RequestResetPassword,
				returnAtomic *models.User,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				rpr.On("IsTokenResetValid", mock.Anything, token).Return(true)
				ur.On("GetUserByEmail", mock.Anything, body.Email).Return(modelUser2, nil)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrUnauthorization,
		},
		{
			name:    "should failed get user by email",
			token:   "ini_token",
			userId:  int64(1),
			reqBody: reqResetPwd,
			mockRepository: func(
				bc *mocks.Bcrypt,
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				rpr *mocks.ResetPasswordRepository,
				gni *mocks.GenerateNumberInterface,
				token string,
				body dtos.RequestResetPassword,
				returnAtomic *models.User,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				rpr.On("IsTokenResetValid", mock.Anything, token).Return(true)
				ur.On("GetUserByEmail", mock.Anything, body.Email).Return(nil, err)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrUserEmailNotExists,
		},
		{
			name:    "should failed when token not valid",
			token:   "ini_token",
			userId:  int64(1),
			reqBody: reqResetPwd,
			mockRepository: func(
				bc *mocks.Bcrypt,
				tr *mocks.TransactionRepository,
				ur *mocks.UserRepository,
				rpr *mocks.ResetPasswordRepository,
				gni *mocks.GenerateNumberInterface,
				token string,
				body dtos.RequestResetPassword,
				returnAtomic *models.User,
				err error,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(returnAtomic, err)
				rpr.On("IsTokenResetValid", mock.Anything, token).Return(false)
			},
			returnAtomic: nil,
			want:         nil,
			wantErr:      apperrors.ErrUserTokenResetPasswordNotValid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockReset := &mocks.ResetPasswordRepository{}
			mockTx := &mocks.TransactionRepository{}
			mockUser := &mocks.UserRepository{}
			mockBc := &mocks.Bcrypt{}
			mockGenerateNumber := &mocks.GenerateNumberInterface{}
			resetPasswordService := services.NewResetPasswordServiceImplementation(mockReset, mockUser, mockTx, mockGenerateNumber, mockBc)
			tt.mockRepository(mockBc, mockTx, mockUser, mockReset, mockGenerateNumber, tt.token, tt.reqBody, tt.returnAtomic, tt.wantErr)

			result, err := resetPasswordService.PutResetPasswordService(context.Background(), tt.token, tt.reqBody, tt.userId)
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
