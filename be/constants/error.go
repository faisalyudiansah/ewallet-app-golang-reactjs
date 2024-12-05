package constants

const (
	ISE                = "internal server error"
	InvalidAccessToken = "invalid access token"
	Unauthorization    = "unauthorization"
	UrlNotFound        = "url not found"
	RequestBodyInvalid = "request body invalid or missing"
	InvalidDateFormat  = "invalid date format"
	InvalidQueryLimit  = "invalid query limit"
	InvalidQueryPage   = "invalid query page"
	ForbiddenAccess    = "forbidden access"
)

const (
	UserInvalidEmailPassword       = "invalid email / password"
	UserEmailNotExists             = "email does not exists"
	UserEmailAlreadyExists         = "email already exists"
	UserFailedRegister             = "there was an error in the register process, try again"
	UserTokenResetPasswordNotValid = "token reset password is not valid"
)

const (
	SOFIdIsNotExists = "source of fund id not exists"
)

const (
	GameZeroChance      = "user does not have an attempt to play game"
	GameInvalidBoxIndex = "invalid box index"
)

const (
	WalletNumberIsNotExists        = "wallet number is not exists"
	WalletTransferToTheirOwnWallet = "user can not transfer to their own wallet"
	WalletBalanceIsInsufficient    = "your balance is not enough"
)
