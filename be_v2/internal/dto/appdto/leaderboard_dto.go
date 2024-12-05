package appdto

import "github.com/shopspring/decimal"

type LeaderboardEntryDto struct {
	UserId           int64           `gorm:"column:user_id"`
	Name             string          `gorm:"column:user_name"`
	Amount           decimal.Decimal `gorm:"column:prize_sum"`
	GameAttemptCount int             `gorm:"column:game_count"`
}
