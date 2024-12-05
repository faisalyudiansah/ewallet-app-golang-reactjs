package controllers

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/constants"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/dtos"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/helpers"
	helpercontext "git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/helpers/helper_context"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/services"
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
