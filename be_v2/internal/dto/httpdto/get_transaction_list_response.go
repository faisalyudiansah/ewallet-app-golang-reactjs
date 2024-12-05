package httpdto

import (
	"time"

	"ewallet-server-v2/internal/dto/appdto"
	"ewallet-server-v2/internal/dto/pagedto"

	"github.com/shopspring/decimal"
)

type GetTransactionListEntryResponse struct {
	TransactionId                 int64           `json:"transaction_id"`
	TransactionRefId              *int64          `json:"transaction_ref_id"`
	WalletId                      int64           `json:"wallet_id"`
	TransactionTypeId             int64           `json:"transaction_type_id"`
	TransactionAdditionalDetailId int64           `json:"transaction_additional_detail_id"`
	Amount                        decimal.Decimal `json:"amount"`
	Description                   string          `json:"description"`
	TransactionDate               time.Time       `json:"transaction_date"`
}

type GetTransactionListResponse struct {
	Entries  []GetTransactionListEntryResponse `json:"entries"`
	PageInfo pagedto.PageInfoDto               `json:"page_info"`
}

func ConvertToGetTransactionListResponse(transactionList *appdto.TransactionListDto) GetTransactionListResponse {
	entries := []GetTransactionListEntryResponse{}
	for _, t := range transactionList.Entries {
		entries = append(entries, GetTransactionListEntryResponse{
			TransactionId:                 t.TransactionId,
			TransactionRefId:              t.TransactionRefId,
			WalletId:                      t.WalletId,
			TransactionTypeId:             t.TransactionTypeId,
			TransactionAdditionalDetailId: t.TransactionAdditionalDetailId,
			Amount:                        t.Amount,
			Description:                   t.Description,
			TransactionDate:               t.CreatedAt,
		})
	}

	return GetTransactionListResponse{
		Entries:  entries,
		PageInfo: transactionList.PageInfo,
	}
}
