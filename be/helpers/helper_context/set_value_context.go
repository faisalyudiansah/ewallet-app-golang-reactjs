package helpercontext

import (
	"context"
	"database/sql"

	"ewallet-server-v1/constants"
)

func SetTx(c context.Context, tx *sql.Tx) context.Context {
	var ctx constants.Ctx = "ctx"
	return context.WithValue(c, ctx, tx)
}
