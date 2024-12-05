package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"
	"time"

	"ewallet-server-v2/internal/config"
	"ewallet-server-v2/internal/handler/httphandler"
	"ewallet-server-v2/internal/httpserver/middleware"
	"ewallet-server-v2/internal/pkg/apputils"
	"ewallet-server-v2/internal/pkg/database"
	"ewallet-server-v2/internal/pkg/encryptutils"
	"ewallet-server-v2/internal/pkg/jwtutils"
	"ewallet-server-v2/internal/pkg/logger"
	"ewallet-server-v2/internal/pkg/randutils"
	"ewallet-server-v2/internal/repository"
	"ewallet-server-v2/internal/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

func initServer(cfg *config.Config) *http.Server {
	// dependencies
	db := database.InitGorm(cfg)

	// database
	gormWrapper := database.NewGormWrapper(db)
	transactor := database.NewTransactor(db)

	// utils
	jwtUtil := jwtutils.NewJwtUtil(cfg.Jwt)
	randomUtil := randutils.NewStdLibRandomUtil()
	passwordEncryptor := encryptutils.NewBcryptPasswordEncryptor(cfg.App.BCryptCost)
	walletFormatter := apputils.NewWalletNumberFormatter(cfg.App)

	// repositories
	gameBoxRepository := repository.NewGameBoxRepository(gormWrapper)
	gameAttemptRepository := repository.NewGameAttemptRepository(gormWrapper)
	resetPasswordAttemptRepository := repository.NewResetPasswordAttemptRepository(gormWrapper)
	transactionRepository := repository.NewTransactionRepository(gormWrapper)
	userRepository := repository.NewUserRepository(gormWrapper)
	walletRepository := repository.NewWalletRepository(gormWrapper)
	sourceOfFundsRepository := repository.NewSourceOfFundRepository(gormWrapper)

	// usecases
	walletUsecase := usecase.NewWalletUsecase(walletRepository, walletFormatter)
	sourceOfFundsUsecase := usecase.NewSourceOfFundsUsecase(sourceOfFundsRepository)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepository, walletUsecase, sourceOfFundsUsecase)
	gameBoxUsecase := usecase.NewGameBoxUsecase(gameBoxRepository)
	userUsecase := usecase.NewUserUsecaseImpl(userRepository)
	gameAttemptUsecase := usecase.NewGameAttemptUsecase(transactionUsecase, gameAttemptRepository, gameBoxUsecase, walletUsecase)
	authUsecase := usecase.NewAuthUsecase(jwtUtil, randomUtil, passwordEncryptor, userUsecase, walletUsecase, resetPasswordAttemptRepository)

	// handlers
	appHandler := httphandler.NewAppHandler()
	gameBoxHandler := httphandler.NewGameBoxHandler(gameBoxUsecase)
	gameAttemptHandler := httphandler.NewGameAttemptHandler(gameAttemptUsecase, walletUsecase, transactor)
	authHandler := httphandler.NewAuthHandler(authUsecase, userUsecase, transactor)
	transactionHandler := httphandler.NewTransactionHandler(transactionUsecase, userUsecase, walletUsecase, transactor)
	userHandler := httphandler.NewUserHandler(userUsecase, walletUsecase)

	// to remove the Gin's warning
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.ContextWithFallback = true // enables .Done(), .Err(), and .Value()

	registerValidators()

	// init middlewares
	authMiddleware := middleware.NewAuthMiddleware(jwtUtil)

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	// registering middlewares
	middlewares := []gin.HandlerFunc{
		middleware.ErrorHandler(),
		middleware.Logger(),
		gin.Recovery(),
		cors.New(corsConfig),
	}
	r.Use(middlewares...)

	// registering routes
	r.NoRoute(appHandler.RouteNotFound)
	r.GET("/", appHandler.Index)

	r.POST("/v1/auth/login", authHandler.Login)
	r.POST("/v1/auth/register", authHandler.Register)

	r.POST("/auth/reset-passwords", authHandler.ResetPassword)
	r.PUT("/auth/reset-passwords", authHandler.ConfirmResetPassword)

	ar := r.Group("")
	ar.Use(authMiddleware.RequireToken())
	{
		ar.GET("/game-attempts/chances", gameAttemptHandler.GetChances)
		ar.POST("/game-attempts", gameAttemptHandler.PlayGame)

		ar.GET("/game-boxes", gameBoxHandler.GetAllBoxes)

		ar.POST("/v1/transactions/top-ups", transactionHandler.TopUp)
		ar.POST("/v1/transactions/transfers", transactionHandler.Transfer)
		ar.GET("/v1/transactions/types", transactionHandler.GetTransactiontype)
		ar.GET("/v1/transactions", transactionHandler.GetTransactionList)
		ar.GET("/v1/transactions/expense/month", transactionHandler.GetThisMonthExpenseSum)
		ar.GET("/v1/transactions/expense/:month", transactionHandler.GetExpenseSumByMonth)

		ar.GET("/v1/users/me", userHandler.GetUserDetails)
		ar.POST("/v1/users/me", userHandler.UpdateProfile)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port),
		Handler: r,
	}

	return srv
}

func StartGinHttpServer(cfg *config.Config) {
	srv := initServer(cfg)

	// graceful shutdown
	go func() {
		logger.Log.Info("running server on port :", cfg.HttpServer.Port)
		if err := srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Log.Fatal("error while server listen and serve: ", err)
			}
		}
		logger.Log.Info("server is not receiving new requests...")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	graceDuration := time.Duration(cfg.HttpServer.GracePeriod) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), graceDuration)
	defer cancel()

	logger.Log.Info("attempt to shutting down the server...")
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("error shutting down server: ", err)
	}

	logger.Log.Info("http server is shutting down gracefully")
}

func registerValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}

			return name
		})

		v.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
			if valuer, ok := field.Interface().(decimal.Decimal); ok {
				return valuer.String()
			}
			return nil
		}, decimal.Decimal{})

		v.RegisterValidation("dgte", func(fl validator.FieldLevel) bool {
			data, ok := fl.Field().Interface().(string)
			if !ok {
				return false
			}
			value, err := decimal.NewFromString(data)
			if err != nil {
				return false
			}
			baseValue, err := decimal.NewFromString(fl.Param())
			if err != nil {
				return false
			}
			return value.GreaterThanOrEqual(baseValue)
		})

		v.RegisterValidation("dlte", func(fl validator.FieldLevel) bool {
			data, ok := fl.Field().Interface().(string)
			if !ok {
				return false
			}
			value, err := decimal.NewFromString(data)
			if err != nil {
				return false
			}
			baseValue, err := decimal.NewFromString(fl.Param())
			if err != nil {
				return false
			}
			return value.LessThanOrEqual(baseValue)
		})
	}
}
