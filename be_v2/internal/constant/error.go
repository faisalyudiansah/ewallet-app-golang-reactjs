package constant

const (
	EntityNotFoundErrorMessage           = "%s not found"
	InternalServerErrorMessage           = "currently our server is facing unexpected error, please try again later"
	ResetPasswordErrorMessage            = "please try again later"
	InvalidResetPasswordCodeErrorMessage = "invalid reset password code"
	TransferToSameWalletErrorMessage     = "you cannot transfer to your own wallet"
	InsufficientWalletFundErrorMessage   = "insufficient wallet fund"
	InsufficientTopUpChanceErrorMessage  = "top up more to play more games"
	InvalidLoginCredentialsErrorMessage  = "invalid combination of email or password"
	UserAlreadyRegisteredError           = "another user already registered with the same email"
	ValidationError                      = "input validation error"
	InvalidJsonUnmarshallError           = "invalid JSON format"
	JsonSyntaxError                      = "invalid JSON syntax"
	InvalidJsonValueTypeError            = "invalid value for %s"
)
