package repositories

import (
	"context"
	"database/sql"
	"fmt"

	helpercontext "git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/helpers/helper_context"
)

type TransactionRepository interface {
	Atomic(c context.Context, fn func(context.Context) (any, error)) (any, error)
}

type TransactionRepositoryImplementation struct {
	db *sql.DB
}

func NewTransactionRepositoryImplementation(db *sql.DB) *TransactionRepositoryImplementation {
	return &TransactionRepositoryImplementation{
		db: db,
	}
}

func (dc *TransactionRepositoryImplementation) Atomic(c context.Context, fn func(context.Context) (any, error)) (any, error) {
	tx, err := dc.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		}
		err = tx.Commit()
	}()

	result, err := fn(helpercontext.SetTx(c, tx))
	if err != nil {
		return nil, err
	}
	return result, nil
}
