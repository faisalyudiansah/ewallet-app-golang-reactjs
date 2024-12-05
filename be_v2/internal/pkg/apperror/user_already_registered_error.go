package apperror

import (
	"errors"

	"ewallet-server-v2/internal/constant"
)

func NewUserAlreadyRegisteredError() AppError {
	msg := constant.UserAlreadyRegisteredError

	err := errors.New(msg)

	return NewAppError(err, DefaultClientErrorCode, msg, nil)
}
