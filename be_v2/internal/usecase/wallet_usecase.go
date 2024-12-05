package usecase

import (
	"context"

	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/apperror"
	"ewallet-server-v2/internal/pkg/apputils"
	"ewallet-server-v2/internal/repository"

	"github.com/shopspring/decimal"
)

type WalletUsecase interface {
	CreateOne(ctx context.Context, userId int64) (*model.Wallet, error)
	UpdateOneAmountById(ctx context.Context, walletId int64, amount decimal.Decimal) (*model.Wallet, error)
	GetOneByUserId(ctx context.Context, userId int64) (*model.Wallet, error)
	GetOneByNumber(ctx context.Context, walletNumber string) (*model.Wallet, error)
	GetOneByIdWithLock(ctx context.Context, walletId int64) (*model.Wallet, error)
	GetOneByNumberWithLock(ctx context.Context, walletNumber string) (*model.Wallet, error)
}

type walletUsecaseImpl struct {
	walletRepository repository.WalletRepository
	walletFormatter  apputils.WalletNumberFormatter
}

func NewWalletUsecase(
	walletRepository repository.WalletRepository,
	walletFormatter apputils.WalletNumberFormatter,
) *walletUsecaseImpl {
	return &walletUsecaseImpl{
		walletRepository: walletRepository,
		walletFormatter:  walletFormatter,
	}
}

func (u *walletUsecaseImpl) CreateOne(ctx context.Context, userId int64) (*model.Wallet, error) {
	newWallet := model.Wallet{
		UserId:       userId,
		Amount:       decimal.NewFromFloat(0),
		WalletNumber: u.walletFormatter.Format(userId),
	}

	createdWallet, err := u.walletRepository.CreateOne(ctx, newWallet)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	return createdWallet, nil
}

func (u *walletUsecaseImpl) UpdateOneAmountById(ctx context.Context, walletId int64, amount decimal.Decimal) (*model.Wallet, error) {
	wallet, err := u.walletRepository.GetOneByIdWithLock(ctx, walletId)
	if err != nil {
		return nil, apperror.NewServerError(err)
	} else if wallet == nil {
		return nil, apperror.NewEntityNotFoundError(nil, "wallet")
	}

	wallet.Amount = amount

	updatedWallet, err := u.walletRepository.SaveOne(ctx, *wallet)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	return updatedWallet, nil
}

func (u *walletUsecaseImpl) GetOneByUserId(ctx context.Context, userId int64) (*model.Wallet, error) {
	wallet, err := u.walletRepository.GetOneByUserId(ctx, userId)
	if err != nil {
		return nil, apperror.NewServerError(err)
	} else if wallet == nil {
		return nil, apperror.NewEntityNotFoundError(nil, "wallet")
	}

	return wallet, nil
}

func (u *walletUsecaseImpl) GetOneByNumber(ctx context.Context, walletNumber string) (*model.Wallet, error) {
	wallet, err := u.walletRepository.GetOneByNumber(ctx, walletNumber)
	if err != nil {
		return nil, apperror.NewServerError(err)
	} else if wallet == nil {
		return nil, apperror.NewEntityNotFoundError(nil, "wallet")
	}

	return wallet, nil
}

func (u *walletUsecaseImpl) GetOneByIdWithLock(ctx context.Context, walletId int64) (*model.Wallet, error) {
	wallet, err := u.walletRepository.GetOneByIdWithLock(ctx, walletId)
	if err != nil {
		return nil, apperror.NewServerError(err)
	} else if wallet == nil {
		return nil, apperror.NewEntityNotFoundError(nil, "wallet")
	}

	return wallet, nil
}

func (u *walletUsecaseImpl) GetOneByNumberWithLock(ctx context.Context, walletNumber string) (*model.Wallet, error) {
	wallet, err := u.walletRepository.GetOneByNumberWithLock(ctx, walletNumber)
	if err != nil {
		return nil, apperror.NewServerError(err)
	} else if wallet == nil {
		return nil, apperror.NewEntityNotFoundError(nil, "wallet")
	}

	return wallet, nil
}
