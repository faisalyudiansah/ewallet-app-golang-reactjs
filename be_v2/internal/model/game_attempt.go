package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type GameAttempt struct {
	GameAttemptId int64           `gorm:"primaryKey;column:game_attempt_id"`
	WalletId      int64           `gorm:"column:wallet_id"`
	Amount        decimal.Decimal `gorm:"column:amount"`
	GameBoxesid   int64           `gorm:"column:game_boxes_id"`
	CreatedAt     time.Time       `gorm:"column:created_at"`
	UpdatedAt     time.Time       `gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt  `gorm:"column:deleted_at"`
}
