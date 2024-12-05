package apperror

import (
	"errors"

	"ewallet-server-v2/internal/constant"
)

func NewInsufficientTopUpChanceError() AppError {
	msg := constant.InsufficientTopUpChanceErrorMessage

	err := errors.New(msg)

	return NewAppError(err, DefaultClientErrorCode, msg, nil)
}
