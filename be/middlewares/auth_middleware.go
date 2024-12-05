package middlewares

import (
	"context"

	"strings"

	"ewallet-server-v1/apperrors"
	"ewallet-server-v1/constants"
	"ewallet-server-v1/helpers"

	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	reqToken := c.GetHeader("Authorization")
	if reqToken == "" || len(reqToken) == 0 {
		c.Error(apperrors.ErrUnauthorization)
		c.Abort()
		return
	}
	splitToken := strings.Split(reqToken, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		c.Error(apperrors.ErrUnauthorization)
		c.Abort()
		return
	}
	jwtProvider := helpers.NewJWTProviderHS256()
	result, err := jwtProvider.VerifyToken(splitToken[1])
	if err != nil {
		c.Error(apperrors.ErrUnauthorization)
		c.Abort()
		return
	}

	var userId constants.ID = "userId"
	ctxId := context.WithValue(c.Request.Context(), userId, result.UserID)
	c.Request = c.Request.WithContext(ctxId)

	c.Next()
}
