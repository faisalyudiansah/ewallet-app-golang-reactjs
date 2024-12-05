package httphandler

import (
	"ewallet-server-v2/internal/dto/httpdto"
	"ewallet-server-v2/internal/pkg/ginutils"
	"ewallet-server-v2/internal/usecase"

	"github.com/gin-gonic/gin"
)

type GameBoxHandler struct {
	gameBoxUsecase usecase.GameBoxUsecase
}

func NewGameBoxHandler(gameBoxUsecase usecase.GameBoxUsecase) *GameBoxHandler {
	return &GameBoxHandler{
		gameBoxUsecase: gameBoxUsecase,
	}
}

func (h *GameBoxHandler) GetAllBoxes(c *gin.Context) {
	boxes, err := h.gameBoxUsecase.GetAll(c)
	if err != nil {
		c.Error(err)
		return
	}

	boxesDto := httpdto.ConvertToGetAllBoxesResponse(boxes)

	ginutils.ResponseOKData(c, boxesDto)
}
