package httpdto

import "github.com/shopspring/decimal"

type TopUpRequest struct {
	SourceOfFundsId int64           `json:"source_of_funds_id" binding:"required"`
	Amount          decimal.Decimal `json:"amount" binding:"required,numeric,dgte=50000,dlte=10000000"`
}
