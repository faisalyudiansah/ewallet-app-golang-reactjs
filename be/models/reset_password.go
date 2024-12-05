package models

import "time"

type ResetPassword struct {
	ID        int64
	UserId    int64
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *time.Time
}
