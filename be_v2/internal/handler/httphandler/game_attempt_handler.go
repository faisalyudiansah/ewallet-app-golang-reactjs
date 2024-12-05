package httphandler

import (
	"context"

	"ewallet-server-v2/internal/dto/httpdto"
	"ewallet-server-v2/internal/pkg/ctxutils"
	"ewallet-server-v2/internal/pkg/database"
	"ewallet-server-v2/internal/pkg/ginutils"
	"ewallet-server-v2/internal/usecase"

	"github.com/gin-gonic/gin"
)

type GameAttemptHandler struct {
	gameAttemptUsecase usecase.GameAttemptUsecase
	walletUsecase      usecase.WalletUsecase
	transactor         database.Transactor
}

func NewGameAttemptHandler(
	gameAttemptUsecase usecase.GameAttemptUsecase,
	walletUsecase usecase.WalletUsecase,
	transactor database.Transactor,
) *GameAttemptHandler {
	return &GameAttemptHandler{
		gameAttemptUsecase: gameAttemptUsecase,
		walletUsecase:      walletUsecase,
		transactor:         transactor,
	}
}

func (h *GameAttemptHandler) GetChances(c *gin.Context) {
	var res httpdto.GetChancesResponse

	userId := ctxutils.GetUserId(c)

	wallet, err := h.walletUsecase.GetOneByUserId(c, userId)
	if err != nil {
		c.Error(err)
		return
	}

	chances, err := h.gameAttemptUsecase.CalculateChances(c, wallet.WalletId)
	if err != nil {
		c.Error(err)
		return
	}

	res.Chances = chances

	ginutils.ResponseOKData(c, res)
}

func (h *GameAttemptHandler) PlayGame(c *gin.Context) {
	var req httpdto.PlayGameRequest
	var res httpdto.PlayGameResponse

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err)
		return
	}

	userId := ctxutils.GetUserId(c)

	err = h.transactor.Transaction(c, func(txCtx context.Context) error {
		wallet, err := h.walletUsecase.GetOneByUserId(txCtx, userId)
		if err != nil {
			return err
		}

		gameAttempt, err := h.gameAttemptUsecase.Attempt(txCtx, wallet.WalletId, req.GameBoxId)
		if err != nil {
			return err
		}

		res = httpdto.ConvertToPlayGameResponse(gameAttempt)

		return nil
	})
	if err != nil {
		c.Error(err)
		return
	}

	ginutils.ResponseOKData(c, res)
}
