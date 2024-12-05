package usecase

import (
	"context"

	"ewallet-server-v2/internal/constant"
	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/apperror"
	"ewallet-server-v2/internal/repository"

	"github.com/shopspring/decimal"
)

type GameAttemptUsecase interface {
	CalculateChances(ctx context.Context, walletId int64) (int64, error)
	Attempt(ctx context.Context, walletId int64, boxId int64) (*model.GameAttempt, error)
}

type gameAttemptUsecaseImpl struct {
	transactionUsecase    TransactionUsecase
	gameAttemptRepository repository.GameAttemptRepository
	gameBoxUsecase        GameBoxUsecase
	walletUsecase         WalletUsecase
}

func NewGameAttemptUsecase(
	transactionUsecase TransactionUsecase,
	gameAttemptRepository repository.GameAttemptRepository,
	gameBoxUsecase GameBoxUsecase,
	walletUsecase WalletUsecase,
) *gameAttemptUsecaseImpl {
	return &gameAttemptUsecaseImpl{
		transactionUsecase:    transactionUsecase,
		gameAttemptRepository: gameAttemptRepository,
		gameBoxUsecase:        gameBoxUsecase,
		walletUsecase:         walletUsecase,
	}
}

func (u *gameAttemptUsecaseImpl) CalculateChances(ctx context.Context, walletId int64) (int64, error) {
	sourceOfFundsToQuery := []int64{
		constant.SourceOfFundBankTransfer,
		constant.SourceOfFundCreditCard,
		constant.SourceOfFundCash,
	}

	topUpMultiply := decimal.NewFromInt(constant.TopupAmountToParticipate)

	totalAttempts, err := u.transactionUsecase.GetTransactionTypeSumMultiple(ctx, walletId, constant.TransactionTypeTopUp, sourceOfFundsToQuery, topUpMultiply)
	if err != nil {
		return 0, err
	}

	attemptCount, err := u.gameAttemptRepository.GetCountByWalletId(ctx, walletId)
	if err != nil {
		return 0, apperror.NewServerError(err)
	}

	return totalAttempts.IntPart() - attemptCount, nil
}

func (u *gameAttemptUsecaseImpl) Attempt(ctx context.Context, walletId int64, boxId int64) (*model.GameAttempt, error) {
	chances, err := u.CalculateChances(ctx, walletId)
	if err != nil {
		return nil, err
	}

	if chances < 1 {
		return nil, apperror.NewInsufficientTopUpChanceError()
	}

	gameBox, err := u.gameBoxUsecase.GetOneById(ctx, boxId)
	if err != nil {
		return nil, err
	}

	newAttempt := model.GameAttempt{
		WalletId:    walletId,
		GameBoxesid: boxId,
		Amount:      gameBox.Amount,
	}

	createdAttempt, err := u.gameAttemptRepository.CreateOne(ctx, newAttempt)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	_, err = u.transactionUsecase.Topup(ctx, walletId, gameBox.Amount, constant.SourceOfFundReward)
	if err != nil {
		return nil, err
	}

	return createdAttempt, nil
}
