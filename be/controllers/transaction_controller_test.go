package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ewallet-server-v1/apperrors"
	"ewallet-server-v1/constants"
	"ewallet-server-v1/controllers"
	"ewallet-server-v1/dtos"
	"ewallet-server-v1/helpers"
	helpercontext "ewallet-server-v1/helpers/helper_context"
	"ewallet-server-v1/middlewares"
	"ewallet-server-v1/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTransactionUserController_GetListTransactionsController(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		statusCode  int
		msg         string
		resService  []dtos.ResponseSingleDataTransactionUser
		mockService func(*gin.Context, *mocks.TransactionServiceInterface, error)
		wantErr     error
	}{
		{
			name:       "should successful in getting a list of transactions of users who are currently logged in",
			url:        "/transactions?s=cash&sortBy=amount&sort=asc&limit=3&page=1&startDate=2024-01-01&endDate=2024-12-12",
			statusCode: http.StatusOK,
			msg:        constants.Ok,
			resService: []dtos.ResponseSingleDataTransactionUser{},
			mockService: func(c *gin.Context, tsi *mocks.TransactionServiceInterface, err error) {
				tsi.On("GetListTransactionsUserService", c, helpercontext.GetValueUserIdFromToken(c), "cash", "amount", "asc", int64(3), int64(0), "2024-01-01", "2024-12-12").Return([]dtos.ResponseSingleDataTransactionUser{}, int64(4), nil)
			},
			wantErr: nil,
		},
		{
			name:       "should error when something wrong in system to load list transactions user",
			url:        "/transactions?s=cash&sortBy=amount&sort=asc&limit=3&page=1&startDate=2024-01-01&endDate=2024-12-12",
			statusCode: http.StatusInternalServerError,
			msg:        constants.ISE,
			resService: []dtos.ResponseSingleDataTransactionUser{},
			mockService: func(c *gin.Context, tsi *mocks.TransactionServiceInterface, err error) {
				tsi.On("GetListTransactionsUserService", c, helpercontext.GetValueUserIdFromToken(c), "cash", "amount", "asc", int64(3), int64(0), "2024-01-01", "2024-12-12").Return(nil, int64(0), err)
			},
			wantErr: apperrors.ErrISE,
		},
		{
			name:       "should error when user input invalid date in query startDate and endDate",
			url:        "/transactions?s=cash&sortBy=amount&sort=asc&limit=3&page=1&startDate=202401&endDate=202412",
			statusCode: http.StatusBadRequest,
			msg:        constants.InvalidDateFormat,
			resService: []dtos.ResponseSingleDataTransactionUser{},
			mockService: func(c *gin.Context, tsi *mocks.TransactionServiceInterface, err error) {
			},
			wantErr: apperrors.ErrInvalidDateFormat,
		},
		{
			name:       "should error when user input invalid date in query startDate and endDate",
			url:        "/transactions?s=cash&sortBy=amount&sort=asc&limit=SALAH&page=1&startDate=2024-01-01&endDate=2024-12-12",
			statusCode: http.StatusBadRequest,
			msg:        constants.InvalidQueryLimit,
			resService: []dtos.ResponseSingleDataTransactionUser{},
			mockService: func(c *gin.Context, tsi *mocks.TransactionServiceInterface, err error) {
			},
			wantErr: apperrors.ErrInvalidQueryLimit,
		},
		{
			name:       "should error when user input invalid date in query startDate and endDate",
			url:        "/transactions?s=cash&sortBy=amount&sort=asc&limit=3&page=INISALAH&startDate=2024-01-01&endDate=2024-12-12",
			statusCode: http.StatusBadRequest,
			msg:        constants.InvalidQueryPage,
			resService: []dtos.ResponseSingleDataTransactionUser{},
			mockService: func(c *gin.Context, tsi *mocks.TransactionServiceInterface, err error) {
			},
			wantErr: apperrors.ErrInvalidQueryPage,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Set("userId", userId)

			mockTransactionService := &mocks.TransactionServiceInterface{}
			mockGetParam := &mocks.GetParamInterface{}

			tt.mockService(ctx, mockTransactionService, tt.wantErr)
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			ctx.Request = req

			txController := controllers.NewTransactionUserController(mockTransactionService, mockGetParam)

			txController.GetListTransactionsController(ctx)
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
				expected := helpers.FormatterSuccessGetListTransaction(tt.resService, tt.msg, int64(3), int64(1), int64(2), int64(4))
				expectedToJson, _ := json.Marshal(expected)
				assert.Equal(t, string(expectedToJson), rec.Body.String())
				assert.Equal(t, expected.Message, response["message"])
			}
		})
	}
}
