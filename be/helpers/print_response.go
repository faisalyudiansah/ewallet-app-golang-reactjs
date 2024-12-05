package helpers

import (
	"ewallet-server-v1/dtos"

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
