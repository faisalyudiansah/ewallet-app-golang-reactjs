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
