package services

import (
	"context"
	"math/rand"
	"time"

	"ewallet-server-v1/apperrors"
	"ewallet-server-v1/repositories"
)

type GameServiceInterface interface {
	PostStartGameService(context.Context, int64) ([]int, error)
	PostChooseGachaBoxService(context.Context, int64, int) (float64, error)
}

type GameServiceImplementation struct {
	GameRepository            repositories.GameRepositoryInterface
	TransactionUserRepository repositories.TransactionUserRepository
	TransactionsRepository    repositories.TransactionRepository
	UserRepository            repositories.UserRepository
}

func NewGameServiceImplementation(
	gr repositories.GameRepositoryInterface,
	txu repositories.TransactionUserRepository,
	tx repositories.TransactionRepository,
	us repositories.UserRepository,
) *GameServiceImplementation {
	return &GameServiceImplementation{
		GameRepository:            gr,
		TransactionUserRepository: txu,
		TransactionsRepository:    tx,
		UserRepository:            us,
	}
}

func (gs *GameServiceImplementation) PostStartGameService(ctx context.Context, userId int64) ([]int, error) {
	result, err := gs.TransactionsRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		user, err := gs.UserRepository.GetUserById(cForTx, userId)
		if err != nil || user.ID == 0 {
			return nil, apperrors.ErrInvalidAccessToken
		}
		if user.ChanceGame == 0 {
			return nil, apperrors.ErrGameZeroChance
		}
		rewards := make([]int, 9)
		for i := range rewards {
			rand.Seed(time.Now().UnixNano())
			rewards[i] = rand.Intn(1000) * 1000
		}
		err = gs.GameRepository.CreateGameBox(cForTx, userId, rewards)
		if err != nil {
			return nil, apperrors.ErrISE
		}
		dataUser, err := gs.UserRepository.PutAttemptGame(cForTx, (user.ChanceGame - 1), user.ID)
		if err != nil || dataUser.ID == 0 {
			return nil, apperrors.ErrISE
		}
		var rewardsShow = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		return rewardsShow, nil
	})
	if err != nil {
		return nil, err
	}
	return result.([]int), nil
}

func (gs *GameServiceImplementation) PostChooseGachaBoxService(ctx context.Context, userId int64, boxIndex int) (float64, error) {
	result, err := gs.TransactionsRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		user, err := gs.UserRepository.GetUserById(cForTx, userId)
		if err != nil || user.ID == 0 {
			return nil, apperrors.ErrInvalidAccessToken
		}
		gameBox, err := gs.GameRepository.GetGameBox(cForTx, userId)
		if err != nil {
			return 0, apperrors.ErrGameZeroChance
		}
		rewards := []float64{gameBox.Box1, gameBox.Box2, gameBox.Box3, gameBox.Box4, gameBox.Box5, gameBox.Box6, gameBox.Box7, gameBox.Box8, gameBox.Box9}
		if boxIndex < 1 || boxIndex > len(rewards) {
			return nil, apperrors.ErrGameInvalidBoxIndex
		}
		selectedReward := rewards[boxIndex-1]
		err = gs.GameRepository.ChangeStatusBox(cForTx, gameBox.ID)
		if err != nil {
			return 0, apperrors.ErrISE
		}
		tx, err := gs.TransactionUserRepository.PostNewTransaction(cForTx, user.ID, 4, selectedReward, "Reward", "")
		if err != nil {
			return nil, apperrors.ErrISE
		}
		err = gs.TransactionUserRepository.PostTransactionUserPivot(cForTx, user.ID, tx.ID)
		if err != nil {
			return nil, apperrors.ErrISE
		}
		return selectedReward, nil
	})
	if err != nil {
		return 0, err
	}
	return result.(float64), nil
}
