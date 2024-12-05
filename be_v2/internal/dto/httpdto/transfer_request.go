package httpdto

import "github.com/shopspring/decimal"

type TransferRequest struct {
	WalletTo    string          `json:"wallet_to" binding:"required,len=13"`
	Description string          `json:"description" binding:"max=32"`
	Amount      decimal.Decimal `json:"amount" binding:"required,dgte=1000,dlte=50000000"`
}
