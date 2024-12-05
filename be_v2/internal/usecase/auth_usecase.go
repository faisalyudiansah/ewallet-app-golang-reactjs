package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"ewallet-server-v2/internal/constant"
	"ewallet-server-v2/internal/dto/authdto"
	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/apperror"
	"ewallet-server-v2/internal/pkg/encryptutils"
	"ewallet-server-v2/internal/pkg/jwtutils"
	"ewallet-server-v2/internal/pkg/randutils"
	"ewallet-server-v2/internal/repository"
)

type AuthUsecase interface {
	Register(ctx context.Context, email string, password string, username string, fullName string) (*authdto.RegisterDto, error)
	Login(ctx context.Context, email string, password string) (*authdto.LoginDto, error)
	ResetPassword(ctx context.Context, email string) (*model.ResetPasswordAttempt, error)
	ConfirmResetPassword(ctx context.Context, email, password, code string) error
}

type authUsecaseImpl struct {
	jwtUtil                        jwtutils.JwtUtil
	randomUtil                     randutils.RandomUtil
	passwordEncryptor              encryptutils.PasswordEncryptor
	userUsecase                    UserUsecase
	walletUsecase                  WalletUsecase
	resetPasswordAttemptRepository repository.ResetPasswordAttemptRepository
}

func NewAuthUsecase(
	jwtUtil jwtutils.JwtUtil,
	randomUtil randutils.RandomUtil,
	passwordEncryptor encryptutils.PasswordEncryptor,
	userUsecase UserUsecase,
	walletUsecase WalletUsecase,
	resetPasswordAttemptRepository repository.ResetPasswordAttemptRepository,
) *authUsecaseImpl {
	return &authUsecaseImpl{
		jwtUtil:                        jwtUtil,
		randomUtil:                     randomUtil,
		passwordEncryptor:              passwordEncryptor,
		userUsecase:                    userUsecase,
		walletUsecase:                  walletUsecase,
		resetPasswordAttemptRepository: resetPasswordAttemptRepository,
	}
}

func (u *authUsecaseImpl) Register(ctx context.Context, email string, password string, username string, fullName string) (*authdto.RegisterDto, error) {
	userCheck, err := u.userUsecase.GetOneByEmail(ctx, email)
	if err != nil {
		// this error handling can be improved
		var appErr apperror.AppError
		switch {
		case errors.As(err, &appErr) && appErr.GetCode() != apperror.NotFoundErrorCode:
			return nil, err
		}
	} else if userCheck != nil {
		return nil, apperror.NewUserAlreadyRegisteredError()
	}

	hashed, err := u.passwordEncryptor.Hash(password)
	if err != nil {
		return nil, err
	}

	createdUser, err := u.userUsecase.CreateOne(ctx, email, hashed, username, fullName)
	if err != nil {
		return nil, err
	}

	createdWallet, err := u.walletUsecase.CreateOne(ctx, createdUser.UserId)
	if err != nil {
		return nil, err
	}

	return &authdto.RegisterDto{
		User:   *createdUser,
		Wallet: *createdWallet,
	}, nil
}

func (u *authUsecaseImpl) Login(ctx context.Context, email string, password string) (*authdto.LoginDto, error) {
	user, err := u.userUsecase.GetOneByEmail(ctx, email)
	if err != nil {
		// this error handling can be improved
		var appErr apperror.AppError
		switch {
		case errors.As(err, &appErr) && appErr.GetCode() == apperror.NotFoundErrorCode:
			return nil, apperror.NewInvalidLoginCredentials(err)
		}
		return nil, err
	}

	log.Println("start hashing")
	isValid := u.passwordEncryptor.Check(password, user.Password)
	if !isValid {
		return nil, apperror.NewInvalidLoginCredentials(nil)
	}
	log.Println("end hashing")

	token, err := u.jwtUtil.Sign(user.UserId)
	if err != nil {
		return nil, apperror.NewInvalidLoginCredentials(err)
	}

	return &authdto.LoginDto{
		Token: token,
	}, nil
}

func (u *authUsecaseImpl) ResetPassword(ctx context.Context, email string) (*model.ResetPasswordAttempt, error) {
	user, err := u.userUsecase.GetOneByEmail(ctx, email)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, apperror.NewResetPasswordError()
	}

	// removing all attempted reset password before creating a new one
	all, err := u.resetPasswordAttemptRepository.GetAllByUserId(ctx, user.UserId)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	attemptIds := []int64{}
	for _, m := range all {
		attemptIds = append(attemptIds, m.ResetPasswordAttemptId)
	}

	if len(attemptIds) > 0 {
		err = u.resetPasswordAttemptRepository.SoftDeleteByIds(ctx, attemptIds)
		if err != nil {
			return nil, apperror.NewServerError(err)
		}
	}

	resetCode := u.randomUtil.StringAlphaNum(constant.ResetPasswordCodeLength)
	newResetPass := model.ResetPasswordAttempt{
		Code:      resetCode,
		UserId:    user.UserId,
		ExpiredAt: time.Now().Add(constant.ResetPasswordValidDuration * time.Minute),
	}

	createdResetPass, err := u.resetPasswordAttemptRepository.CreateOne(ctx, newResetPass)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	return createdResetPass, nil
}

func (u *authUsecaseImpl) ConfirmResetPassword(ctx context.Context, email, password, code string) error {
	user, err := u.userUsecase.GetOneByEmail(ctx, email)
	if err != nil {
		return err
	}

	attempt, err := u.resetPasswordAttemptRepository.GetOneActiveByUserIdAndCode(ctx, user.UserId, code)
	if err != nil {
		return apperror.NewServerError(err)
	} else if attempt == nil {
		return apperror.NewResetPasswordConfirmError(nil)
	}

	user.Password, err = u.passwordEncryptor.Hash(password)
	if err != nil {
		return apperror.NewServerError(err)
	}

	_, err = u.userUsecase.SaveOne(ctx, *user)
	if err != nil {
		return apperror.NewServerError(err)
	}

	err = u.resetPasswordAttemptRepository.SoftDeleteByIds(ctx, []int64{
		attempt.ResetPasswordAttemptId,
	})
	if err != nil {
		return apperror.NewServerError(err)
	}

	return nil
}
