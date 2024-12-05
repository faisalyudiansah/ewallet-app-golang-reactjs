package models

import "time"

type SourceOfFund struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *time.Time
}
