package httpdto

import (
	"time"

	"ewallet-server-v2/internal/model"

	"github.com/shopspring/decimal"
)

type TopUpResponse struct {
	TransactionId     int64           `json:"transaction_id"`
	TransactionRefId  *int64          `json:"transaction_ref_id"`
	WalletId          int64           `json:"wallet_id"`
	WalletNumber      string          `json:"wallet_number"`
	TransactionTypeId int64           `json:"transaction_type_id"`
	SourceOfFundsId   int64           `json:"source_of_funds_id"`
	Amount            decimal.Decimal `json:"amount"`
	Description       string          `json:"description"`
	TransactionDate   time.Time       `json:"transaction_date"`
}

func ConvertToTopUpResponse(transaction *model.Transaction, wallet *model.Wallet) TopUpResponse {
	return TopUpResponse{
		TransactionId:     transaction.TransactionId,
		TransactionRefId:  transaction.TransactionRefId,
		WalletId:          transaction.WalletId,
		WalletNumber:      wallet.WalletNumber,
		TransactionTypeId: transaction.TransactionTypeId,
		SourceOfFundsId:   transaction.TransactionAdditionalDetailId,
		Amount:            transaction.Amount,
		Description:       transaction.Description,
		TransactionDate:   transaction.CreatedAt,
	}
}
