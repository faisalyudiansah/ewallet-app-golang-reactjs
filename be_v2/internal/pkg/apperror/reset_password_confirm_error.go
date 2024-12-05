package apperror

import (
	"errors"

	"ewallet-server-v2/internal/constant"
)

func NewResetPasswordConfirmError(err error) AppError {
	msg := constant.InvalidResetPasswordCodeErrorMessage

	if err == nil {
		err = errors.New(msg)
	}

	return NewAppError(err, DefaultClientErrorCode, msg, nil)
}
