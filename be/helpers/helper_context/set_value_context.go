package helpercontext

import (
	"context"
	"database/sql"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/constants"
)

func SetTx(c context.Context, tx *sql.Tx) context.Context {
	var ctx constants.Ctx = "ctx"
	return context.WithValue(c, ctx, tx)
}
