package apperror

import "ewallet-server-v2/internal/constant"

func NewServerError(err error) AppError {
	msg := constant.InternalServerErrorMessage

	return NewAppError(err, DefaultServerErrorCode, msg, nil)
}
