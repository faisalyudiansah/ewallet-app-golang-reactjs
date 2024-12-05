package apperrors

import (
	"errors"

	"ewallet-server-v1/constants"
)

var (
	ErrISE                = errors.New(constants.ISE)
	ErrInvalidAccessToken = errors.New(constants.InvalidAccessToken)
	ErrUnauthorization    = errors.New(constants.Unauthorization)
	ErrUrlNotFound        = errors.New(constants.UrlNotFound)
	ErrRequestBodyInvalid = errors.New(constants.RequestBodyInvalid)
	ErrInvalidDateFormat  = errors.New(constants.InvalidDateFormat)
	ErrInvalidQueryLimit  = errors.New(constants.InvalidQueryLimit)
	ErrInvalidQueryPage   = errors.New(constants.InvalidQueryPage)
	ErrForbiddenAccess    = errors.New(constants.ForbiddenAccess)
)

var (
	ErrUserEmailAlreadyExists         = errors.New(constants.UserEmailAlreadyExists)
	ErrUserEmailNotExists             = errors.New(constants.UserEmailNotExists)
	ErrUserFailedRegister             = errors.New(constants.UserFailedRegister)
	ErrUserInvalidEmailPassword       = errors.New(constants.UserInvalidEmailPassword)
	ErrUserTokenResetPasswordNotValid = errors.New(constants.UserTokenResetPasswordNotValid)
)

var (
	ErrSOFIdIsNotExists = errors.New(constants.SOFIdIsNotExists)
)

var (
	ErrGameZeroChance      = errors.New(constants.GameZeroChance)
	ErrGameInvalidBoxIndex = errors.New(constants.GameInvalidBoxIndex)
)

var (
	ErrWalletNumberIsNotExists        = errors.New(constants.WalletNumberIsNotExists)
	ErrWalletTransferToTheirOwnWallet = errors.New(constants.WalletTransferToTheirOwnWallet)
	ErrWalletBalanceIsInsufficient    = errors.New(constants.WalletBalanceIsInsufficient)
)
