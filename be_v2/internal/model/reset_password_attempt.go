package model

import (
	"time"

	"gorm.io/gorm"
)

type ResetPasswordAttempt struct {
	ResetPasswordAttemptId int64          `gorm:"primaryKey;column:reset_password_attempt_id"`
	UserId                 int64          `gorm:"column:user_id"`
	Code                   string         `gorm:"column:reset_code"`
	ExpiredAt              time.Time      `gorm:"column:expired_at"`
	CreatedAt              time.Time      `gorm:"column:created_at"`
	UpdatedAt              time.Time      `gorm:"column:updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"column:deleted_at"`
}
