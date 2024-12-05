package constant

const (
	TopupAmountToParticipate = 10_000_000
)

const (
	SourceOfFundBankTransfer = 1
	SourceOfFundCreditCard   = 2
	SourceOfFundCash         = 3
	SourceOfFundReward       = 4
)

const (
	StringSourceOfFundBankTransfer = "BANK_TRANSFER"
	StringSourceOfFundCreditCard   = "CREDIT_CARD"
	StringSourceOfFundCash         = "CASH"
	StringSourceOfFundReward       = "REWARD"
)

var sourceOfFundString = map[int64]string{
	SourceOfFundBankTransfer: StringSourceOfFundBankTransfer,
	SourceOfFundCreditCard:   StringSourceOfFundCreditCard,
	SourceOfFundCash:         StringSourceOfFundCash,
	SourceOfFundReward:       StringSourceOfFundReward,
}

func ConvertSourceOfFundToString(sourceOfFundId int64) string {
	return sourceOfFundString[sourceOfFundId]
}

const (
	ReadableSourceOfFundBankTransfer = "Bank Transfer"
	ReadableSourceOfFundCreditCard   = "Credit Card"
	ReadableSourceOfFundCash         = "Cash"
	ReadableSourceOfFundReward       = "Reward"
)

var sourceOfFundReadable = map[int64]string{
	SourceOfFundBankTransfer: ReadableSourceOfFundBankTransfer,
	SourceOfFundCreditCard:   ReadableSourceOfFundCreditCard,
	SourceOfFundCash:         ReadableSourceOfFundCash,
	SourceOfFundReward:       ReadableSourceOfFundReward,
}

func ConvertSourceOfFundToReadable(sourceOfFundId int64) string {
	return sourceOfFundReadable[sourceOfFundId]
}

const (
	TransactionTypeTopUp    = 1
	TransactionTypeTransfer = 2
)

const (
	GameBoxLimit = 9
)

const (
	ResetPasswordCodeLength = 100
)

const (
	ResetPasswordValidDuration = 24 * 60 // in minutes
)

const (
	ContextUserId = "ctx-user-id"
)

const (
	MessageTopUpDescription = "Top Up from %s"
	MessageResponseSuccess  = "success"
)

var timeLayoutTranslate map[string]string = map[string]string{
	"2006-01-02": "YYYY-MM-DD",
}

func ConvertGoTimeLayoutToReadable(layout string) string {
	return timeLayoutTranslate[layout]
}
