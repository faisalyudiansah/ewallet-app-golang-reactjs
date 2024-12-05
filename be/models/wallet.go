package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID           int64
	UserId       int64
	WalletNumber string
	Balance      decimal.Decimal
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeleteAt     *time.Time
}
