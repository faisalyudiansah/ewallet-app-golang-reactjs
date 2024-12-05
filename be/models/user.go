package models

import "time"

type User struct {
	ID         int64
	Email      string
	Password   string
	ChanceGame int
	FullName   string
	BirthDate  time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteAt   *time.Time
}
