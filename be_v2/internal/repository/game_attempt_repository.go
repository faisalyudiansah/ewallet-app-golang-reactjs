package repository

import (
	"context"

	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/database"
)

type GameAttemptRepository interface {
	GetCountByWalletId(ctx context.Context, walletId int64) (int64, error)
	CreateOne(ctx context.Context, gameAttempt model.GameAttempt) (*model.GameAttempt, error)
}

type gameAttemptRepositoryPostgreSQL struct {
	db *database.GormWrapper
}

func NewGameAttemptRepository(db *database.GormWrapper) *gameAttemptRepositoryPostgreSQL {
	return &gameAttemptRepositoryPostgreSQL{
		db: db,
	}
}

func (r *gameAttemptRepositoryPostgreSQL) GetCountByWalletId(ctx context.Context, walletId int64) (int64, error) {
	var res int64

	sql := `
		SELECT
			COALESCE(COUNT(ga.game_attempt_id), 0)
		FROM
			game_attempts ga
		WHERE
			ga.wallet_id = ?
		AND ga.deleted_at IS NULL
	`

	q := r.db.Start(ctx).Raw(sql, walletId)

	if err := q.Scan(&res).Error; err != nil {
		return int64(0), err
	}

	return res, nil
}

func (r *gameAttemptRepositoryPostgreSQL) CreateOne(ctx context.Context, gameAttempt model.GameAttempt) (*model.GameAttempt, error) {
	if err := r.db.Start(ctx).Create(&gameAttempt).Debug().Error; err != nil {
		return nil, err
	}

	return &gameAttempt, nil
}
