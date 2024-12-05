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

type GameController struct {
	GameService       services.GameServiceInterface
	ValidationReqBody helpers.ValidationReqBodyInterface
}

func NewGameController(gs services.GameServiceInterface, vrb helpers.ValidationReqBodyInterface) *GameController {
	return &GameController{
		GameService:       gs,
		ValidationReqBody: vrb,
	}
}

func (gc *GameController) PostStartGameController(c *gin.Context) {
	rewards, err := gc.GameService.PostStartGameService(c, helpercontext.GetValueUserIdFromToken(c))
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusCreated, helpers.FormatterSuccessCreateGachaBox(rewards, constants.GameSuccessGenerateGachaBox))
}

func (gc *GameController) PostChooseGachaBoxController(c *gin.Context) {
	var reqBody dtos.RequestChooseBox
	if err := gc.ValidationReqBody.ValidationReqBody(c, &reqBody); err != nil {
		c.Error(err)
		return
	}
	rewards, err := gc.GameService.PostChooseGachaBoxService(c, helpercontext.GetValueUserIdFromToken(c), reqBody.BoxIndex)
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusOK, helpers.FormatterSuccessChooseGame(rewards, constants.GameSuccessChooseGachaBox))
}
