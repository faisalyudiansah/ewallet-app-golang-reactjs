package httpdto

import (
	"time"

	"ewallet-server-v2/internal/model"

	"github.com/shopspring/decimal"
)

type TransferResponse struct {
	TransactionId     int64           `json:"transaction_id"`
	TransactionRefId  *int64          `json:"transaction_ref_id"`
	WalletFromId      int64           `json:"wallet_from_id"`
	WalletFromNumber  string          `json:"wallet_from_number"`
	TransactionTypeId int64           `json:"transaction_type_id"`
	WalletToId        int64           `json:"wallet_to_id"`
	WalletToNumber    string          `json:"wallet_to_number"`
	Amount            decimal.Decimal `json:"amount"`
	Description       string          `json:"description"`
	TransactionDate   time.Time       `json:"transaction_date"`
}

func ConvertToTransferResponse(transaction *model.Transaction, walletFrom, walletTo *model.Wallet) TransferResponse {
	return TransferResponse{
		TransactionId:     transaction.TransactionId,
		TransactionRefId:  transaction.TransactionRefId,
		WalletFromId:      transaction.WalletId,
		WalletFromNumber:  walletFrom.WalletNumber,
		TransactionTypeId: transaction.TransactionTypeId,
		WalletToId:        transaction.TransactionAdditionalDetailId,
		WalletToNumber:    walletTo.WalletNumber,
		Amount:            transaction.Amount,
		Description:       transaction.Description,
		TransactionDate:   transaction.CreatedAt,
	}
}
