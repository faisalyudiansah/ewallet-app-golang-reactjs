package dtos

import (
	"time"

	"github.com/shopspring/decimal"
)

type ResponseMessageOnly struct {
	Message string `json:"message"`
}

type ResponseApiError struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

type ResponseAccessToken struct {
	AccessToken string `json:"access_token"`
}

type ResponseTokenResetPassword struct {
	TokenResetPassword string `json:"token"`
	LinkResetPassword  string `json:"URL_reset_password"`
}

type ResponseLoginUser struct {
	Message string              `json:"message"`
	Result  ResponseAccessToken `json:"result"`
}

type ResponseForgetPassword struct {
	Message string                     `json:"message"`
	Result  ResponseTokenResetPassword `json:"result"`
}

type ResponseSuccessTransferWithMessage struct {
	Message string                  `json:"message"`
	Result  ResponseSuccessTransfer `json:"result"`
}

type ResponseListTransactions struct {
	Message        string                              `json:"message"`
	Limit          int64                               `json:"per_page"`
	Page           int64                               `json:"page"`
	PageCount      int64                               `json:"page_count"`
	TotalCountData int64                               `json:"total_count_data"`
	Result         []ResponseSingleDataTransactionUser `json:"data"`
}

type ResponseSingleDataTransactionUser struct {
	ID              int64           `json:"id"`
	SourceId        int64           `json:"source_id"`
	SourceOfFund    SourceOfFund    `json:"source_of_fund_data"`
	RecipientId     int64           `json:"recipient_id"`
	Recipient       ResponseUser    `json:"recipient_Data"`
	TransactionTime time.Time       `json:"transaction_time"`
	Amount          decimal.Decimal `json:"amount"`
	Description     string          `json:"description"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeleteAt        *time.Time      `json:"deleted_at"`
}

type SourceOfFund struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeleteAt  *time.Time `json:"deleted_at"`
}

type ResponseUser struct {
	ID         int64      `json:"id"`
	Email      string     `json:"email,omitempty"`
	ChanceGame int        `json:"chance_game,omitempty"`
	FullName   string     `json:"fullname,omitempty"`
	BirthDate  time.Time  `json:"birthdate"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeleteAt   *time.Time `json:"deleted_at"`
}

type ResponseWallet struct {
	ID           int64           `json:"id"`
	UserId       int64           `json:"user_id"`
	WalletNumber string          `json:"wallet_number"`
	Balance      decimal.Decimal `json:"balance"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeleteAt     *time.Time      `json:"deleted_at"`
}

type ResponseUserAndWallet struct {
	ID         int64          `json:"id"`
	Email      string         `json:"email"`
	ChanceGame int            `json:"chance_game"`
	FullName   string         `json:"fullname"`
	BirthDate  time.Time      `json:"birthdate"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeleteAt   *time.Time     `json:"deleted_at"`
	Wallet     ResponseWallet `json:"wallet"`
}

type ResponseShowOneUser struct {
	Message string       `json:"message"`
	Result  ResponseUser `json:"result"`
}

type ResponseShowOneUserWithWallet struct {
	Message string                `json:"message"`
	Result  ResponseUserAndWallet `json:"result"`
}

type ResponseShowListGachaBox struct {
	Message string `json:"message"`
	Result  []int  `json:"box"`
}

type ResponseShowReward struct {
	Message string `json:"message"`
	Result  string `json:"reward"`
}

type ResponseSuccessTransfer struct {
	UserId           int64           `json:"sender_id"`
	RecipientId      int64           `json:"recipient_id"`
	Amount           decimal.Decimal `json:"amount"`
	RemainingBelance decimal.Decimal `json:"your_remaining_balance"`
	TransactionTime  time.Time       `json:"transaction_time"`
	Description      string          `json:"description"`
}
