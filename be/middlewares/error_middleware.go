package middlewares

import (
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/apperrors"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) == 0 {
		return
	}
	if len(c.Errors) > 0 {
		var ve validator.ValidationErrors
		if errors.As(c.Errors[0].Err, &ve) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": helpers.FormatterErrorInput(ve)})
			return
		}

		errorMappings := map[error]int{
			apperrors.ErrISE:                            http.StatusInternalServerError,
			apperrors.ErrInvalidAccessToken:             http.StatusUnauthorized,
			apperrors.ErrUnauthorization:                http.StatusUnauthorized,
			apperrors.ErrUrlNotFound:                    http.StatusNotFound,
			apperrors.ErrRequestBodyInvalid:             http.StatusBadRequest,
			apperrors.ErrUserEmailAlreadyExists:         http.StatusBadRequest,
			apperrors.ErrUserFailedRegister:             http.StatusBadRequest,
			apperrors.ErrUserInvalidEmailPassword:       http.StatusBadRequest,
			apperrors.ErrUserEmailNotExists:             http.StatusBadRequest,
			apperrors.ErrUserTokenResetPasswordNotValid: http.StatusBadRequest,
			apperrors.ErrInvalidDateFormat:              http.StatusBadRequest,
			apperrors.ErrInvalidQueryLimit:              http.StatusBadRequest,
			apperrors.ErrInvalidQueryPage:               http.StatusBadRequest,
			apperrors.ErrSOFIdIsNotExists:               http.StatusBadRequest,
			apperrors.ErrForbiddenAccess:                http.StatusForbidden,
			apperrors.ErrGameZeroChance:                 http.StatusBadRequest,
			apperrors.ErrGameInvalidBoxIndex:            http.StatusBadRequest,
			apperrors.ErrWalletNumberIsNotExists:        http.StatusBadRequest,
			apperrors.ErrWalletTransferToTheirOwnWallet: http.StatusBadRequest,
			apperrors.ErrWalletBalanceIsInsufficient:    http.StatusBadRequest,
		}
		if statusCode, exists := errorMappings[c.Errors[0].Err]; exists {
			helpers.PrintError(c, statusCode, c.Errors[0].Err.Error())
			return
		}
		helpers.PrintError(c, http.StatusInternalServerError, apperrors.ErrISE.Error())
	}
}
