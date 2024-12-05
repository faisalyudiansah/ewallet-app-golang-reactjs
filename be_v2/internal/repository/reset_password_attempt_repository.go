package repository

import (
	"context"

	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/database"

	"gorm.io/gorm"
)

type ResetPasswordAttemptRepository interface {
	CreateOne(ctx context.Context, resetPasswordAttempt model.ResetPasswordAttempt) (*model.ResetPasswordAttempt, error)
	GetOneActiveByUserIdAndCode(ctx context.Context, userId int64, code string) (*model.ResetPasswordAttempt, error)
	GetAllByUserId(ctx context.Context, userId int64) ([]model.ResetPasswordAttempt, error)
	SoftDeleteByIds(ctx context.Context, attemptIds []int64) error
}

type resetPasswordAttemptRepositoryPostgreSQL struct {
	db *database.GormWrapper
}

func NewResetPasswordAttemptRepository(db *database.GormWrapper) *resetPasswordAttemptRepositoryPostgreSQL {
	return &resetPasswordAttemptRepositoryPostgreSQL{
		db: db,
	}
}

func (r *resetPasswordAttemptRepositoryPostgreSQL) CreateOne(
	ctx context.Context,
	resetPasswordAttempt model.ResetPasswordAttempt,
) (*model.ResetPasswordAttempt, error) {
	if err := r.db.Start(ctx).Create(&resetPasswordAttempt).Error; err != nil {
		return nil, err
	}

	return &resetPasswordAttempt, nil
}

func (r *resetPasswordAttemptRepositoryPostgreSQL) GetOneActiveByUserIdAndCode(
	ctx context.Context,
	userId int64,
	code string,
) (*model.ResetPasswordAttempt, error) {
	var res model.ResetPasswordAttempt

	q := r.db.Start(ctx).Where("expired_at >= NOW()").Where("user_id", userId).Where("reset_code", code)

	if err := q.First(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &res, nil
}

func (r *resetPasswordAttemptRepositoryPostgreSQL) GetAllByUserId(
	ctx context.Context,
	userId int64,
) ([]model.ResetPasswordAttempt, error) {
	var res []model.ResetPasswordAttempt

	q := r.db.Start(ctx).Where("user_id", userId)

	if err := q.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (r *resetPasswordAttemptRepositoryPostgreSQL) SoftDeleteByIds(ctx context.Context, attemptIds []int64) error {
	if err := r.db.Start(ctx).Delete(&model.ResetPasswordAttempt{}, attemptIds).Error; err != nil {
		return err
	}

	return nil
}
