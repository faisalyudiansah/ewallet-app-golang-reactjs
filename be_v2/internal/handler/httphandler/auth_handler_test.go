package httphandler

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"ewallet-server-v2/internal/dto/authdto"
	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewAuthHandler(t *testing.T) {
	dep := NewAuthHandler(nil, nil, nil)
	assert.NotNil(t, dep)
}

func TestAuthHandler_ResetPassword(t *testing.T) {
	type fields struct {
		authUsecase *mocks.AuthUsecase
		userUsecase *mocks.UserUsecase
		transactor  *mocks.Transactor
	}
	tests := []struct {
		name        string
		body        io.Reader
		mock        func(c *gin.Context, f fields)
		wantCode    int
		wantMessage string
	}{
		{
			name: "success 1",
			body: bytes.NewReader([]byte(`{"email":"lala@mail.com"}`)),
			mock: func(c *gin.Context, f fields) {
				f.transactor.On(
					"Transaction",
					c,
					mock.Anything,
				).Return(nil)
			},
			wantCode:    http.StatusCreated,
			wantMessage: `{"message":"success","data":{"email":"","code":"","expired_at":"0001-01-01T00:00:00Z"}}`,
		},
		{
			name: "success 2",
			body: bytes.NewReader([]byte(`{"email":"lala@mail.com"}`)),
			mock: func(c *gin.Context, f fields) {
				f.authUsecase.On(
					"ResetPassword",
					mock.Anything,
					"lala@mail.com",
				).Return(
					&model.ResetPasswordAttempt{
						Code:      "token",
						ExpiredAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					}, nil,
				)

				f.transactor.On(
					"Transaction",
					c,
					mock.MatchedBy(func(txFunc func(context.Context) error) bool {
						err := txFunc(context.Background())
						return err == nil
					}),
				).Return(nil)
			},
			wantCode:    http.StatusCreated,
			wantMessage: `{"message":"success","data":{"email":"lala@mail.com","code":"token","expired_at":"2022-01-01T00:00:00Z"}}`,
		},
		{
			name: "error ResetPassword",
			body: bytes.NewReader([]byte(`{"email":"lala@mail.com"}`)),
			mock: func(c *gin.Context, f fields) {
				f.authUsecase.On(
					"ResetPassword",
					mock.Anything,
					"lala@mail.com",
				).Return(
					nil, fmt.Errorf("error"),
				)

				f.transactor.On(
					"Transaction",
					c,
					mock.MatchedBy(func(txFunc func(context.Context) error) bool {
						err := txFunc(context.Background())
						return err != nil
					}),
				).Return(fmt.Errorf("error"))
			},
			wantCode:    http.StatusOK,
			wantMessage: "",
		},
		{
			name: "error transactor",
			body: bytes.NewReader([]byte(`{"email":"lala@mail.com"}`)),
			mock: func(c *gin.Context, f fields) {
				f.transactor.On(
					"Transaction",
					c,
					mock.Anything,
				).Return(fmt.Errorf("error"))
			},
			wantCode:    http.StatusOK,
			wantMessage: "",
		},
		{
			name:        "error bind JSON",
			body:        bytes.NewReader([]byte(`{"emails":"lala@mail.com"}`)),
			mock:        func(c *gin.Context, f fields) {},
			wantCode:    http.StatusOK,
			wantMessage: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDep := fields{
				authUsecase: mocks.NewAuthUsecase(t),
				userUsecase: mocks.NewUserUsecase(t),
				transactor:  mocks.NewTransactor(t),
			}

			h := &AuthHandler{
				authUsecase: mockDep.authUsecase,
				userUsecase: mockDep.userUsecase,
				transactor:  mockDep.transactor,
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			r := httptest.NewRequest("POST", "/auth/reset-passwords", tt.body)
			c.Request = r

			tt.mock(c, mockDep)

			h.ResetPassword(c)

			assert.Equal(t, tt.wantCode, w.Result().StatusCode)
			assert.Equal(t, tt.wantMessage, w.Body.String())
		})
	}
}

func TestAuthHandler_ConfirmResetPassword(t *testing.T) {
	type fields struct {
		authUsecase *mocks.AuthUsecase
		userUsecase *mocks.UserUsecase
		transactor  *mocks.Transactor
	}

	tests := []struct {
		name        string
		body        io.Reader
		mock        func(c *gin.Context, f fields)
		wantCode    int
		wantMessage string
	}{
		{
			name: "success 1",
			body: bytes.NewReader([]byte(`{"code":"token","email":"lala@mail.com","password":"password"}`)),
			mock: func(c *gin.Context, f fields) {
				f.transactor.On(
					"Transaction",
					c,
					mock.Anything,
				).Return(nil)
			},
			wantCode:    http.StatusOK,
			wantMessage: `{"message":"success"}`,
		},
		{
			name: "success 2",
			body: bytes.NewReader([]byte(`{"code":"token","email":"lala@mail.com","password":"password"}`)),
			mock: func(c *gin.Context, f fields) {
				f.authUsecase.On(
					"ConfirmResetPassword",
					mock.Anything,
					"lala@mail.com",
					"password",
					"token",
				).Return(nil)

				f.transactor.On(
					"Transaction",
					c,
					mock.MatchedBy(func(txFunc func(context.Context) error) bool {
						err := txFunc(context.Background())
						return err == nil
					}),
				).Return(nil)
			},
			wantCode:    http.StatusOK,
			wantMessage: `{"message":"success"}`,
		},
		{
			name: "error ConfirmResetPassword",
			body: bytes.NewReader([]byte(`{"code":"token","email":"lala@mail.com","password":"password"}`)),
			mock: func(c *gin.Context, f fields) {
				f.authUsecase.On(
					"ConfirmResetPassword",
					mock.Anything,
					"lala@mail.com",
					"password",
					"token",
				).Return(fmt.Errorf("error"))

				f.transactor.On(
					"Transaction",
					c,
					mock.MatchedBy(func(txFunc func(context.Context) error) bool {
						err := txFunc(context.Background())
						return err != nil
					}),
				).Return(fmt.Errorf("error"))
			},
			wantCode:    http.StatusOK,
			wantMessage: "",
		},
		{
			name: "error transactor",
			body: bytes.NewReader([]byte(`{"code":"token","email":"lala@mail.com","password":"password"}`)),
			mock: func(c *gin.Context, f fields) {
				f.transactor.On(
					"Transaction",
					c,
					mock.Anything,
				).Return(fmt.Errorf("error"))
			},
			wantCode:    http.StatusOK,
			wantMessage: "",
		},
		{
			name:        "error bind JSON",
			body:        bytes.NewReader([]byte(`{"codes":"token","emails":"lala@mail.com","passwords":"password"}`)),
			mock:        func(c *gin.Context, f fields) {},
			wantCode:    http.StatusOK,
			wantMessage: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDep := fields{
				authUsecase: mocks.NewAuthUsecase(t),
				userUsecase: mocks.NewUserUsecase(t),
				transactor:  mocks.NewTransactor(t),
			}

			h := &AuthHandler{
				authUsecase: mockDep.authUsecase,
				userUsecase: mockDep.userUsecase,
				transactor:  mockDep.transactor,
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			r := httptest.NewRequest("POST", "/auth/reset-passwords", tt.body)
			c.Request = r

			tt.mock(c, mockDep)

			h.ConfirmResetPassword(c)

			assert.Equal(t, tt.wantCode, w.Result().StatusCode)
			assert.Equal(t, tt.wantMessage, w.Body.String())
		})
	}
}

func TestAuthHandler_Register(t *testing.T) {
	type fields struct {
		authUsecase *mocks.AuthUsecase
		userUsecase *mocks.UserUsecase
		transactor  *mocks.Transactor
	}

	tests := []struct {
		name        string
		body        io.Reader
		mock        func(c *gin.Context, f fields)
		wantCode    int
		wantMessage string
	}{
		{
			name: "success 1",
			body: bytes.NewReader([]byte(`{"username":"lala","email":"lala@mail.com","password":"password"}`)),
			mock: func(c *gin.Context, f fields) {
				f.transactor.On(
					"Transaction",
					c,
					mock.Anything,
				).Return(nil)
			},
			wantCode:    http.StatusCreated,
			wantMessage: `{"message":"success","data":{"username":"","email":"","wallet_id":0,"wallet_number":""}}`,
		},
		{
			name: "success 2",
			body: bytes.NewReader([]byte(`{"username":"lala","email":"lala@mail.com","password":"password"}`)),
			mock: func(c *gin.Context, f fields) {
				f.authUsecase.On(
					"Register",
					mock.Anything,
					"lala@mail.com",
					"password",
					"lala",
				).Return(&authdto.RegisterDto{
					User: model.User{
						Name:  "lala",
						Email: "lala@mail.com",
					},
					Wallet: model.Wallet{
						WalletId:     1,
						WalletNumber: "0001",
					},
				}, nil)

				f.transactor.On(
					"Transaction",
					c,
					mock.MatchedBy(func(txFunc func(context.Context) error) bool {
						err := txFunc(context.Background())
						return err == nil
					}),
				).Return(nil)
			},
			wantCode:    http.StatusCreated,
			wantMessage: `{"message":"success","data":{"username":"lala","email":"lala@mail.com","wallet_id":1,"wallet_number":"0001"}}`,
		},
		{
			name: "error Register",
			body: bytes.NewReader([]byte(`{"username":"lala","email":"lala@mail.com","password":"password"}`)),
			mock: func(c *gin.Context, f fields) {
				f.authUsecase.On(
					"Register",
					mock.Anything,
					"lala@mail.com",
					"password",
					"lala",
				).Return(nil, fmt.Errorf("error"))

				f.transactor.On(
					"Transaction",
					mock.Anything,
					mock.MatchedBy(func(txFunc func(context.Context) error) bool {
						err := txFunc(context.Background())
						return err != nil
					}),
				).Return(fmt.Errorf("error"))
			},
			wantCode:    http.StatusOK,
			wantMessage: "",
		},
		{
			name: "error transactor",
			body: bytes.NewReader([]byte(`{"username":"lala","email":"lala@mail.com","password":"password"}`)),
			mock: func(c *gin.Context, f fields) {
				f.transactor.On(
					"Transaction",
					c,
					mock.Anything,
				).Return(fmt.Errorf("error"))
			},
			wantCode:    http.StatusOK,
			wantMessage: "",
		},
		{
			name:        "error bind JSON",
			body:        bytes.NewReader([]byte(`{"usernames":"lala","emails":"lala@mail.com","passwords":"password"}`)),
			mock:        func(c *gin.Context, f fields) {},
			wantCode:    http.StatusOK,
			wantMessage: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDep := fields{
				authUsecase: mocks.NewAuthUsecase(t),
				userUsecase: mocks.NewUserUsecase(t),
				transactor:  mocks.NewTransactor(t),
			}

			h := &AuthHandler{
				authUsecase: mockDep.authUsecase,
				userUsecase: mockDep.userUsecase,
				transactor:  mockDep.transactor,
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			r := httptest.NewRequest("POST", "/auth/users", tt.body)
			c.Request = r

			tt.mock(c, mockDep)

			h.Register(c)

			assert.Equal(t, tt.wantCode, w.Result().StatusCode)
			assert.Equal(t, tt.wantMessage, w.Body.String())
		})
	}
}

func TestAuthHandler_Login(t *testing.T) {
	type fields struct {
		authUsecase *mocks.AuthUsecase
		userUsecase *mocks.UserUsecase
		transactor  *mocks.Transactor
	}

	tests := []struct {
		name        string
		body        io.Reader
		mock        func(c *gin.Context, f fields)
		wantCode    int
		wantMessage string
	}{
		{
			name: "success",
			body: bytes.NewReader([]byte(`{"username":"lala","email":"lala@mail.com","password":"password"}`)),
			mock: func(c *gin.Context, f fields) {
				f.authUsecase.On(
					"Login",
					c,
					"lala@mail.com",
					"password",
				).Return(&authdto.LoginDto{
					Token: "token",
				}, nil)
			},
			wantCode:    http.StatusCreated,
			wantMessage: `{"message":"success","data":{"token":"token"}}`,
		},
		{
			name: "error login",
			body: bytes.NewReader([]byte(`{"username":"lala","email":"lala@mail.com","password":"password"}`)),
			mock: func(c *gin.Context, f fields) {
				f.authUsecase.On(
					"Login",
					c,
					"lala@mail.com",
					"password",
				).Return(nil, fmt.Errorf("error"))
			},
			wantCode:    http.StatusOK,
			wantMessage: "",
		},
		{
			name:        "error bind JSON",
			body:        bytes.NewReader([]byte(`{"usernames":"lala","emails":"lala@mail.com","passwords":"password"}`)),
			mock:        func(c *gin.Context, f fields) {},
			wantCode:    http.StatusOK,
			wantMessage: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDep := fields{
				authUsecase: mocks.NewAuthUsecase(t),
				userUsecase: mocks.NewUserUsecase(t),
				transactor:  mocks.NewTransactor(t),
			}

			h := &AuthHandler{
				authUsecase: mockDep.authUsecase,
				userUsecase: mockDep.userUsecase,
				transactor:  mockDep.transactor,
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			r := httptest.NewRequest("POST", "/auth/sessions", tt.body)
			c.Request = r

			tt.mock(c, mockDep)

			h.Login(c)

			assert.Equal(t, tt.wantCode, w.Result().StatusCode)
			assert.Equal(t, tt.wantMessage, w.Body.String())
		})
	}
}
