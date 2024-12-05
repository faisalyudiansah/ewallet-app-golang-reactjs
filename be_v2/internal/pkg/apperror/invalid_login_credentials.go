package apperror

import (
	"errors"

	"ewallet-server-v2/internal/constant"
)

func NewInvalidLoginCredentials(err error) AppError {
	msg := constant.InvalidLoginCredentialsErrorMessage

	if err == nil {
		err = errors.New(msg)
	}

	return NewAppError(err, DefaultClientErrorCode, msg, nil)
}
