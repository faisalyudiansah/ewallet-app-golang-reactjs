package model

import (
	"time"

	"gorm.io/gorm"
)

type TransactionType struct {
	TransactionTypeId int64          `gorm:"primaryKey;column:transaction_type_id"`
	TypeName          string         `gorm:"column:type_name"`
	CreatedAt         time.Time      `gorm:"column:created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at"`
}
