package servers

import (
	"ewallet-server-v1/apperrors"

	"github.com/gin-gonic/gin"
)

func InvalidRoute(c *gin.Context) {
	c.Error(apperrors.ErrUrlNotFound)
	c.Abort()
}
