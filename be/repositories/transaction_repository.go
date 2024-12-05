package repositories

import (
	"context"
	"database/sql"
	"fmt"

	helpercontext "ewallet-server-v1/helpers/helper_context"
	"ewallet-server-v1/models"
)

type TransactionUserRepository interface {
	GetListTransactionsRepository(context.Context, int64, string, string, string, int64, int64, string, string) ([]models.TransactionUserAndSourceOfFund, int64, error)
	PostTransactionUserPivot(context.Context, int64, int64) error
	PostNewTransaction(context.Context, int64, int64, float64, string, string) (*models.Transaction, error)
}

type TransactionUserRepositoryImplementation struct {
	db *sql.DB
}

func NewTransactionUserRepositoryImplementation(db *sql.DB) *TransactionUserRepositoryImplementation {
	return &TransactionUserRepositoryImplementation{
		db: db,
	}
}

func (br *TransactionUserRepositoryImplementation) GetListTransactionsRepository(ctx context.Context, userId int64, query string, sortBy string, sort string, limit int64, offset int64, startDate string, endDate string) ([]models.TransactionUserAndSourceOfFund, int64, error) {
	searchKey := "%" + query + "%"
	var countSql string
	var countArgs []interface{}

	countSql = `
		SELECT COUNT(*)
		FROM transactions t 
			JOIN transactionusers tu ON t.id = tu.transaction_id 
			JOIN users u ON u.id = tu.user_id 
			JOIN sourcefunds s ON s.id = t.source_id 
			JOIN users u2 ON t.recipient_id = u2.id
		WHERE u.id = $1 AND t.description ILIKE $2`
	countArgs = []interface{}{userId, searchKey}
	if startDate != "" && endDate != "" {
		countSql += " AND t.transaction_time BETWEEN $3 AND $4"
		countArgs = append(countArgs, startDate, endDate)
	}

	var totalCount int64
	err := br.db.QueryRowContext(ctx, countSql, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	var sqlStr string
	var args []interface{}

	sqlStr = `
		SELECT 
			t.id, t.source_id, t.recipient_id, t.transaction_time, t.amount, t.description, t.created_at, t.updated_at, t.deleted_at,
			u2.id, u2.fullname, u2.created_at, u2.updated_at, u2.deleted_at,
			s.id, s.name, s.created_at, s.updated_at, s.deleted_at
		FROM transactions t 
			JOIN transactionusers tu ON t.id = tu.transaction_id 
			JOIN users u ON u.id = tu.user_id 
			JOIN sourcefunds s ON s.id = t.source_id 
			JOIN users u2 ON t.recipient_id = u2.id
		WHERE u.id = $1 AND t.description ILIKE $2`
	args = []interface{}{userId, searchKey}

	if startDate != "" && endDate != "" {
		sqlStr += " AND t.transaction_time BETWEEN $3 AND $4"
		args = append(args, startDate, endDate)
	}

	if startDate != "" && endDate != "" {
		sqlStr += fmt.Sprintf(" ORDER BY %s %s LIMIT $5 OFFSET $6", sortBy, sort)
	} else {
		sqlStr += fmt.Sprintf(" ORDER BY %s %s LIMIT $3 OFFSET $4", sortBy, sort)
	}
	args = append(args, limit, offset)

	rows, err := br.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}
	defer rows.Close()

	transactionsList := []models.TransactionUserAndSourceOfFund{}
	for rows.Next() {
		var transaction models.Transaction
		var user models.User
		var sourceOfFund models.SourceOfFund
		err := rows.Scan(
			&transaction.ID, &transaction.SourceId, &transaction.RecipientId, &transaction.TransactionTime, &transaction.Amount, &transaction.Description, &transaction.CreatedAt, &transaction.UpdatedAt, &transaction.DeleteAt,
			&user.ID, &user.FullName, &user.CreatedAt, &user.UpdatedAt, &user.DeleteAt,
			&sourceOfFund.ID, &sourceOfFund.Name, &sourceOfFund.CreatedAt, &sourceOfFund.UpdatedAt, &sourceOfFund.DeleteAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, 0, err
		}
		resp := models.TransactionUserAndSourceOfFund{
			Transaction:  transaction,
			User:         user,
			SourceOfFund: sourceOfFund,
		}
		transactionsList = append(transactionsList, resp)
	}

	return transactionsList, totalCount, nil
}

func (us *TransactionUserRepositoryImplementation) PostTransactionUserPivot(ctx context.Context, userId int64, txId int64) error {
	sql := `
		INSERT INTO TransactionUsers (user_id, transaction_id, created_at, updated_at) VALUES 
		($1, $2, NOW(), NOW())
		RETURNING id, user_id, transaction_id, created_at, updated_at, deleted_at;
	`
	var TransactionUser models.TransactionUser
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, userId, txId).Scan(
			&TransactionUser.ID,
			&TransactionUser.UserId,
			&TransactionUser.TransactionId,
			&TransactionUser.CreatedAt,
			&TransactionUser.UpdatedAt,
			&TransactionUser.DeleteAt,
		)
	} else {
		err = us.db.QueryRowContext(ctx, sql, userId, txId).Scan(
			&TransactionUser.ID,
			&TransactionUser.UserId,
			&TransactionUser.TransactionId,
			&TransactionUser.CreatedAt,
			&TransactionUser.UpdatedAt,
			&TransactionUser.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (us *TransactionUserRepositoryImplementation) PostNewTransaction(ctx context.Context, recipientId int64, sofId int64, amount float64, sofName string, customDesc string) (*models.Transaction, error) {
	sql := `
		INSERT INTO Transactions (source_id, recipient_id, transaction_time, amount, description, created_at, updated_at) VALUES 
		($1, $2, NOW(), $3, $4, NOW(), NOW())
		RETURNING id, source_id, recipient_id, transaction_time, amount, description, created_at, updated_at, deleted_at;
	`
	var tx models.Transaction
	var err error
	var description = customDesc
	if customDesc == "" {
		description = "Top Up from " + sofName
	}
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, sofId, recipientId, amount, description).Scan(
			&tx.ID,
			&tx.SourceId,
			&tx.RecipientId,
			&tx.TransactionTime,
			&tx.Amount,
			&tx.Description,
			&tx.CreatedAt,
			&tx.UpdatedAt,
			&tx.DeleteAt,
		)
	} else {
		err = us.db.QueryRowContext(ctx, sql, sofId, recipientId, amount, description).Scan(
			&tx.ID,
			&tx.SourceId,
			&tx.RecipientId,
			&tx.TransactionTime,
			&tx.Amount,
			&tx.Description,
			&tx.CreatedAt,
			&tx.UpdatedAt,
			&tx.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &tx, nil
}
