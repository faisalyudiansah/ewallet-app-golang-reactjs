package apperror

import (
	"errors"

	"ewallet-server-v2/internal/constant"
)

func NewResetPasswordError() AppError {
	msg := constant.ResetPasswordErrorMessage

	err := errors.New(msg)

	return NewAppError(err, DefaultClientErrorCode, msg, nil)
}
