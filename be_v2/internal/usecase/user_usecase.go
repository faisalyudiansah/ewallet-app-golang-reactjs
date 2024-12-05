package usecase

import (
	"context"

	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/apperror"
	"ewallet-server-v2/internal/repository"
)

type UserUsecase interface {
	CreateOne(ctx context.Context, email string, hashedPassword string, username string, fullName string) (*model.User, error)
	GetOneByEmail(ctx context.Context, email string) (*model.User, error)
	GetOneById(ctx context.Context, id int64) (*model.User, error)
	GetOneByWalletId(ctx context.Context, walletId int64) (*model.User, error)
	SaveOne(ctx context.Context, user model.User) (*model.User, error)
	UpdateOne(ctx context.Context, userId int64, user model.User) (*model.User, error)
}

type userUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUsecaseImpl(userRepository repository.UserRepository) *userUsecaseImpl {
	return &userUsecaseImpl{
		userRepository: userRepository,
	}
}

func (u *userUsecaseImpl) CreateOne(ctx context.Context, email string, hashedPassword string, username string, fullName string) (*model.User, error) {
	newUser := model.User{
		Email:    email,
		Password: hashedPassword,
		Name:     username,
		FullName: fullName,
	}

	createdUser, err := u.userRepository.CreateOne(ctx, newUser)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	return createdUser, nil
}

func (u *userUsecaseImpl) GetOneByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := u.userRepository.GetOneByEmail(ctx, email)
	if err != nil {
		return nil, apperror.NewServerError(err)
	} else if user == nil {
		return nil, apperror.NewEntityNotFoundError(err, "user")
	}

	return user, nil
}

func (u *userUsecaseImpl) GetOneById(ctx context.Context, userId int64) (*model.User, error) {
	user, err := u.userRepository.GetOneById(ctx, userId)
	if err != nil {
		return nil, apperror.NewServerError(err)
	} else if user == nil {
		return nil, apperror.NewEntityNotFoundError(err, "user")
	}

	return user, nil
}

func (u *userUsecaseImpl) GetOneByWalletId(ctx context.Context, walletId int64) (*model.User, error) {
	user, err := u.userRepository.GetOneByWalletId(ctx, walletId)
	if err != nil {
		return nil, apperror.NewServerError(err)
	} else if user == nil {
		return nil, apperror.NewEntityNotFoundError(err, "user")
	}

	return user, nil
}

func (u *userUsecaseImpl) SaveOne(ctx context.Context, user model.User) (*model.User, error) {
	savedUser, err := u.userRepository.SaveOne(ctx, user)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	return savedUser, nil
}

func (u *userUsecaseImpl) UpdateOne(ctx context.Context, userId int64, user model.User) (*model.User, error) {
	updatedUser, err := u.userRepository.UpdateOne(ctx, user)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	return updatedUser, nil
}
