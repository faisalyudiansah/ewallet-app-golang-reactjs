package appdto

import (
	"time"

	"ewallet-server-v2/internal/dto/pagedto"
	"ewallet-server-v2/internal/model"
)

type TransactionListDto struct {
	Entries  []model.Transaction
	PageInfo pagedto.PageInfoDto
}

type TransactionTypeDto struct {
	Entries []model.TransactionType
}

type TransactionTypeResponse struct {
	TypeId   int64  `json:"type_id"`
	TypeName string `json:"type_name"`
}

type TransactionSum struct {
	Sum   float64   `json:"sum"`
	Month time.Time `json:"month"`
}

type TransactionByMonthParams struct {
	Month string `uri:"month"`
}

func ConvertToTransactionTypeDto(transactionTypes []model.TransactionType) []TransactionTypeResponse {
	entries := []TransactionTypeResponse{}
	for _, t := range transactionTypes {
		entries = append(entries, TransactionTypeResponse{
			TypeId:   t.TransactionTypeId,
			TypeName: t.TypeName,
		})
	}

	return entries
}
