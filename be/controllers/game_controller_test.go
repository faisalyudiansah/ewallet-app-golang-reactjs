package controllers_test

import (
	"bytes"
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

var resStartGame = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

var reqChooseGame = dtos.RequestChooseBox{
	BoxIndex: 1,
}

func TestGameController_PostStartGameController(t *testing.T) {
	tests := []struct {
		name        string
		statusCode  int
		msg         string
		resService  []int
		mockService func(*gin.Context, *mocks.GameServiceInterface, error)
		wantErr     error
	}{
		{
			name:       "should success when user start the game without error",
			statusCode: http.StatusCreated,
			msg:        constants.GameSuccessGenerateGachaBox,
			resService: resStartGame,
			mockService: func(c *gin.Context, gsi *mocks.GameServiceInterface, err error) {
				gsi.On("PostStartGameService", c, helpercontext.GetValueUserIdFromToken(c)).Return(resStartGame, nil)
			},
			wantErr: nil,
		},
		{
			name:       "should error when user does not have any attempt to start the game",
			statusCode: http.StatusBadRequest,
			msg:        constants.GameZeroChance,
			resService: nil,
			mockService: func(c *gin.Context, gsi *mocks.GameServiceInterface, err error) {
				gsi.On("PostStartGameService", c, helpercontext.GetValueUserIdFromToken(c)).Return(nil, err)
			},
			wantErr: apperrors.ErrGameZeroChance,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Set("userId", userId)

			mockGameService := &mocks.GameServiceInterface{}
			mockValidateReqBody := &mocks.ValidationReqBodyInterface{}

			tt.mockService(ctx, mockGameService, tt.wantErr)
			req := httptest.NewRequest(http.MethodPost, "/game/start", nil)
			ctx.Request = req

			gameController := controllers.NewGameController(mockGameService, mockValidateReqBody)

			gameController.PostStartGameController(ctx)
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
				expected := helpers.FormatterSuccessCreateGachaBox(tt.resService, tt.msg)
				expectedToJson, _ := json.Marshal(expected)
				assert.Equal(t, string(expectedToJson), rec.Body.String())
				assert.Equal(t, expected.Message, response["message"])
			}
		})
	}
}

func TestGameController_PostChooseGachaBoxController(t *testing.T) {
	tests := []struct {
		name        string
		statusCode  int
		msg         string
		reqBody     interface{}
		resService  float64
		mockService func(*gin.Context, *mocks.GameServiceInterface, *mocks.ValidationReqBodyInterface, dtos.RequestChooseBox, error)
		wantErr     error
	}{
		{
			name:       "should success choose game and get the reward when user hit api without error",
			statusCode: http.StatusOK,
			msg:        constants.GameSuccessChooseGachaBox,
			reqBody:    reqChooseGame,
			resService: float64(20000),
			mockService: func(c *gin.Context, gs *mocks.GameServiceInterface, vrbi *mocks.ValidationReqBodyInterface, rtuw dtos.RequestChooseBox, err error) {
				var reqBodyReg dtos.RequestChooseBox
				vrbi.On("ValidationReqBody", c, &reqBodyReg).Return(nil)
				gs.On("PostChooseGachaBoxService", c, helpercontext.GetValueUserIdFromToken(c), reqBodyReg.BoxIndex).Return(float64(20000), nil)
			},
			wantErr: nil,
		},
		{
			name:       "should error on validate request body when user input request on choose game api",
			statusCode: http.StatusBadRequest,
			msg:        constants.RequestBodyInvalid,
			reqBody:    reqChooseGame,
			resService: 0,
			mockService: func(c *gin.Context, gs *mocks.GameServiceInterface, vrbi *mocks.ValidationReqBodyInterface, rtuw dtos.RequestChooseBox, err error) {
				var reqBodyReg dtos.RequestChooseBox
				vrbi.On("ValidationReqBody", c, &reqBodyReg).Return(err)
			},
			wantErr: apperrors.ErrRequestBodyInvalid,
		},
		{
			name:       "should error choose box when user input wrong box index",
			statusCode: http.StatusBadRequest,
			msg:        constants.GameInvalidBoxIndex,
			reqBody:    reqChooseGame,
			resService: 0,
			mockService: func(c *gin.Context, gs *mocks.GameServiceInterface, vrbi *mocks.ValidationReqBodyInterface, rtuw dtos.RequestChooseBox, err error) {
				var reqBodyReg dtos.RequestChooseBox
				vrbi.On("ValidationReqBody", c, &reqBodyReg).Return(nil)
				gs.On("PostChooseGachaBoxService", c, helpercontext.GetValueUserIdFromToken(c), reqBodyReg.BoxIndex).Return(float64(0), err)
			},
			wantErr: apperrors.ErrGameInvalidBoxIndex,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Set("userId", userId)

			mockGameService := &mocks.GameServiceInterface{}
			mockValidateReqBody := &mocks.ValidationReqBodyInterface{}

			tt.mockService(ctx, mockGameService, mockValidateReqBody, tt.reqBody.(dtos.RequestChooseBox), tt.wantErr)
			reqBody, _ := json.Marshal(tt.reqBody.(dtos.RequestChooseBox))
			req := httptest.NewRequest(http.MethodPost, "/game/choose", bytes.NewBuffer(reqBody))
			ctx.Request = req

			gameController := controllers.NewGameController(mockGameService, mockValidateReqBody)

			gameController.PostChooseGachaBoxController(ctx)
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
				expected := helpers.FormatterSuccessChooseGame(tt.resService, tt.msg)
				expectedToJson, _ := json.Marshal(expected)
				assert.Equal(t, string(expectedToJson), rec.Body.String())
				assert.Equal(t, expected.Message, response["message"])
			}
		})
	}
}
