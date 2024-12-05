package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type GameBox struct {
	GameBoxId int64           `gorm:"primaryKey;column:game_boxes_id"`
	Amount    decimal.Decimal `gorm:"column:amount"`
	CreatedAt time.Time       `gorm:"column:created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`
}
