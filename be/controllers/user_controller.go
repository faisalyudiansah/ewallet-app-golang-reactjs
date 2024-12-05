package controllers

import (
	"net/http"

	"ewallet-server-v1/constants"
	"ewallet-server-v1/dtos"
	"ewallet-server-v1/helpers"
	helpercontext "ewallet-server-v1/helpers/helper_context"
	"ewallet-server-v1/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService       services.UserServiceInterface
	ValidationReqBody helpers.ValidationReqBodyInterface
}

func NewUserController(us services.UserServiceInterface, vrb helpers.ValidationReqBodyInterface) *UserController {
	return &UserController{
		UserService:       us,
		ValidationReqBody: vrb,
	}
}

func (uc *UserController) PostRegisterUserController(c *gin.Context) {
	var reqBody dtos.RequestRegisterUser
	if err := uc.ValidationReqBody.ValidationReqBody(c, &reqBody); err != nil {
		c.Error(err)
		return
	}
	result, err := uc.UserService.PostRegisterUserService(c, reqBody)
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusCreated, helpers.FormatterMessageWithOneUserAndWallet(result, constants.UserSuccessRegister))
}

func (uc *UserController) PostLoginUserController(c *gin.Context) {
	var reqBody dtos.RequestLoginUser
	if err := uc.ValidationReqBody.ValidationReqBody(c, &reqBody); err != nil {
		c.Error(err)
		return
	}
	result, err := uc.UserService.PostLoginUserService(c, reqBody)
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusOK, helpers.FormatterSuccessLogin(result, constants.UserSuccessLogin))
}

func (uc *UserController) GetUserDetail(c *gin.Context) {
	result, err := uc.UserService.GetUserDetailService(c, helpercontext.GetValueUserIdFromToken(c))
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusOK, helpers.FormatterMessageWithOneUserAndWallet(result, constants.Ok))
}
