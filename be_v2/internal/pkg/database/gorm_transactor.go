package database

import (
	"context"

	"gorm.io/gorm"
)

type gormTxKey struct{}

type Transactor interface {
	Transaction(ctx context.Context, txFunc func(txCtx context.Context) error) error
}

type transactor struct {
	db *gorm.DB
}

func NewTransactor(db *gorm.DB) Transactor {
	return &transactor{db: db}
}

func (t *transactor) Transaction(ctx context.Context, txFunc func(txCtx context.Context) error) error {
	err := t.db.Transaction(func(tx *gorm.DB) error {
		err := txFunc(injectTx(ctx, tx))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// injectTx injects transaction to context
func injectTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, gormTxKey{}, tx)
}

// extractTx extracts transaction from context
func extractTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(gormTxKey{}).(*gorm.DB); ok {
		return tx
	}
	return nil
}

// thin wrapper to enable transaction in repository
type GormWrapper struct {
	db *gorm.DB
}

func NewGormWrapper(db *gorm.DB) *GormWrapper {
	return &GormWrapper{
		db: db,
	}
}

func (w *GormWrapper) Start(ctx context.Context) *gorm.DB {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.WithContext(ctx)
	}
	return w.db.WithContext(ctx)
}
