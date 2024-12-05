package httphandler

import (
	"context"
	"net/http"

	"ewallet-server-v2/internal/constant"
	"ewallet-server-v2/internal/dto/httpdto"
	"ewallet-server-v2/internal/pkg/database"
	"ewallet-server-v2/internal/pkg/ginutils"
	"ewallet-server-v2/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
	userUsecase usecase.UserUsecase
	transactor  database.Transactor
}

func NewAuthHandler(
	authUsecase usecase.AuthUsecase,
	userUsecase usecase.UserUsecase,
	transactor database.Transactor,
) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
		userUsecase: userUsecase,
		transactor:  transactor,
	}
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req httpdto.ResetPasswordRequest
	var res httpdto.ResetPasswordResponse

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err)
		return
	}

	err = h.transactor.Transaction(c, func(txCtx context.Context) error {
		resetPass, err := h.authUsecase.ResetPassword(txCtx, req.Email)
		if err != nil {
			return err
		}

		res = httpdto.ConvertToResetPasswordResponse(req.Email, resetPass)

		return nil
	})
	if err != nil {
		c.Error(err)
		return
	}

	ginutils.ResponseSuccessJSON(c, http.StatusCreated, constant.MessageResponseSuccess, res)
}

func (h *AuthHandler) ConfirmResetPassword(c *gin.Context) {
	var req httpdto.ConfirmResetPasswordRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err)
		return
	}

	err = h.transactor.Transaction(c, func(txCtx context.Context) error {
		err = h.authUsecase.ConfirmResetPassword(txCtx, req.Email, req.Password, req.Code)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		c.Error(err)
		return
	}

	ginutils.ResponseOKPlain(c)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req httpdto.LoginRequest
	var res httpdto.LoginResponse

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err)
		return
	}

	login, err := h.authUsecase.Login(c, req.Email, req.Password)
	if err != nil {
		c.Error(err)
		return
	}

	res = httpdto.ConvertToLoginResponse(login)

	ginutils.ResponseSuccessJSON(c, http.StatusCreated, constant.MessageResponseSuccess, res)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req httpdto.RegisterRequest
	var res httpdto.RegisterResponse

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err)
		return
	}

	err = h.transactor.Transaction(c, func(txCtx context.Context) error {
		reg, err := h.authUsecase.Register(txCtx, req.Email, req.Password, req.Username, req.FullName)
		if err != nil {
			return err
		}

		res = httpdto.ConvertToRegisterResponse(reg)
		return nil
	})
	if err != nil {
		c.Error(err)
		return
	}

	ginutils.ResponseSuccessJSON(c, http.StatusCreated, constant.MessageResponseSuccess, res)
}
