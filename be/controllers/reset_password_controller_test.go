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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func SetUpRouterResetPassword(resetController *controllers.ResetPassowordController) *gin.Engine {
	h := &servers.HandlerOps{
		ResetPasswordController: resetController,
	}
	return servers.SetupRoute(h)
}

var reqForgetPwd = dtos.RequestForgetPassword{
	Email: emailUser,
}

var reqResetPwd = dtos.RequestResetPassword{
	Email:       emailUser,
	NewPassword: pwdUser,
}

var resForgetPwdService = &dtos.ResponseTokenResetPassword{
	TokenResetPassword: "frieren sihirin saya",
	LinkResetPassword:  "frieren.com",
}

var resUser = &dtos.ResponseUser{
	ID:         userId,
	Email:      emailUser,
	ChanceGame: 0,
	FullName:   fullnameUser,
	BirthDate:  time.Time{},
	CreatedAt:  time.Time{},
	UpdatedAt:  time.Time{},
	DeleteAt:   &time.Time{},
}

func TestResetPassowordController_PostForgetPasswordController(t *testing.T) {
	tests := []struct {
		name        string
		statusCode  int
		msg         string
		reqBody     interface{}
		resService  *dtos.ResponseTokenResetPassword
		mockService func(*mocks.ResetPasswordInterface, *mocks.ValidationReqBodyInterface, *mocks.GetParamInterface, dtos.RequestForgetPassword, error)
		wantErr     error
	}{
		{
			name:       "should success get token for reset password on forget password api url",
			statusCode: http.StatusCreated,
			msg:        constants.UserSuccessForgetPassword,
			reqBody:    reqForgetPwd,
			resService: resForgetPwdService,
			mockService: func(rpi *mocks.ResetPasswordInterface, vrbi *mocks.ValidationReqBodyInterface, gpi *mocks.GetParamInterface, rfp dtos.RequestForgetPassword, err error) {
				var reqBodyReg dtos.RequestForgetPassword
				vrbi.On("ValidationReqBody", mock.AnythingOfType("*gin.Context"), &reqBodyReg).Return(nil)
				rpi.On("PostForgetPasswordService", mock.AnythingOfType("*gin.Context"), reqBodyReg).Return(resForgetPwdService, nil)
			},
			wantErr: nil,
		},
		{
			name:       "should error on validate request body when user input request on forget password",
			statusCode: http.StatusBadRequest,
			msg:        constants.RequestBodyInvalid,
			reqBody:    reqForgetPwd,
			resService: nil,
			mockService: func(rpi *mocks.ResetPasswordInterface, vrbi *mocks.ValidationReqBodyInterface, gpi *mocks.GetParamInterface, rfp dtos.RequestForgetPassword, err error) {
				var reqBodyReg dtos.RequestForgetPassword
				vrbi.On("ValidationReqBody", mock.AnythingOfType("*gin.Context"), &reqBodyReg).Return(err)
			},
			wantErr: apperrors.ErrRequestBodyInvalid,
		},
		{
			name:       "should error when email does not exists",
			statusCode: http.StatusBadRequest,
			msg:        constants.UserEmailNotExists,
			reqBody:    reqForgetPwd,
			resService: nil,
			mockService: func(rpi *mocks.ResetPasswordInterface, vrbi *mocks.ValidationReqBodyInterface, gpi *mocks.GetParamInterface, rfp dtos.RequestForgetPassword, err error) {
				var reqBodyReg dtos.RequestForgetPassword
				vrbi.On("ValidationReqBody", mock.AnythingOfType("*gin.Context"), &reqBodyReg).Return(nil)
				rpi.On("PostForgetPasswordService", mock.AnythingOfType("*gin.Context"), reqBodyReg).Return(nil, err)
			},
			wantErr: apperrors.ErrUserEmailNotExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockResetPasswordService := &mocks.ResetPasswordInterface{}
			mockValidateReqBody := &mocks.ValidationReqBodyInterface{}
			mockGetParam := &mocks.GetParamInterface{}
			tt.mockService(mockResetPasswordService, mockValidateReqBody, mockGetParam, tt.reqBody.(dtos.RequestForgetPassword), tt.wantErr)
			resetController := controllers.NewResetPassowordController(mockResetPasswordService, mockValidateReqBody, mockGetParam)

			g := SetUpRouterResetPassword(resetController)
			reqBody, _ := json.Marshal(tt.reqBody.(dtos.RequestForgetPassword))
			req, _ := http.NewRequest(http.MethodPost, "/user/forget-password", bytes.NewBuffer(reqBody))
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
				expected := helpers.FormatterSuccessForgetPassword(tt.resService, tt.msg)
				expectedToJson, _ := json.Marshal(expected)
				assert.Equal(t, string(expectedToJson), rec.Body.String())
				assert.Equal(t, expected.Message, response["message"])
			}
		})
	}
}

func TestResetPassowordController_PutResetPasswordController(t *testing.T) {
	tests := []struct {
		name        string
		statusCode  int
		msg         string
		reqBody     interface{}
		resService  *dtos.ResponseUser
		mockService func(*gin.Context, *mocks.ResetPasswordInterface, *mocks.ValidationReqBodyInterface, *mocks.GetParamInterface, dtos.RequestResetPassword, error)
		wantErr     error
	}{
		{
			name:       "should success reset password user without error",
			statusCode: http.StatusOK,
			msg:        constants.UserSuccessResetPassword,
			reqBody:    reqResetPwd,
			resService: resUser,
			mockService: func(c *gin.Context, rpi *mocks.ResetPasswordInterface, vrbi *mocks.ValidationReqBodyInterface, gpi *mocks.GetParamInterface, rfp dtos.RequestResetPassword, err error) {
				var reqBodyReg dtos.RequestResetPassword
				vrbi.On("ValidationReqBody", c, &reqBodyReg).Return(nil)
				rpi.On("PutResetPasswordService", c, mock.Anything, reqBodyReg, helpercontext.GetValueUserIdFromToken(c)).Return(resUser, nil)
			},
			wantErr: nil,
		},
		{
			name:       "should error on validate request body when user input request on reset password",
			statusCode: http.StatusBadRequest,
			msg:        constants.RequestBodyInvalid,
			reqBody:    reqResetPwd,
			resService: nil,
			mockService: func(c *gin.Context, rpi *mocks.ResetPasswordInterface, vrbi *mocks.ValidationReqBodyInterface, gpi *mocks.GetParamInterface, rfp dtos.RequestResetPassword, err error) {
				var reqBodyReg dtos.RequestResetPassword
				vrbi.On("ValidationReqBody", c, &reqBodyReg).Return(err)
			},
			wantErr: apperrors.ErrRequestBodyInvalid,
		},
		{
			name:       "should error when email does not exists",
			statusCode: http.StatusBadRequest,
			msg:        constants.UserEmailNotExists,
			reqBody:    reqResetPwd,
			resService: nil,
			mockService: func(c *gin.Context, rpi *mocks.ResetPasswordInterface, vrbi *mocks.ValidationReqBodyInterface, gpi *mocks.GetParamInterface, rfp dtos.RequestResetPassword, err error) {
				var reqBodyReg dtos.RequestResetPassword
				vrbi.On("ValidationReqBody", c, &reqBodyReg).Return(nil)
				rpi.On("PutResetPasswordService", c, mock.Anything, reqBodyReg, helpercontext.GetValueUserIdFromToken(c)).Return(nil, err)
			},
			wantErr: apperrors.ErrUserEmailNotExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Set("userId", userId)

			mockResetPasswordService := &mocks.ResetPasswordInterface{}
			mockValidateReqBody := &mocks.ValidationReqBodyInterface{}
			mockGetParam := &mocks.GetParamInterface{}

			tt.mockService(ctx, mockResetPasswordService, mockValidateReqBody, mockGetParam, tt.reqBody.(dtos.RequestResetPassword), tt.wantErr)
			reqBody, _ := json.Marshal(tt.reqBody.(dtos.RequestResetPassword))
			req := httptest.NewRequest(http.MethodPut, "/user/reset-password/202407111523046", bytes.NewBuffer(reqBody))
			ctx.Request = req

			resetController := controllers.NewResetPassowordController(mockResetPasswordService, mockValidateReqBody, mockGetParam)

			resetController.PutResetPasswordController(ctx)
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
				expected := helpers.FormatterMessageWithOneUser(tt.resService, tt.msg)
				expectedToJson, _ := json.Marshal(expected)
				assert.Equal(t, string(expectedToJson), rec.Body.String())
				assert.Equal(t, expected.Message, response["message"])
			}
		})
	}
}
