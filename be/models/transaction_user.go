package models

import "time"

type TransactionUser struct {
	ID            int64
	UserId        int64
	TransactionId int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeleteAt      *time.Time
}
