package model

import (
	"time"

	"gorm.io/gorm"
)

type SourceOfFund struct {
	SourceOfFundId int64          `gorm:"primaryKey;column:source_of_fund_id"`
	Name           string         `gorm:"column:source_name"`
	CreatedAt      time.Time      `gorm:"column:created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`
}
