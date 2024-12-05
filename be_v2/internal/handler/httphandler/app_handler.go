package httphandler

import (
	"net/http"

	"ewallet-server-v2/internal/dto/httpdto"
	"ewallet-server-v2/internal/pkg/ginutils"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (h *AppHandler) RouteNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, httpdto.ErrorResponse{
		Message: "Route not found",
	})
}

func (h *AppHandler) Index(c *gin.Context) {
	ginutils.ResponseOK(c, "Welcome to Sea Wallet API!", gin.H{
		"documentation": "See the documentation: https://documenter.getpostman.com/view/12104547/2sA3Bt2pXq",
	})
}
