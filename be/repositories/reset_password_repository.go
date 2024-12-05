package repositories

import (
	"context"
	"database/sql"
	"fmt"

	helpercontext "ewallet-server-v1/helpers/helper_context"
	"ewallet-server-v1/models"
)

type ResetPasswordRepository interface {
	PostNewDataResetPassword(context.Context, int64, string) (*models.ResetPassword, error)
	IsTokenResetValid(context.Context, string) bool
	DeleteToken(context.Context, string) error
}

type ResetPasswordRepositoryImplementation struct {
	db *sql.DB
}

func NewResetPasswordRepositoryImplementation(db *sql.DB) *ResetPasswordRepositoryImplementation {
	return &ResetPasswordRepositoryImplementation{
		db: db,
	}
}

func (rp *ResetPasswordRepositoryImplementation) PostNewDataResetPassword(ctx context.Context, userId int64, token string) (*models.ResetPassword, error) {
	sql := `
		INSERT INTO ResetPasswords (user_id, token, created_at, updated_at) VALUES 
		($1, $2, NOW(), NOW()) 
		ON CONFLICT (user_id)
		DO UPDATE SET 
			token = excluded.token, 
			created_at = excluded.created_at, 
			updated_at = excluded.updated_at
		RETURNING id, user_id, token, created_at, updated_at, deleted_at;
	`
	var resetPasswordUserData models.ResetPassword
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, userId, token).Scan(
			&resetPasswordUserData.ID,
			&resetPasswordUserData.UserId,
			&resetPasswordUserData.Token,
			&resetPasswordUserData.CreatedAt,
			&resetPasswordUserData.UpdatedAt,
			&resetPasswordUserData.DeleteAt,
		)
	} else {
		err = rp.db.QueryRowContext(ctx, sql, userId, token).Scan(
			&resetPasswordUserData.ID,
			&resetPasswordUserData.UserId,
			&resetPasswordUserData.Token,
			&resetPasswordUserData.CreatedAt,
			&resetPasswordUserData.UpdatedAt,
			&resetPasswordUserData.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &resetPasswordUserData, nil
}

func (rp *ResetPasswordRepositoryImplementation) IsTokenResetValid(ctx context.Context, token string) bool {
	sql := `
		SELECT
		id
		FROM ResetPasswords
		WHERE token = $1 AND created_at > NOW() - INTERVAL'5 minutes';
	`
	var resetPasswordUserData models.ResetPassword
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		txFromCtx.QueryRowContext(ctx, sql, token).Scan(
			&resetPasswordUserData.ID,
		)
	} else {
		rp.db.QueryRowContext(ctx, sql, token).Scan(
			&resetPasswordUserData.ID,
		)
	}
	return resetPasswordUserData.ID != 0
}

func (rp *ResetPasswordRepositoryImplementation) DeleteToken(ctx context.Context, token string) error {
	sql := `
		DELETE FROM ResetPasswords WHERE token = $1;
	`
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		_, err := txFromCtx.ExecContext(ctx, sql, token)
		return err
	}
	_, err := rp.db.ExecContext(ctx, sql, token)
	return err
}
