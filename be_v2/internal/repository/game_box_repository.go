package repository

import (
	"context"

	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/database"

	"gorm.io/gorm"
)

type GameBoxRepository interface {
	GetAll(ctx context.Context, limit int) ([]model.GameBox, error)
	GetOneById(ctx context.Context, boxId int64) (*model.GameBox, error)
}

type gameBoxRepositoryPostgreSQL struct {
	db *database.GormWrapper
}

func NewGameBoxRepository(db *database.GormWrapper) *gameBoxRepositoryPostgreSQL {
	return &gameBoxRepositoryPostgreSQL{
		db: db,
	}
}

func (r *gameBoxRepositoryPostgreSQL) GetAll(ctx context.Context, limit int) ([]model.GameBox, error) {
	gameBoxes := []model.GameBox{}

	if err := r.db.Start(ctx).Limit(limit).Find(&gameBoxes).Error; err != nil {
		return nil, err
	}
	return gameBoxes, nil
}

func (r *gameBoxRepositoryPostgreSQL) GetOneById(ctx context.Context, boxId int64) (*model.GameBox, error) {
	var gameBox model.GameBox

	if err := r.db.Start(ctx).First(&gameBox, boxId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &gameBox, nil
}
