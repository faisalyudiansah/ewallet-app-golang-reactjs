package services

import (
	"context"

	"ewallet-server-v1/apperrors"
	"ewallet-server-v1/dtos"
	"ewallet-server-v1/helpers"
	"ewallet-server-v1/models"
	"ewallet-server-v1/repositories"

	"github.com/shopspring/decimal"
)

type WalletServiceInterface interface {
	PutTopWalletService(context.Context, dtos.RequestTopUpWallet, int64) (*dtos.ResponseUserAndWallet, error)
	PostTransferFundService(context.Context, dtos.RequestTransferFund, int64) (*dtos.ResponseSuccessTransfer, error)
}

type WalletServiceImplementation struct {
	UserRepository            repositories.UserRepository
	WalletRepository          repositories.WalletRepository
	TransactionUserRepository repositories.TransactionUserRepository
	TransactionsRepository    repositories.TransactionRepository
	SourceFundsRepository     repositories.SourceOfFundRepository
}

func NewWalletServiceImplementation(
	us repositories.UserRepository,
	w repositories.WalletRepository,
	txu repositories.TransactionUserRepository,
	tx repositories.TransactionRepository,
	sof repositories.SourceOfFundRepository,
) *WalletServiceImplementation {
	return &WalletServiceImplementation{
		UserRepository:            us,
		WalletRepository:          w,
		TransactionUserRepository: txu,
		TransactionsRepository:    tx,
		SourceFundsRepository:     sof,
	}
}

func (ws *WalletServiceImplementation) PutTopWalletService(ctx context.Context, reqBody dtos.RequestTopUpWallet, userId int64) (*dtos.ResponseUserAndWallet, error) {
	result, err := ws.TransactionsRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		if reqBody.SourceOfFund == 4 {
			return nil, apperrors.ErrForbiddenAccess
		}
		var user *models.User
		user, err := ws.UserRepository.GetUserById(cForTx, userId)
		if err != nil || user.ID == 0 {
			return nil, apperrors.ErrInvalidAccessToken
		}
		getSof, err := ws.SourceFundsRepository.GetSourceOfFundById(cForTx, reqBody.SourceOfFund)
		if err != nil || getSof.ID == 0 {
			return nil, apperrors.ErrSOFIdIsNotExists
		}
		wallet, err := ws.WalletRepository.GetWalletByIdUser(cForTx, user.ID)
		if err != nil || wallet.ID == 0 {
			return nil, apperrors.ErrInvalidAccessToken
		}
		wallet, err = ws.WalletRepository.PutChangeBalanceWallet(cForTx, (wallet.Balance.InexactFloat64() + reqBody.Amount), user.ID)
		if err != nil {
			return nil, apperrors.ErrISE
		}
		tx, err := ws.TransactionUserRepository.PostNewTransaction(cForTx, user.ID, getSof.ID, reqBody.Amount, getSof.Name, "")
		if err != nil {
			return nil, apperrors.ErrISE
		}
		err = ws.TransactionUserRepository.PostTransactionUserPivot(cForTx, user.ID, tx.ID)
		if err != nil {
			return nil, apperrors.ErrISE
		}
		if reqBody.Amount == 10000000 {
			user, err = ws.UserRepository.PutAttemptGame(cForTx, (user.ChanceGame + 1), user.ID)
			if err != nil {
				return nil, apperrors.ErrISE
			}
		}
		dataUser := &models.UserAndWallet{
			User:   *user,
			Wallet: *wallet,
		}
		return dataUser, nil
	})
	if err != nil {
		return nil, err
	}
	return helpers.ToResponseUserAndWallet(result.(*models.UserAndWallet)), nil
}

func (ws *WalletServiceImplementation) PostTransferFundService(ctx context.Context, reqBody dtos.RequestTransferFund, userId int64) (*dtos.ResponseSuccessTransfer, error) {
	result, err := ws.TransactionsRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		user, err := ws.UserRepository.GetUserById(cForTx, userId)
		if err != nil || user.ID == 0 {
			return nil, apperrors.ErrInvalidAccessToken
		}
		wallet, err := ws.WalletRepository.GetWalletByIdUser(cForTx, user.ID)
		if err != nil || wallet.ID == 0 {
			return nil, apperrors.ErrInvalidAccessToken
		}
		if wallet.WalletNumber == reqBody.ToWalletNumber {
			return nil, apperrors.ErrWalletTransferToTheirOwnWallet
		}
		if wallet.Balance.InexactFloat64() < reqBody.Amount {
			return nil, apperrors.ErrWalletBalanceIsInsufficient
		}
		targetWallet, err := ws.WalletRepository.GetWalletByWalletNumber(cForTx, reqBody.ToWalletNumber)
		if err != nil || targetWallet.ID == 0 {
			return nil, apperrors.ErrWalletNumberIsNotExists
		}
		result, err := ws.WalletRepository.PutChangeBalanceWallet(cForTx, (targetWallet.Balance.InexactFloat64() + reqBody.Amount), targetWallet.UserId)
		if err != nil {
			return nil, apperrors.ErrISE
		}
		tx, err := ws.TransactionUserRepository.PostNewTransaction(cForTx, result.UserId, 1, reqBody.Amount, "Bank Transfer", reqBody.Description)
		if err != nil {
			return nil, apperrors.ErrISE
		}
		err = ws.TransactionUserRepository.PostTransactionUserPivot(cForTx, user.ID, tx.ID)
		if err != nil {
			return nil, apperrors.ErrISE
		}
		res := &dtos.ResponseSuccessTransfer{
			UserId:           userId,
			RecipientId:      targetWallet.UserId,
			Amount:           decimal.NewFromFloat(reqBody.Amount),
			RemainingBelance: result.Balance,
			TransactionTime:  tx.TransactionTime,
			Description:      tx.Description,
		}
		return res, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*dtos.ResponseSuccessTransfer), nil
}
