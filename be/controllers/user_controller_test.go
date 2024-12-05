package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/apperrors"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/constants"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/controllers"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/dtos"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/helpers"
	helpercontext "git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/helpers/helper_context"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/middlewares"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/servers"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func SetUpRouterUser(userController *controllers.UserController) *gin.Engine {
	h := &servers.HandlerOps{
		UserController: userController,
	}
	return servers.SetupRoute(h)
}

var (
	userId          = int64(1)
	walletId        = int64(1)
	emailUser       = "frieren@gmail.com"
	pwdUser         = "12345"
	fullnameUser    = "Mbak Frieren"
	bdtUser         = "1908-10-10"
	walletNumber    = "7770000000001"
	emptyBalance, _ = decimal.NewFromString("0")
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

var resAccessToken = &dtos.ResponseAccessToken{
	AccessToken: "Frieren guguk, kemari guguk ayo nari nari",
}

func TestUserController_PostRegisterUserController(t *testing.T) {
	tests := []struct {
		name        string
		statusCode  int
		msg         string
		reqBody     interface{}
		resService  *dtos.ResponseUserAndWallet
		mockService func(*mocks.UserServiceInterface, *mocks.ValidationReqBodyInterface, dtos.RequestRegisterUser, error)
		wantErr     error
	}{
		{
			name:       "should success register without an error",
			statusCode: http.StatusCreated,
			msg:        constants.UserSuccessRegister,
			reqBody:    reqBodyRegister,
			resService: resUserAndWallet,
			mockService: func(usi *mocks.UserServiceInterface, vr *mocks.ValidationReqBodyInterface, reqBody dtos.RequestRegisterUser, err error) {
				var reqBodyReg dtos.RequestRegisterUser
				vr.On("ValidationReqBody", mock.AnythingOfType("*gin.Context"), &reqBodyReg).Return(nil)
				usi.On("PostRegisterUserService", mock.AnythingOfType("*gin.Context"), reqBodyReg).Return(resUserAndWallet, nil)
			},
			wantErr: nil,
		},
		{
			name:       "should error on validate request body when user register",
			statusCode: http.StatusBadRequest,
			msg:        constants.RequestBodyInvalid,
			reqBody:    reqBodyRegister,
			resService: nil,
			mockService: func(usi *mocks.UserServiceInterface, vr *mocks.ValidationReqBodyInterface, reqBody dtos.RequestRegisterUser, err error) {
				var reqBodyReg dtos.RequestRegisterUser
				vr.On("ValidationReqBody", mock.AnythingOfType("*gin.Context"), &reqBodyReg).Return(err)
			},
			wantErr: apperrors.ErrRequestBodyInvalid,
		},
		{
			name:       "should failed register a when user send the same email with another user",
			statusCode: http.StatusBadRequest,
			msg:        constants.UserEmailAlreadyExists,
			reqBody:    reqBodyRegister,
			resService: nil,
			mockService: func(usi *mocks.UserServiceInterface, vr *mocks.ValidationReqBodyInterface, reqBody dtos.RequestRegisterUser, err error) {
				var reqBodyReg dtos.RequestRegisterUser
				vr.On("ValidationReqBody", mock.AnythingOfType("*gin.Context"), &reqBodyReg).Return(nil)
				usi.On("PostRegisterUserService", mock.AnythingOfType("*gin.Context"), reqBodyReg).Return(nil, err)
			},
			wantErr: apperrors.ErrUserEmailAlreadyExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService := &mocks.UserServiceInterface{}
			mockValidateReqBody := &mocks.ValidationReqBodyInterface{}
			tt.mockService(mockUserService, mockValidateReqBody, tt.reqBody.(dtos.RequestRegisterUser), tt.wantErr)
			userController := controllers.NewUserController(mockUserService, mockValidateReqBody)

			g := SetUpRouterUser(userController)
			reqBody, _ := json.Marshal(tt.reqBody.(dtos.RequestRegisterUser))
			req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBuffer(reqBody))
			rec := httptest.NewRecorder()
			g.ServeHTTP(rec, req)

			var response map[string]interface{}
			json.Unmarshal(rec.Body.Bytes(), &response)

			assert.Equal(t, tt.statusCode, rec.Code)
			if tt.wantErr != nil {
				expected := dtos.ResponseMessageOnly{
					Message: tt.msg,
				}
				expectedToJson, _ := json.Marshal(expected)
				assert.Equal(t, string(expectedToJson), rec.Body.String())
			} else {
				expected := helpers.FormatterMessageWithOneUserAndWallet(tt.resService, tt.msg)
				expectedToJson, _ := json.Marshal(expected)
				assert.Equal(t, string(expectedToJson), rec.Body.String())
				assert.Equal(t, expected.Message, response["message"])
			}
		})
	}
}

func TestUserController_PostLoginUserController(t *testing.T) {
	tests := []struct {
		name        string
		statusCode  int
		msg         string
		reqBody     interface{}
		resService  *dtos.ResponseAccessToken
		mockService func(*mocks.UserServiceInterface, *mocks.ValidationReqBodyInterface, dtos.RequestLoginUser, error)
		wantErr     error
	}{
		{
			name:       "should success login without an error",
			statusCode: http.StatusOK,
			msg:        constants.UserSuccessLogin,
			reqBody:    reqBodyLogin,
			resService: resAccessToken,
			mockService: func(usi *mocks.UserServiceInterface, vr *mocks.ValidationReqBodyInterface, reqBody dtos.RequestLoginUser, err error) {
				var reqBodyReg dtos.RequestLoginUser
				vr.On("ValidationReqBody", mock.AnythingOfType("*gin.Context"), &reqBodyReg).Return(nil)
				usi.On("PostLoginUserService", mock.AnythingOfType("*gin.Context"), reqBodyReg).Return(resAccessToken, nil)
			},
			wantErr: nil,
		},
		{
			name:       "should error on validate request body when user login",
			statusCode: http.StatusBadRequest,
			msg:        constants.RequestBodyInvalid,
			reqBody:    reqBodyLogin,
			resService: resAccessToken,
			mockService: func(usi *mocks.UserServiceInterface, vr *mocks.ValidationReqBodyInterface, reqBody dtos.RequestLoginUser, err error) {
				var reqBodyReg dtos.RequestLoginUser
				vr.On("ValidationReqBody", mock.AnythingOfType("*gin.Context"), &reqBodyReg).Return(err)
			},
			wantErr: apperrors.ErrRequestBodyInvalid,
		},
		{
			name:       "should error on when user input invalid email / password",
			statusCode: http.StatusBadRequest,
			msg:        constants.UserInvalidEmailPassword,
			reqBody:    reqBodyLogin,
			resService: resAccessToken,
			mockService: func(usi *mocks.UserServiceInterface, vr *mocks.ValidationReqBodyInterface, reqBody dtos.RequestLoginUser, err error) {
				var reqBodyReg dtos.RequestLoginUser
				vr.On("ValidationReqBody", mock.AnythingOfType("*gin.Context"), &reqBodyReg).Return(nil)
				usi.On("PostLoginUserService", mock.AnythingOfType("*gin.Context"), reqBodyReg).Return(nil, err)
			},
			wantErr: apperrors.ErrUserInvalidEmailPassword,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService := &mocks.UserServiceInterface{}
			mockValidateReqBody := &mocks.ValidationReqBodyInterface{}

			tt.mockService(mockUserService, mockValidateReqBody, tt.reqBody.(dtos.RequestLoginUser), tt.wantErr)
			userController := controllers.NewUserController(mockUserService, mockValidateReqBody)

			g := SetUpRouterUser(userController)
			reqBody, _ := json.Marshal(tt.reqBody.(dtos.RequestLoginUser))
			req, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBuffer(reqBody))
			rec := httptest.NewRecorder()
			g.ServeHTTP(rec, req)

			var response map[string]interface{}
			json.Unmarshal(rec.Body.Bytes(), &response)

			assert.Equal(t, tt.statusCode, rec.Code)
			if tt.wantErr != nil {
				expected := dtos.ResponseMessageOnly{
					Message: tt.msg,
				}
				expectedToJson, _ := json.Marshal(expected)
				assert.Equal(t, string(expectedToJson), rec.Body.String())
			} else {
				expected := helpers.FormatterSuccessLogin(tt.resService, tt.msg)
				expectedToJson, _ := json.Marshal(expected)
				assert.Equal(t, string(expectedToJson), rec.Body.String())
				assert.Equal(t, expected.Message, response["message"])
			}
		})
	}
}

func TestUserController_GetUserDetail(t *testing.T) {
	tests := []struct {
		name        string
		statusCode  int
		msg         string
		resService  *dtos.ResponseUserAndWallet
		mockService func(c *gin.Context, usi *mocks.UserServiceInterface, err error)
		wantErr     error
	}{
		{
			name:       "should success get detail of user who is login without error",
			statusCode: http.StatusOK,
			msg:        constants.Ok,
			resService: resUserAndWallet,
			mockService: func(c *gin.Context, usi *mocks.UserServiceInterface, err error) {
				usi.On("GetUserDetailService", c, helpercontext.GetValueUserIdFromToken(c)).Return(resUserAndWallet, nil)
			},
			wantErr: nil,
		},
		{
			name:       "should error get detail of user who is login when user access with invalid token",
			statusCode: http.StatusUnauthorized,
			msg:        constants.InvalidAccessToken,
			resService: resUserAndWallet,
			mockService: func(c *gin.Context, usi *mocks.UserServiceInterface, err error) {
				usi.On("GetUserDetailService", c, helpercontext.GetValueUserIdFromToken(c)).Return(nil, err)
			},
			wantErr: apperrors.ErrInvalidAccessToken,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Set("userId", userId)

			mockUserService := &mocks.UserServiceInterface{}
			mockValidateReqBody := &mocks.ValidationReqBodyInterface{}
			tt.mockService(ctx, mockUserService, tt.wantErr)
			req := httptest.NewRequest(http.MethodGet, "/user/me", nil)
			ctx.Request = req

			userController := controllers.NewUserController(mockUserService, mockValidateReqBody)

			userController.GetUserDetail(ctx)
			middlewares.ErrorHandler(ctx)

			var response map[string]interface{}
			json.Unmarshal(rec.Body.Bytes(), &response)

			assert.Equal(t, tt.statusCode, rec.Code)
			if tt.wantErr != nil {
				expected := dtos.ResponseMessageOnly{
					Message: tt.msg,
				}
				expectedToJson, _ := json.Marshal(expected)
				assert.Equal(t, string(expectedToJson), rec.Body.String())
			} else {
				expected := helpers.FormatterMessageWithOneUserAndWallet(tt.resService, tt.msg)
				expectedToJson, _ := json.Marshal(expected)
				assert.Equal(t, string(expectedToJson), rec.Body.String())
				assert.Equal(t, expected.Message, response["message"])
			}
		})
	}
}
