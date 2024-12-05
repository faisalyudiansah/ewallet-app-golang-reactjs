package usecase

import (
	"context"
	"fmt"
	"time"

	"ewallet-server-v2/internal/constant"
	"ewallet-server-v2/internal/dto/appdto"
	"ewallet-server-v2/internal/dto/pagedto"
	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/apperror"
	"ewallet-server-v2/internal/repository"

	"github.com/shopspring/decimal"
)

type TransactionUsecase interface {
	GetListByWalletId(ctx context.Context, walletId int64, pageDto pagedto.PageSortDto) (*appdto.TransactionListDto, error)
	GetTransactionTypeSumMultiple(ctx context.Context, walletId, transactionTypeId int64, transactionAdditionalDetailId []int64, minAmount decimal.Decimal) (*decimal.Decimal, error)
	GetTransactionType(ctx context.Context) ([]model.TransactionType, error)
	Transfer(ctx context.Context, fromId int64, to string, amount decimal.Decimal, description string) (*model.Transaction, error)
	Topup(ctx context.Context, walletId int64, amount decimal.Decimal, sourceOfFundId int64) (*model.Transaction, error)
	GetThisMonthExpenseSum(ctx context.Context, walletId int64, date time.Time) (*appdto.TransactionSum, error)
	GetExpenseSumByMonth(ctx context.Context, walletId int64, month time.Month) (*appdto.TransactionSum, error)
}

type transactionUsecaseImpl struct {
	transactionRepository repository.TransactionRepository
	walletUsecase         WalletUsecase
	sourceOfFundsUsecase  SourceOfFundsUsecase
}

func NewTransactionUsecase(
	transactionRepository repository.TransactionRepository,
	walletUsecase WalletUsecase,
	sourceOfFundsUsecase SourceOfFundsUsecase,
) *transactionUsecaseImpl {
	return &transactionUsecaseImpl{
		transactionRepository: transactionRepository,
		walletUsecase:         walletUsecase,
		sourceOfFundsUsecase:  sourceOfFundsUsecase,
	}
}

func (u *transactionUsecaseImpl) GetListByWalletId(ctx context.Context, walletId int64, pageDto pagedto.PageSortDto) (*appdto.TransactionListDto, error) {
	res, err := u.transactionRepository.GetListByWalletId(ctx, walletId, pageDto)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	return res, nil
}

func (u *transactionUsecaseImpl) GetTransactionTypeSumMultiple(
	ctx context.Context,
	walletId int64,
	transactionTypeId int64,
	transactionAdditionalDetailId []int64,
	minAmount decimal.Decimal,
) (*decimal.Decimal, error) {
	sum, err := u.transactionRepository.GetTransactionTypeSumMultiple(ctx, walletId, transactionTypeId, transactionAdditionalDetailId, minAmount)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	return sum, nil
}

func (u *transactionUsecaseImpl) GetTransactionType(ctx context.Context) ([]model.TransactionType, error) {
	res, err := u.transactionRepository.GetTransactionType(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *transactionUsecaseImpl) Transfer(ctx context.Context, fromId int64, to string, amount decimal.Decimal, description string) (*model.Transaction, error) {
	toWallet, err := u.walletUsecase.GetOneByNumberWithLock(ctx, to)
	if err != nil {
		return nil, err
	}

	if fromId == toWallet.WalletId {
		return nil, apperror.NewTransferToSameWalletError()
	}

	fromWallet, err := u.walletUsecase.GetOneByIdWithLock(ctx, fromId)
	if err != nil {
		return nil, err
	}

	if fromWallet.Amount.LessThan(amount) {
		return nil, apperror.NewInsufficientWalletFundError()
	}

	newFromWalletAmount := fromWallet.Amount.Sub(amount)
	_, err = u.walletUsecase.UpdateOneAmountById(ctx, fromId, newFromWalletAmount)
	if err != nil {
		return nil, err
	}

	newToWalletAmount := toWallet.Amount.Add(amount)
	_, err = u.walletUsecase.UpdateOneAmountById(ctx, toWallet.WalletId, newToWalletAmount)
	if err != nil {
		return nil, err
	}

	fromTransaction := model.Transaction{
		WalletId:                      fromId,
		TransactionRefId:              nil,
		TransactionAdditionalDetailId: toWallet.WalletId,
		TransactionTypeId:             constant.TransactionTypeTransfer,
		Amount:                        amount.Neg(),
		Description:                   description,
	}

	createdFromTransaction, err := u.transactionRepository.CreateOne(ctx, fromTransaction)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	toTransaction := model.Transaction{
		WalletId:                      toWallet.WalletId,
		TransactionRefId:              &createdFromTransaction.TransactionId,
		TransactionAdditionalDetailId: fromId,
		TransactionTypeId:             constant.TransactionTypeTransfer,
		Amount:                        amount,
		Description:                   description,
	}

	_, err = u.transactionRepository.CreateOne(ctx, toTransaction)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	return createdFromTransaction, nil
}

func (u *transactionUsecaseImpl) Topup(ctx context.Context, walletId int64, amount decimal.Decimal, sourceOfFundId int64) (*model.Transaction, error) {
	_, err := u.sourceOfFundsUsecase.GetOneById(ctx, sourceOfFundId)
	if err != nil {
		return nil, err
	}

	wallet, err := u.walletUsecase.GetOneByIdWithLock(ctx, walletId)
	if err != nil {
		return nil, err
	}

	newAmount := wallet.Amount.Add(amount)

	_, err = u.walletUsecase.UpdateOneAmountById(ctx, walletId, newAmount)
	if err != nil {
		return nil, err
	}

	newTransaction := model.Transaction{
		WalletId:                      walletId,
		TransactionRefId:              nil,
		Amount:                        amount,
		TransactionTypeId:             constant.TransactionTypeTopUp,
		TransactionAdditionalDetailId: sourceOfFundId,
		Description:                   fmt.Sprintf(constant.MessageTopUpDescription, constant.ConvertSourceOfFundToReadable(sourceOfFundId)),
	}

	createdTransaction, err := u.transactionRepository.CreateOne(ctx, newTransaction)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	return createdTransaction, nil
}

func (u *transactionUsecaseImpl) GetThisMonthExpenseSum(ctx context.Context, walletId int64, date time.Time) (*appdto.TransactionSum, error) {
	res, err := u.transactionRepository.GetThisMonthTransactionSum(ctx, walletId, date, date)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *transactionUsecaseImpl) GetExpenseSumByMonth(ctx context.Context, walletId int64, month time.Month) (*appdto.TransactionSum, error) {
	currentTime := time.Now()
	currentYear := currentTime.Year()

	startOfMonth := time.Date(currentYear, month, 1, 0, 0, 0, 0, currentTime.Location())

	res, err := u.transactionRepository.GetThisMonthTransactionSum(ctx, walletId, startOfMonth, startOfMonth)
	if err != nil {
		return nil, err
	}

	return res, nil
}
