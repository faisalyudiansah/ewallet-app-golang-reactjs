package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID              int64
	SourceId        int64
	RecipientId     int64
	TransactionTime time.Time
	Amount          decimal.Decimal
	Description     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeleteAt        *time.Time
}
