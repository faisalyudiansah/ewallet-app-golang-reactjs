package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetParamInterface interface {
	GetParamNumber(*gin.Context, string) (int, error)
}

type GetParam struct{}

func NewGetParams() *GetParam {
	return &GetParam{}
}

func (gp *GetParam) GetParamNumber(c *gin.Context, key string) (int, error) {
	paramNumber, err := strconv.Atoi(c.Param(key))
	return paramNumber, err
}
