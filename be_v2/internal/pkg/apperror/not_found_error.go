package apperror

import (
	"errors"
	"fmt"

	"ewallet-server-v2/internal/constant"
)

func NewEntityNotFoundError(err error, entityName string) AppError {
	msg := fmt.Sprintf(constant.EntityNotFoundErrorMessage, entityName)

	if err == nil {
		err = errors.New(msg)
	}

	return NewAppError(err, NotFoundErrorCode, msg, nil)
}
