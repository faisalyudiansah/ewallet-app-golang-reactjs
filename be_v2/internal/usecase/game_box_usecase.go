package usecase

import (
	"context"
	"math/rand"

	"ewallet-server-v2/internal/constant"
	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/apperror"
	"ewallet-server-v2/internal/repository"
)

type GameBoxUsecase interface {
	GetAll(ctx context.Context) ([]model.GameBox, error)
	GetOneById(ctx context.Context, boxId int64) (*model.GameBox, error)
}

type gameBoxUsecaseImpl struct {
	gameBoxRepository repository.GameBoxRepository
}

func NewGameBoxUsecase(gameBoxRepository repository.GameBoxRepository) *gameBoxUsecaseImpl {
	return &gameBoxUsecaseImpl{
		gameBoxRepository: gameBoxRepository,
	}
}

func (u *gameBoxUsecaseImpl) GetAll(ctx context.Context) ([]model.GameBox, error) {
	res, err := u.gameBoxRepository.GetAll(ctx, constant.GameBoxLimit)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(res); i++ {
		num := rand.Intn(len(res))

		res[i], res[num] = res[num], res[i]
	}

	return res, nil
}

func (u *gameBoxUsecaseImpl) GetOneById(ctx context.Context, boxId int64) (*model.GameBox, error) {
	res, err := u.gameBoxRepository.GetOneById(ctx, boxId)
	if err != nil {
		return nil, apperror.NewServerError(err)
	} else if res == nil {
		return nil, apperror.NewEntityNotFoundError(nil, "game box")
	}

	return res, nil
}
