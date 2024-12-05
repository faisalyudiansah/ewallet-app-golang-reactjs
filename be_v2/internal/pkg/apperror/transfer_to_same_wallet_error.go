package apperror

import (
	"errors"

	"ewallet-server-v2/internal/constant"
)

func NewTransferToSameWalletError() AppError {
	msg := constant.TransferToSameWalletErrorMessage

	err := errors.New(msg)

	return NewAppError(err, DefaultClientErrorCode, msg, nil)
}
