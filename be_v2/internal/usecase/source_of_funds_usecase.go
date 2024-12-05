package usecase

import (
	"context"

	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/apperror"
	"ewallet-server-v2/internal/repository"
)

type SourceOfFundsUsecase interface {
	GetOneById(ctx context.Context, sourceId int64) (*model.SourceOfFund, error)
}

type sourceOfFundsUsecaseImpl struct {
	sourceOfFundsRepository repository.SourceOfFundRepository
}

func NewSourceOfFundsUsecase(sourceOfFundsRepository repository.SourceOfFundRepository) *sourceOfFundsUsecaseImpl {
	return &sourceOfFundsUsecaseImpl{
		sourceOfFundsRepository: sourceOfFundsRepository,
	}
}

func (u *sourceOfFundsUsecaseImpl) GetOneById(ctx context.Context, sourceId int64) (*model.SourceOfFund, error) {
	res, err := u.sourceOfFundsRepository.GetOneById(ctx, sourceId)
	if err != nil {
		return nil, apperror.NewServerError(err)
	} else if res == nil {
		return nil, apperror.NewEntityNotFoundError(nil, "source of funds")
	}

	return res, nil
}
