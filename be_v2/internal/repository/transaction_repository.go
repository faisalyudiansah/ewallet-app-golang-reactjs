package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"ewallet-server-v2/internal/constant"
	"ewallet-server-v2/internal/dto/appdto"
	"ewallet-server-v2/internal/dto/pagedto"
	"ewallet-server-v2/internal/model"
	"ewallet-server-v2/internal/pkg/database"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateOne(ctx context.Context, transaction model.Transaction) (*model.Transaction, error)
	GetListByWalletId(ctx context.Context, walletId int64, pageDto pagedto.PageSortDto) (*appdto.TransactionListDto, error)
	GetTransactionTypeSumMultiple(ctx context.Context, walletId, transactionTypeId int64, transactionAdditionalDetailId []int64, minAmount decimal.Decimal) (*decimal.Decimal, error)
	GetTransactionType(ctx context.Context) ([]model.TransactionType, error)
	GetThisMonthTransactionSum(ctx context.Context, walletId int64, startDate time.Time, endDate time.Time) (*appdto.TransactionSum, error)
}

type transactionRepositoryPostgreSQL struct {
	db *database.GormWrapper
}

func NewTransactionRepository(db *database.GormWrapper) *transactionRepositoryPostgreSQL {
	return &transactionRepositoryPostgreSQL{
		db: db,
	}
}

func (r *transactionRepositoryPostgreSQL) CreateOne(ctx context.Context, transaction model.Transaction) (*model.Transaction, error) {
	if err := r.db.Start(ctx).Create(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *transactionRepositoryPostgreSQL) GetListByWalletId(ctx context.Context, walletId int64, pageDto pagedto.PageSortDto) (*appdto.TransactionListDto, error) {
	allowedSorts := map[string]string{
		"date":   "created_at",
		"amount": "amount",
		"to":     "transaction_additional_detail_id",
	}

	allowedOrderDir := map[string]struct{}{
		"asc":  {},
		"desc": {},
	}

	orderBy := []string{}

	res := appdto.TransactionListDto{
		Entries: []model.Transaction{},
		PageInfo: pagedto.PageInfoDto{
			Page:            pageDto.Page,
			Limit:           pageDto.Limit,
			Sorts:           pageDto.Sorts,
			Search:          pageDto.Search,
			FilterStartDate: pageDto.FilterStartDate,
			FilterEndDate:   pageDto.FilterEndDate,
			FilterType:      pageDto.FilterType,
			TotalRow:        0,
		},
	}

	q := r.db.Start(ctx).Model(&model.Transaction{}).Where("wallet_id", walletId)

	if pageDto.Search != "" {
		q.Where("transaction_description ILIKE ?", fmt.Sprintf("%%%s%%", pageDto.Search))
	}

	if pageDto.FilterStartDate != nil {
		q.Where("created_at >= date_trunc('day', ?::timestamp)", pageDto.FilterStartDate)
	}

	if pageDto.FilterEndDate != nil {
		q.Where("created_at <= date_trunc('day', ?::timestamp) + interval '1 day' - interval '1 sec'", pageDto.FilterEndDate)
	}

	if pageDto.FilterType != 0 {
		q.Where("transaction_type_id = ?", pageDto.FilterType)
	}

	for _, s := range pageDto.Sorts {
		ord, ok := allowedSorts[strings.ToLower(s.Column)]
		if !ok {
			continue
		}

		if s.Column == "to" {
			q.Where("transaction_type_id", constant.TransactionTypeTransfer)
		}

		dir := s.OrderDir
		_, ok = allowedOrderDir[strings.ToLower(s.OrderDir)]
		if !ok {
			dir = "asc"
		}

		orderBy = append(orderBy, fmt.Sprintf("%s %s", ord, dir))
	}

	q = q.Session(&gorm.Session{})

	page := 1
	if pageDto.Page > 0 {
		page = pageDto.Page
	}

	limit := 10
	if pageDto.Limit > 0 {
		limit = pageDto.Limit
	}

	off := limit * (page - 1)

	q2 := q.Limit(limit).Offset(off)

	// handling pagination, can be extracted to its own helper function / method
	if len(orderBy) == 0 {
		q2.Order("created_at DESC")
	}

	for _, o := range orderBy {
		q2.Order(o)
	}

	q2.Order("transaction_id DESC")

	if err := q2.Find(&res.Entries).Error; err != nil {
		return nil, err
	}

	if err := q.Count(&res.PageInfo.TotalRow).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *transactionRepositoryPostgreSQL) GetTransactionType(ctx context.Context) ([]model.TransactionType, error) {
	res := []model.TransactionType{}

	q := r.db.Start(ctx).Model(&model.TransactionType{})

	err := q.Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *transactionRepositoryPostgreSQL) GetTransactionTypeSumMultiple(
	ctx context.Context,
	walletId,
	transactionTypeId int64,
	transactionAdditionalDetailId []int64,
	minAmount decimal.Decimal,
) (*decimal.Decimal, error) {
	res := decimal.NewFromInt(0)

	sql := `
		SELECT
			SUM(t.amount / ?)
		FROM
			transactions t 
		JOIN transaction_types tt ON tt.transaction_type_id = t.transaction_type_id
		WHERE
			t.wallet_id = ?
		AND t.transaction_type_id = ?
		AND t.deleted_at IS NULL
		AND t.amount >= ?
	`

	if len(transactionAdditionalDetailId) > 0 {
		placeholders := strings.TrimSuffix(strings.Repeat("?,", len(transactionAdditionalDetailId)), ",")

		sql += fmt.Sprintf("\nAND t.transaction_additional_detail_id IN (%s)", placeholders)
	}

	args := []interface{}{minAmount, walletId, transactionTypeId, minAmount}
	for _, tid := range transactionAdditionalDetailId {
		args = append(args, tid)
	}

	q := r.db.Start(ctx).Raw(sql, args...)

	if err := q.Scan(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *transactionRepositoryPostgreSQL) GetThisMonthTransactionSum(
	ctx context.Context,
	walletId int64,
	startDate time.Time,
	endDate time.Time,
) (*appdto.TransactionSum, error) {
	var res appdto.TransactionSum

	sql := `
        SELECT SUM(t.amount), date_trunc('month', ?::timestamp) as month
        FROM "transactions" t
        JOIN wallets w ON w.wallet_id = t.wallet_id
        WHERE t.created_at >= date_trunc('month', ?::timestamp) 
		AND t.created_at <= date_trunc('month', ?::timestamp) + interval '1 month' - interval '1 day'
		AND t.amount < 0
		AND t.wallet_id = ? AND t.deleted_at IS NULL
	`

	q := r.db.Start(ctx).Raw(sql, startDate, startDate, endDate, walletId)

	if err := q.Scan(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}
