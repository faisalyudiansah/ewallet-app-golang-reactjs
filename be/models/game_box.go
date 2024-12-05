package models

import "time"

type GameBox struct {
	ID        int64
	UserID    int64
	IsOpen    bool
	Box1      float64
	Box2      float64
	Box3      float64
	Box4      float64
	Box5      float64
	Box6      float64
	Box7      float64
	Box8      float64
	Box9      float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *time.Time
}
