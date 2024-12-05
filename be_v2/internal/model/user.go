package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId       int64          `gorm:"primaryKey;column:user_id"`
	Name         string         `gorm:"column:user_name"`
	Password     string         `gorm:"column:user_password"`
	Email        string         `gorm:"column:email"`
	FullName     string         `gorm:"column:full_name"`
	ProfileImage string         `gorm:"column:profile_image"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
}
