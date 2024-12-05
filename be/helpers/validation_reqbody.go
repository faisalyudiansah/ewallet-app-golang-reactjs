package helpers

import (
	"io"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/apperrors"
	"github.com/gin-gonic/gin"
)

type ValidationReqBodyInterface interface {
	ValidationReqBody(*gin.Context, interface{}) error
}

type ValidationReqBody struct{}

func NewValidationReqBody() *ValidationReqBody {
	return &ValidationReqBody{}
}

func (vr *ValidationReqBody) ValidationReqBody(c *gin.Context, reqBody interface{}) error {
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		if err == io.EOF {
			return apperrors.ErrRequestBodyInvalid
		}
		return err
	}
	return nil
}
