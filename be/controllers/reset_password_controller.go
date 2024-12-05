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

type ResetPassowordController struct {
	ResetPasswordService services.ResetPasswordInterface
	ValidationReqBody    helpers.ValidationReqBodyInterface
	GetParam             helpers.GetParamInterface
}

func NewResetPassowordController(rp services.ResetPasswordInterface, vrb helpers.ValidationReqBodyInterface, gp helpers.GetParamInterface) *ResetPassowordController {
	return &ResetPassowordController{
		ResetPasswordService: rp,
		ValidationReqBody:    vrb,
		GetParam:             gp,
	}
}

func (uc *ResetPassowordController) PostForgetPasswordController(c *gin.Context) {
	var reqBody dtos.RequestForgetPassword
	if err := uc.ValidationReqBody.ValidationReqBody(c, &reqBody); err != nil {
		c.Error(err)
		return
	}
	result, err := uc.ResetPasswordService.PostForgetPasswordService(c, reqBody)
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusCreated, helpers.FormatterSuccessForgetPassword(result, constants.UserSuccessForgetPassword))
}

func (uc *ResetPassowordController) PutResetPasswordController(c *gin.Context) {
	token := c.Param("token")
	var reqBody dtos.RequestResetPassword
	if err := uc.ValidationReqBody.ValidationReqBody(c, &reqBody); err != nil {
		c.Error(err)
		return
	}
	result, err := uc.ResetPasswordService.PutResetPasswordService(c, token, reqBody, helpercontext.GetValueUserIdFromToken(c))
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusOK, helpers.FormatterMessageWithOneUser(result, constants.UserSuccessResetPassword))
}
