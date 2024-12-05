package services

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/apperrors"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/dtos"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/helpers"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/repositories"
)

type ResetPasswordInterface interface {
	PostForgetPasswordService(context.Context, dtos.RequestForgetPassword) (*dtos.ResponseTokenResetPassword, error)
	PutResetPasswordService(context.Context, string, dtos.RequestResetPassword, int64) (*dtos.ResponseUser, error)
}

type ResetPasswordServiceImplementation struct {
	ResetPasswordRepository repositories.ResetPasswordRepository
	UserRepository          repositories.UserRepository
	TransactionsRepository  repositories.TransactionRepository
	GenerateNumber          helpers.GenerateNumberInterface
	Bcrypt                  helpers.Bcrypt
}

func NewResetPasswordServiceImplementation(
	rp repositories.ResetPasswordRepository,
	us repositories.UserRepository,
	tx repositories.TransactionRepository,
	gn helpers.GenerateNumberInterface,
	bc helpers.Bcrypt,
) *ResetPasswordServiceImplementation {
	return &ResetPasswordServiceImplementation{
		ResetPasswordRepository: rp,
		UserRepository:          us,
		TransactionsRepository:  tx,
		GenerateNumber:          gn,
		Bcrypt:                  bc,
	}
}

func (rp *ResetPasswordServiceImplementation) PostForgetPasswordService(ctx context.Context, reqBody dtos.RequestForgetPassword) (*dtos.ResponseTokenResetPassword, error) {
	result, err := rp.TransactionsRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		user, err := rp.UserRepository.GetUserByEmail(cForTx, reqBody.Email)
		if err != nil {
			return nil, apperrors.ErrUserEmailNotExists
		}
		token := rp.GenerateNumber.GenerateTokenResetPassword(user.ID)
		result, err := rp.ResetPasswordRepository.PostNewDataResetPassword(cForTx, user.ID, token)
		if err != nil {
			return nil, apperrors.ErrISE
		}
		return result.Token, nil
	})
	if err != nil {
		return nil, err
	}
	return helpers.ToResponseTokenResetPassword(result.(string)), nil
}

func (rp *ResetPasswordServiceImplementation) PutResetPasswordService(ctx context.Context, token string, reqBody dtos.RequestResetPassword, userId int64) (*dtos.ResponseUser, error) {
	result, err := rp.TransactionsRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		if isValid := rp.ResetPasswordRepository.IsTokenResetValid(cForTx, token); !isValid {
			return nil, apperrors.ErrUserTokenResetPasswordNotValid
		}
		user, err := rp.UserRepository.GetUserByEmail(cForTx, reqBody.Email)
		if err != nil {
			return nil, apperrors.ErrUserEmailNotExists
		}
		if user.ID != userId {
			return nil, apperrors.ErrUnauthorization
		}
		hashPassword, err := rp.Bcrypt.HashPassword(reqBody.NewPassword, 10)
		if err != nil {
			return nil, apperrors.ErrUserFailedRegister
		}
		result, err := rp.UserRepository.PutResetPassword(cForTx, string(hashPassword), userId)
		if err != nil {
			return nil, apperrors.ErrISE
		}
		err = rp.ResetPasswordRepository.DeleteToken(cForTx, token)
		if err != nil {
			return nil, apperrors.ErrISE
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}
	return helpers.ToResponseUser(result.(*models.User)), nil
}
