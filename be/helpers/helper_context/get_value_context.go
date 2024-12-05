package helpercontext

import (
	"context"
	"database/sql"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/constants"
)

func GetTx(c context.Context) *sql.Tx {
	var ctx constants.Ctx = "ctx"
	if tx, ok := c.Value(ctx).(*sql.Tx); ok {
		return tx
	}
	return nil
}

func GetValueUserIdFromToken(c context.Context) int64 {
	var key constants.ID = "userId"
	if userId, ok := c.Value(key).(int64); ok {
		return userId
	}
	return 0
}
