package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Wallet struct {
	WalletId     int64           `gorm:"primaryKey;column:wallet_id"`
	WalletNumber string          `gorm:"column:wallet_number"`
	UserId       int64           `gorm:"column:user_id"`
	Amount       decimal.Decimal `gorm:"column:amount"`
	CreatedAt    time.Time       `gorm:"column:created_at"`
	UpdatedAt    time.Time       `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt  `gorm:"column:deleted_at"`
}
