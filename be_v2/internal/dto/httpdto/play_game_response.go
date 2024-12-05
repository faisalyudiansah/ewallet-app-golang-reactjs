package httpdto

import (
	"ewallet-server-v2/internal/model"

	"github.com/shopspring/decimal"
)

type PlayGameResponse struct {
	GameAttemptId int64           `json:"game_attempt_id"`
	PrizeAmount   decimal.Decimal `json:"prize_amount"`
}

func ConvertToPlayGameResponse(gameAttempt *model.GameAttempt) PlayGameResponse {
	return PlayGameResponse{
		GameAttemptId: gameAttempt.GameAttemptId,
		PrizeAmount:   gameAttempt.Amount,
	}
}
