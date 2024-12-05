package repository

import (
	"context"

	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/database"

	"gorm.io/gorm"
)

type SourceOfFundRepository interface {
	GetAll(ctx context.Context) ([]model.SourceOfFund, error)
	GetOneById(ctx context.Context, sourceId int64) (*model.SourceOfFund, error)
}

type sourceOfFundRepositoryPostgreSQL struct {
	db *database.GormWrapper
}

func NewSourceOfFundRepository(db *database.GormWrapper) *sourceOfFundRepositoryPostgreSQL {
	return &sourceOfFundRepositoryPostgreSQL{
		db: db,
	}
}

func (r *sourceOfFundRepositoryPostgreSQL) GetAll(ctx context.Context) ([]model.SourceOfFund, error) {
	res := []model.SourceOfFund{}

	if err := r.db.Start(ctx).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (r *sourceOfFundRepositoryPostgreSQL) GetOneById(ctx context.Context, sourceId int64) (*model.SourceOfFund, error) {
	res := model.SourceOfFund{}

	if err := r.db.Start(ctx).First(&res, sourceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &res, nil
}
