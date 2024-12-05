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
)

func SetUpRouterWalletController(walletController *controllers.WalletController) *gin.Engine {
	h := &servers.HandlerOps{
		WalletController: walletController,
	}
	return servers.SetupRoute(h)
}

var reqTopUp = dtos.RequestTopUpWallet{
	Amount:       float64(20242024),
	SourceOfFund: 1,
}

var reqTransferFund = dtos.RequestTransferFund{
	ToWalletNumber: walletNumber,
	Amount:         reqTopUp.Amount,
	Description:    "tabungan frieren",
}

var resTransferFund = &dtos.ResponseSuccessTransfer{
	UserId:           userId,
	RecipientId:      int64(2),
	Amount:           decimal.NewFromFloat(float64(700000)),
	RemainingBelance: decimal.NewFromFloat(float64(9999999999999)),
	TransactionTime:  time.Time{},
	Description:      "frieren minta makan",
}

func TestWalletController_PutTopUpWalletController(t *testing.T) {
	tests := []struct {
		name        string
		statusCode  int
		msg         string
		reqBody     interface{}
		resService  *dtos.ResponseUserAndWallet
		mockService func(*gin.Context, *mocks.WalletServiceInterface, *mocks.ValidationReqBodyInterface, dtos.RequestTopUpWallet, error)
		wantErr     error
	}{
		{
			name:       "should success top up when user hit api top up without error",
			statusCode: http.StatusOK,
			msg:        constants.UserWalletSuccessTopUp,
			reqBody:    reqTopUp,
			resService: resUserAndWallet,
			mockService: func(c *gin.Context, wsi *mocks.WalletServiceInterface, vrbi *mocks.ValidationReqBodyInterface, rtuw dtos.RequestTopUpWallet, err error) {
				var reqBodyReg dtos.RequestTopUpWallet
				vrbi.On("ValidationReqBody", c, &reqBodyReg).Return(nil)
				wsi.On("PutTopWalletService", c, reqBodyReg, helpercontext.GetValueUserIdFromToken(c)).Return(resUserAndWallet, nil)
			},
			wantErr: nil,
		},
		{
			name:       "should error on validate request body when user input request on top up api",
			statusCode: http.StatusBadRequest,
			msg:        constants.RequestBodyInvalid,
			reqBody:    reqTopUp,
			resService: nil,
			mockService: func(c *gin.Context, wsi *mocks.WalletServiceInterface, vrbi *mocks.ValidationReqBodyInterface, rtuw dtos.RequestTopUpWallet, err error) {
				var reqBodyReg dtos.RequestTopUpWallet
				vrbi.On("ValidationReqBody", c, &reqBodyReg).Return(err)
			},
			wantErr: apperrors.ErrRequestBodyInvalid,
		},
		{
			name:       "should error when source of fund id not exists",
			statusCode: http.StatusBadRequest,
			msg:        constants.SOFIdIsNotExists,
			reqBody:    reqTopUp,
			resService: nil,
			mockService: func(c *gin.Context, wsi *mocks.WalletServiceInterface, vrbi *mocks.ValidationReqBodyInterface, rtuw dtos.RequestTopUpWallet, err error) {
				var reqBodyReg dtos.RequestTopUpWallet
				vrbi.On("ValidationReqBody", c, &reqBodyReg).Return(nil)
				wsi.On("PutTopWalletService", c, reqBodyReg, helpercontext.GetValueUserIdFromToken(c)).Return(nil, err)
			},
			wantErr: apperrors.ErrSOFIdIsNotExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Set("userId", userId)

			mockWalletService := &mocks.WalletServiceInterface{}
			mockValidateReqBody := &mocks.ValidationReqBodyInterface{}

			tt.mockService(ctx, mockWalletService, mockValidateReqBody, tt.reqBody.(dtos.RequestTopUpWallet), tt.wantErr)
			reqBody, _ := json.Marshal(tt.reqBody.(dtos.RequestTopUpWallet))
			req := httptest.NewRequest(http.MethodPut, "/user/top-up/wallet", bytes.NewBuffer(reqBody))
			ctx.Request = req

			walletController := controllers.NewWalletController(mockWalletService, mockValidateReqBody)

			walletController.PutTopUpWalletController(ctx)
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

func TestWalletController_PostTransferFundController(t *testing.T) {
	tests := []struct {
		name        string
		statusCode  int
		msg         string
		reqBody     interface{}
		resService  *dtos.ResponseSuccessTransfer
		mockService func(*gin.Context, *mocks.WalletServiceInterface, *mocks.ValidationReqBodyInterface, dtos.RequestTransferFund, error)
		wantErr     error
	}{
		{
			name:       "should success top up when user hit api top up without error",
			statusCode: http.StatusOK,
			msg:        constants.WalletSuccessTransfer,
			reqBody:    reqTransferFund,
			resService: resTransferFund,
			mockService: func(c *gin.Context, wsi *mocks.WalletServiceInterface, vrbi *mocks.ValidationReqBodyInterface, rtuw dtos.RequestTransferFund, err error) {
				var reqBodyReg dtos.RequestTransferFund
				vrbi.On("ValidationReqBody", c, &reqBodyReg).Return(nil)
				wsi.On("PostTransferFundService", c, reqBodyReg, helpercontext.GetValueUserIdFromToken(c)).Return(resTransferFund, nil)
			},
			wantErr: nil,
		},
		{
			name:       "should error on validate request body when user input request on top up api",
			statusCode: http.StatusBadRequest,
			msg:        constants.RequestBodyInvalid,
			reqBody:    reqTransferFund,
			resService: nil,
			mockService: func(c *gin.Context, wsi *mocks.WalletServiceInterface, vrbi *mocks.ValidationReqBodyInterface, rtuw dtos.RequestTransferFund, err error) {
				var reqBodyReg dtos.RequestTransferFund
				vrbi.On("ValidationReqBody", c, &reqBodyReg).Return(err)
			},
			wantErr: apperrors.ErrRequestBodyInvalid,
		},
		{
			name:       "should error transfer when balance's user is insufficient",
			statusCode: http.StatusBadRequest,
			msg:        constants.WalletBalanceIsInsufficient,
			reqBody:    reqTransferFund,
			resService: nil,
			mockService: func(c *gin.Context, wsi *mocks.WalletServiceInterface, vrbi *mocks.ValidationReqBodyInterface, rtuw dtos.RequestTransferFund, err error) {
				var reqBodyReg dtos.RequestTransferFund
				vrbi.On("ValidationReqBody", c, &reqBodyReg).Return(nil)
				wsi.On("PostTransferFundService", c, reqBodyReg, helpercontext.GetValueUserIdFromToken(c)).Return(nil, err)
			},
			wantErr: apperrors.ErrWalletBalanceIsInsufficient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Set("userId", userId)

			mockWalletService := &mocks.WalletServiceInterface{}
			mockValidateReqBody := &mocks.ValidationReqBodyInterface{}

			tt.mockService(ctx, mockWalletService, mockValidateReqBody, tt.reqBody.(dtos.RequestTransferFund), tt.wantErr)
			reqBody, _ := json.Marshal(tt.reqBody.(dtos.RequestTransferFund))
			req := httptest.NewRequest(http.MethodPost, "/transactions/transfers", bytes.NewBuffer(reqBody))
			ctx.Request = req

			walletController := controllers.NewWalletController(mockWalletService, mockValidateReqBody)

			walletController.PostTransferFundController(ctx)
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
				expected := helpers.FormatterSuccessTransfer(tt.resService, tt.msg)
				expectedToJson, _ := json.Marshal(expected)
				assert.Equal(t, string(expectedToJson), rec.Body.String())
				assert.Equal(t, expected.Message, response["message"])
			}
		})
	}
}
