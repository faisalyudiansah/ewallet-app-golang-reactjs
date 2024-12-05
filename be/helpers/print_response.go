package helpers

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/dtos"
	"github.com/gin-gonic/gin"
)

func PrintError(c *gin.Context, statusCode int, msg string) {
	res := dtos.ResponseMessageOnly{
		Message: msg,
	}
	c.AbortWithStatusJSON(statusCode, res)
}

func PrintResponse(c *gin.Context, statusCode int, res interface{}) {
	c.JSON(statusCode, res)
}
