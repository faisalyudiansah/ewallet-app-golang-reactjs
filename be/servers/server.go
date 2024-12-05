package servers

import (
	"database/sql"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/controllers"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/helpers"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/helpers/logger"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/repositories"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/services"
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
