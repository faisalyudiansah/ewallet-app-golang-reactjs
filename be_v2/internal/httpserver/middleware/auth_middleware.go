package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"ewallet-server-v2/internal/dto/httpdto"
	"ewallet-server-v2/internal/pkg/ginutils"
	"ewallet-server-v2/internal/pkg/jwtutils"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtUtil jwtutils.JwtUtil
}

func NewAuthMiddleware(jwtUtil jwtutils.JwtUtil) *AuthMiddleware {
	return &AuthMiddleware{
		jwtUtil: jwtUtil,
	}
}

func (m *AuthMiddleware) RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authStr := c.Request.Header.Get("Authorization")
		if authStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httpdto.ErrorResponse{
				Message: "please provide the token",
			})
			return
		}

		authStrs := strings.Split(authStr, " ")
		if len(authStrs) != 2 || authStrs[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httpdto.ErrorResponse{
				Message: "invalid token - malformed",
			})
			return
		}

		tokenString := authStrs[1]

		claims, err := m.jwtUtil.Parse(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httpdto.ErrorResponse{
				Message: fmt.Sprintf("invalid token - %s", err),
			})
		}

		ginutils.SetUserId(c, claims.UserId)
	}
}
