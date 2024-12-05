package ginutils

import (
	"ewallet-server-v2/internal/constant"

	"github.com/gin-gonic/gin"
)

func SetUserId(c *gin.Context, userId int64) error {
	c.Set(constant.ContextUserId, userId)

	return nil
}
