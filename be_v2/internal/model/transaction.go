package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	TransactionId                 int64           `gorm:"primaryKey;column:transaction_id"`
	TransactionRefId              *int64          `gorm:"column:transaction_ref_id"`
	WalletId                      int64           `gorm:"column:wallet_id"`
	TransactionTypeId             int64           `gorm:"column:transaction_type_id"`
	TransactionAdditionalDetailId int64           `gorm:"column:transaction_additional_detail_id"`
	Amount                        decimal.Decimal `gorm:"column:amount"`
	Description                   string          `gorm:"column:transaction_description"`
	CreatedAt                     time.Time       `gorm:"column:created_at"`
	UpdatedAt                     time.Time       `gorm:"column:updated_at"`
	DeletedAt                     gorm.DeletedAt  `gorm:"column:deleted_at"`
}
