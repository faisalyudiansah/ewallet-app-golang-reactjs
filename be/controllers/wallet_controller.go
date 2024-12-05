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

type WalletController struct {
	WalletService     services.WalletServiceInterface
	ValidationReqBody helpers.ValidationReqBodyInterface
}

func NewWalletController(wc services.WalletServiceInterface, vrb helpers.ValidationReqBodyInterface) *WalletController {
	return &WalletController{
		WalletService:     wc,
		ValidationReqBody: vrb,
	}
}

func (wc *WalletController) PutTopUpWalletController(c *gin.Context) {
	var reqBody dtos.RequestTopUpWallet
	if err := wc.ValidationReqBody.ValidationReqBody(c, &reqBody); err != nil {
		c.Error(err)
		return
	}
	result, err := wc.WalletService.PutTopWalletService(c, reqBody, helpercontext.GetValueUserIdFromToken(c))
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusOK, helpers.FormatterMessageWithOneUserAndWallet(result, constants.UserWalletSuccessTopUp))
}

func (wc *WalletController) PostTransferFundController(c *gin.Context) {
	var reqBody dtos.RequestTransferFund
	if err := wc.ValidationReqBody.ValidationReqBody(c, &reqBody); err != nil {
		c.Error(err)
		return
	}
	result, err := wc.WalletService.PostTransferFundService(c, reqBody, helpercontext.GetValueUserIdFromToken(c))
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusOK, helpers.FormatterSuccessTransfer(result, constants.WalletSuccessTransfer))
}
