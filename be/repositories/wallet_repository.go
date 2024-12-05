package repositories

import (
	"context"
	"database/sql"
	"fmt"

	helpercontext "git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/helpers/helper_context"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/models"
)

type WalletRepository interface {
	PostCreateWalletUser(context.Context, int64) (*models.Wallet, error)
	PutChangeBalanceWallet(context.Context, float64, int64) (*models.Wallet, error)
	GetWalletByIdUser(context.Context, int64) (*models.Wallet, error)
	GetWalletByWalletNumber(context.Context, string) (*models.Wallet, error)
}

type WalletRepositoryImplementation struct {
	db *sql.DB
}

func NewWalletRepositoryImplementation(db *sql.DB) *WalletRepositoryImplementation {
	return &WalletRepositoryImplementation{
		db: db,
	}
}

func (w *WalletRepositoryImplementation) PostCreateWalletUser(ctx context.Context, userId int64) (*models.Wallet, error) {
	sql := fmt.Sprintf("INSERT INTO Wallets (user_id, wallet_number, balance, created_at, updated_at) VALUES ($1, CONCAT('777', SUBSTRING(TO_CHAR(%d, '0000000000'), 2)), 0.0, NOW(), NOW()) RETURNING id, user_id, wallet_number, balance, created_at, updated_at, deleted_at;", userId)
	var wallet models.Wallet
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, userId).Scan(
			&wallet.ID,
			&wallet.UserId,
			&wallet.WalletNumber,
			&wallet.Balance,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
			&wallet.DeleteAt,
		)
	} else {
		err = w.db.QueryRowContext(ctx, sql, userId).Scan(
			&wallet.ID,
			&wallet.UserId,
			&wallet.WalletNumber,
			&wallet.Balance,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
			&wallet.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &wallet, nil
}

func (w *WalletRepositoryImplementation) PutChangeBalanceWallet(ctx context.Context, amount float64, userId int64) (*models.Wallet, error) {
	sql := `
		UPDATE Wallets SET 
			balance = $1,
			updated_at = NOW()
		WHERE user_id = $2
		RETURNING id, user_id, wallet_number, balance, created_at, updated_at, deleted_at;
	`
	var wallet models.Wallet
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, amount, userId).Scan(
			&wallet.ID,
			&wallet.UserId,
			&wallet.WalletNumber,
			&wallet.Balance,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
			&wallet.DeleteAt,
		)
	} else {
		err = w.db.QueryRowContext(ctx, sql, amount, userId).Scan(
			&wallet.ID,
			&wallet.UserId,
			&wallet.WalletNumber,
			&wallet.Balance,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
			&wallet.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &wallet, nil
}

func (w *WalletRepositoryImplementation) GetWalletByIdUser(ctx context.Context, userId int64) (*models.Wallet, error) {
	sql := `
		SELECT
		id, user_id, wallet_number, balance, created_at, updated_at, deleted_at
		FROM Wallets
		WHERE user_id = $1 AND deleted_at IS NULL;
	`
	var wallet models.Wallet
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, userId).Scan(
			&wallet.ID,
			&wallet.UserId,
			&wallet.WalletNumber,
			&wallet.Balance,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
			&wallet.DeleteAt,
		)
	} else {
		err = w.db.QueryRowContext(ctx, sql, userId).Scan(
			&wallet.ID,
			&wallet.UserId,
			&wallet.WalletNumber,
			&wallet.Balance,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
			&wallet.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &wallet, nil
}

func (w *WalletRepositoryImplementation) GetWalletByWalletNumber(ctx context.Context, wn string) (*models.Wallet, error) {
	sql := `
		SELECT
		id, user_id, wallet_number, balance, created_at, updated_at, deleted_at
		FROM Wallets
		WHERE wallet_number = $1 AND deleted_at IS NULL;
	`
	var wallet models.Wallet
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, wn).Scan(
			&wallet.ID,
			&wallet.UserId,
			&wallet.WalletNumber,
			&wallet.Balance,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
			&wallet.DeleteAt,
		)
	} else {
		err = w.db.QueryRowContext(ctx, sql, wn).Scan(
			&wallet.ID,
			&wallet.UserId,
			&wallet.WalletNumber,
			&wallet.Balance,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
			&wallet.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &wallet, nil
}
