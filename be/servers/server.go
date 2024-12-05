package servers

import (
	"database/sql"

	"ewallet-server-v1/controllers"
	"ewallet-server-v1/helpers"
	"ewallet-server-v1/helpers/logger"
	"ewallet-server-v1/repositories"
	"ewallet-server-v1/services"
)

type HandlerOps struct {
	UserController            *controllers.UserController
	TransactionUserController *controllers.TransactionUserController
	ResetPasswordController   *controllers.ResetPassowordController
	WalletController          *controllers.WalletController
	GameController            *controllers.GameController
}

func SetupController(db *sql.DB) *HandlerOps {
	logrusLogger := logger.NewLogger()
	logger.SetLogger(logrusLogger)

	bcrypt := helpers.NewBcryptStruct()
	jwt := helpers.NewJWTProviderHS256()
	validationReqBody := helpers.NewValidationReqBody()
	generateNumber := helpers.NewGenerateNumber()
	getParam := helpers.NewGetParams()

	transactionsRepository := repositories.NewTransactionRepositoryImplementation(db)
	userRepository := repositories.NewUserRepositoryImplementation(db)
	resetPasswordRepository := repositories.NewResetPasswordRepositoryImplementation(db)
	walletRepository := repositories.NewWalletRepositoryImplementation(db)
	transactionUserRepository := repositories.NewTransactionUserRepositoryImplementation(db)
	sourceOfFundRepository := repositories.NewSourceOfFundImplementation(db)
	gameRepository := repositories.NewGameRepositoryImplementation(db)

	userService := services.NewUserServiceImplementation(userRepository, walletRepository, transactionsRepository, bcrypt, jwt, generateNumber)
	resetPasswordService := services.NewResetPasswordServiceImplementation(resetPasswordRepository, userRepository, transactionsRepository, generateNumber, bcrypt)
	transactionUserService := services.NewTransactionUserServiceImplementation(transactionUserRepository, userRepository, walletRepository, transactionsRepository)
	walletService := services.NewWalletServiceImplementation(userRepository, walletRepository, transactionUserRepository, transactionsRepository, sourceOfFundRepository)
	gameService := services.NewGameServiceImplementation(gameRepository, transactionUserRepository, transactionsRepository, userRepository)

	userController := controllers.NewUserController(userService, validationReqBody)
	resetPasswordController := controllers.NewResetPassowordController(resetPasswordService, validationReqBody, getParam)
	transactionUserController := controllers.NewTransactionUserController(transactionUserService, getParam)
	walletController := controllers.NewWalletController(walletService, validationReqBody)
	gameController := controllers.NewGameController(gameService, validationReqBody)

	return &HandlerOps{
		UserController:            userController,
		TransactionUserController: transactionUserController,
		ResetPasswordController:   resetPasswordController,
		WalletController:          walletController,
		GameController:            gameController,
	}
}
