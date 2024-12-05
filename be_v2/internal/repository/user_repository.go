package repository

import (
	"context"
	"errors"

	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/database"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateOne(ctx context.Context, user model.User) (*model.User, error)
	GetOneById(ctx context.Context, userId int64) (*model.User, error)
	GetOneByEmail(ctx context.Context, email string) (*model.User, error)
	GetOneByWalletId(ctx context.Context, walletId int64) (*model.User, error)
	SaveOne(ctx context.Context, user model.User) (*model.User, error)
	UpdateOne(ctx context.Context, user model.User) (*model.User, error)
}

type userRepositoryPostgreSQL struct {
	db *database.GormWrapper
}

func NewUserRepository(db *database.GormWrapper) *userRepositoryPostgreSQL {
	return &userRepositoryPostgreSQL{
		db: db,
	}
}

func (r *userRepositoryPostgreSQL) CreateOne(ctx context.Context, user model.User) (*model.User, error) {
	if err := r.db.Start(ctx).Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryPostgreSQL) GetOneById(ctx context.Context, userId int64) (*model.User, error) {
	var user model.User

	if err := r.db.Start(ctx).First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryPostgreSQL) GetOneByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	if err := r.db.Start(ctx).Where("email", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryPostgreSQL) GetOneByWalletId(ctx context.Context, walletId int64) (*model.User, error) {
	var user model.User

	q := r.db.Start(ctx).
		Model(&model.User{}).
		Joins("wallets w ON w.user_id = users.user_id").
		Where("w.wallet_id", walletId)

	if err := q.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryPostgreSQL) SaveOne(ctx context.Context, user model.User) (*model.User, error) {
	if err := r.db.Start(ctx).Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryPostgreSQL) UpdateOne(ctx context.Context, user model.User) (*model.User, error) {
	if err := r.db.Start(ctx).Model(&user).Where("user_id = ?", user.UserId).Updates(model.User{
		Email:        user.Email,
		FullName:     user.FullName,
		ProfileImage: user.ProfileImage,
	}).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
