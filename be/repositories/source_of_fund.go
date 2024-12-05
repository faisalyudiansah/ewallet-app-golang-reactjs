package repositories

import (
	"context"
	"database/sql"
	"fmt"

	helpercontext "ewallet-server-v1/helpers/helper_context"
	"ewallet-server-v1/models"
)

type SourceOfFundRepository interface {
	GetSourceOfFundById(context.Context, int64) (*models.SourceOfFund, error)
}

type SourceOfFundImplementation struct {
	db *sql.DB
}

func NewSourceOfFundImplementation(db *sql.DB) *SourceOfFundImplementation {
	return &SourceOfFundImplementation{
		db: db,
	}
}

func (sof *SourceOfFundImplementation) GetSourceOfFundById(ctx context.Context, id int64) (*models.SourceOfFund, error) {
	sql := `
		SELECT
			s.id,
			s.name,
			s.created_at,
			s.updated_at,
			s.deleted_at
		FROM SourceFunds s
		WHERE id = $1 AND deleted_at IS NULL;
	`
	var source models.SourceOfFund
	var err error
	txFromCtx := helpercontext.GetTx(ctx)
	if txFromCtx != nil {
		err = txFromCtx.QueryRowContext(ctx, sql, id).Scan(
			&source.ID,
			&source.Name,
			&source.CreatedAt,
			&source.UpdatedAt,
			&source.DeleteAt,
		)
	} else {
		err = sof.db.QueryRowContext(ctx, sql, id).Scan(
			&source.ID,
			&source.Name,
			&source.CreatedAt,
			&source.UpdatedAt,
			&source.DeleteAt,
		)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &source, nil
}
