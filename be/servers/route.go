package servers

import (
	"ewallet-server-v1/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoute(h *HandlerOps) *gin.Engine {
	g := gin.New()
	g.ContextWithFallback = true
	g.Use(gin.Recovery(), middlewares.LoggerMiddleware(), middlewares.ErrorHandler)

	g.NoRoute(InvalidRoute)

	SetupAuthenRoutes(g, h)
	SetupTransactionRoutes(g, h)
	SetupGameRoute(g, h)

	return g
}

func SetupAuthenRoutes(g *gin.Engine, h *HandlerOps) {
	u := g.Group("/user")
	u.POST("/register", h.UserController.PostRegisterUserController)
	u.POST("/login", h.UserController.PostLoginUserController)
	u.POST("/forget-password", h.ResetPasswordController.PostForgetPasswordController)
	u.PUT("/reset-password/:token", middlewares.Authorization, h.ResetPasswordController.PutResetPasswordController)
	u.PUT("/top-up/wallet", middlewares.Authorization, h.WalletController.PutTopUpWalletController)
	u.GET("/me", middlewares.Authorization, h.UserController.GetUserDetail)
}

func SetupTransactionRoutes(g *gin.Engine, h *HandlerOps) {
	u := g.Group("/transactions")
	u.GET("", middlewares.Authorization, h.TransactionUserController.GetListTransactionsController)
	u.POST("/transfers", middlewares.Authorization, h.WalletController.PostTransferFundController)
}

func SetupGameRoute(g *gin.Engine, h *HandlerOps) {
	u := g.Group("/game")
	u.POST("/start", middlewares.Authorization, h.GameController.PostStartGameController)
	u.POST("/choose", middlewares.Authorization, h.GameController.PostChooseGachaBoxController)
}
