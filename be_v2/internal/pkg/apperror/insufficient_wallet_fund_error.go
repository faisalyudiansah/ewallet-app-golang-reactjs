package apperror

import (
	"errors"

	"ewallet-server-v2/internal/constant"
)

func NewInsufficientWalletFundError() AppError {
	msg := constant.InsufficientWalletFundErrorMessage

	err := errors.New(msg)

	return NewAppError(err, DefaultClientErrorCode, msg, nil)
}
