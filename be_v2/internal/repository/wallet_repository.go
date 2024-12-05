package repository

import (
	"context"

	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletRepository interface {
	CreateOne(ctx context.Context, wallet model.Wallet) (*model.Wallet, error)
	SaveOne(ctx context.Context, wallet model.Wallet) (*model.Wallet, error)
	GetOneByNumber(ctx context.Context, walletNumber string) (*model.Wallet, error)
	GetOneByIdWithLock(ctx context.Context, walletId int64) (*model.Wallet, error)
	GetOneByNumberWithLock(ctx context.Context, walletNumber string) (*model.Wallet, error)
	GetOneByUserId(ctx context.Context, userId int64) (*model.Wallet, error)
}

type walletRepositoryPostgreSQL struct {
	db *database.GormWrapper
}

func NewWalletRepository(db *database.GormWrapper) *walletRepositoryPostgreSQL {
	return &walletRepositoryPostgreSQL{
		db: db,
	}
}

func (r *walletRepositoryPostgreSQL) CreateOne(ctx context.Context, wallet model.Wallet) (*model.Wallet, error) {
	if err := r.db.Start(ctx).Create(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (r *walletRepositoryPostgreSQL) SaveOne(ctx context.Context, wallet model.Wallet) (*model.Wallet, error) {
	if err := r.db.Start(ctx).Save(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (r *walletRepositoryPostgreSQL) GetOneByNumber(ctx context.Context, walletNumber string) (*model.Wallet, error) {
	var wallet model.Wallet

	if err := r.db.Start(ctx).Where("wallet_number", walletNumber).First(&wallet).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &wallet, nil
}

func (r *walletRepositoryPostgreSQL) GetOneByIdWithLock(ctx context.Context, walletId int64) (*model.Wallet, error) {
	var wallet model.Wallet

	if err := r.db.Start(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
	}).First(&wallet, walletId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &wallet, nil
}

func (r *walletRepositoryPostgreSQL) GetOneByNumberWithLock(ctx context.Context, walletNumber string) (*model.Wallet, error) {
	var wallet model.Wallet

	if err := r.db.Start(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
	}).Where("wallet_number", walletNumber).First(&wallet).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &wallet, nil
}

func (r *walletRepositoryPostgreSQL) GetOneByUserId(ctx context.Context, userId int64) (*model.Wallet, error) {
	var wallet model.Wallet

	if err := r.db.Start(ctx).Where("user_id", userId).First(&wallet).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &wallet, nil
}
