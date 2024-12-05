package services

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/apperrors"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/dtos"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/helpers"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/repositories"
)

type UserServiceInterface interface {
	PostRegisterUserService(context.Context, dtos.RequestRegisterUser) (*dtos.ResponseUserAndWallet, error)
	PostLoginUserService(context.Context, dtos.RequestLoginUser) (*dtos.ResponseAccessToken, error)
	GetUserDetailService(context.Context, int64) (*dtos.ResponseUserAndWallet, error)
}

type UserServiceImplementation struct {
	UserRepository         repositories.UserRepository
	WalletRepository       repositories.WalletRepository
	TransactionsRepository repositories.TransactionRepository
	Bcrypt                 helpers.Bcrypt
	Jwt                    helpers.JWTProvider
	GenerateNumber         helpers.GenerateNumberInterface
}

func NewUserServiceImplementation(
	us repositories.UserRepository,
	w repositories.WalletRepository,
	tx repositories.TransactionRepository,
	bcr helpers.Bcrypt,
	jwt helpers.JWTProvider,
	gn helpers.GenerateNumberInterface,
) *UserServiceImplementation {
	return &UserServiceImplementation{
		UserRepository:         us,
		WalletRepository:       w,
		TransactionsRepository: tx,
		Bcrypt:                 bcr,
		Jwt:                    jwt,
		GenerateNumber:         gn,
	}
}

func (us *UserServiceImplementation) PostRegisterUserService(ctx context.Context, reqBody dtos.RequestRegisterUser) (*dtos.ResponseUserAndWallet, error) {
	result, err := us.TransactionsRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		if IsEmailAlreadyRegistered := us.UserRepository.IsEmailAlreadyRegistered(cForTx, reqBody.Email); IsEmailAlreadyRegistered {
			return nil, apperrors.ErrUserEmailAlreadyExists
		}
		hashPassword, err := us.Bcrypt.HashPassword(reqBody.Password, 10)
		if err != nil {
			return nil, apperrors.ErrUserFailedRegister
		}
		user, err := us.UserRepository.PostUser(cForTx, reqBody, string(hashPassword))
		if err != nil {
			return nil, apperrors.ErrUserFailedRegister
		}
		wallet, err := us.WalletRepository.PostCreateWalletUser(cForTx, user.ID)
		if err != nil {
			return nil, apperrors.ErrUserFailedRegister
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

func (us *UserServiceImplementation) PostLoginUserService(ctx context.Context, reqBody dtos.RequestLoginUser) (*dtos.ResponseAccessToken, error) {
	result, err := us.TransactionsRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		user, err := us.UserRepository.GetUserByEmail(cForTx, reqBody.Email)
		if err != nil {
			return nil, apperrors.ErrUserInvalidEmailPassword
		}
		isValid, err := us.Bcrypt.CheckPassword(reqBody.Password, []byte(user.Password))
		if err != nil || !isValid {
			return nil, apperrors.ErrUserInvalidEmailPassword
		}
		accessToken, err := us.Jwt.CreateToken(int64(user.ID))
		if err != nil {
			return nil, apperrors.ErrISE
		}
		return accessToken, nil
	})
	if err != nil {
		return nil, err
	}
	return helpers.ToResponseAccessToken(result.(string)), nil
}

func (us *UserServiceImplementation) GetUserDetailService(ctx context.Context, userId int64) (*dtos.ResponseUserAndWallet, error) {
	result, err := us.TransactionsRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		user, err := us.UserRepository.GetUserById(cForTx, userId)
		if err != nil {
			return nil, apperrors.ErrInvalidAccessToken
		}
		wallet, err := us.WalletRepository.GetWalletByIdUser(cForTx, user.ID)
		if err != nil || wallet.ID == 0 {
			return nil, apperrors.ErrInvalidAccessToken
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
