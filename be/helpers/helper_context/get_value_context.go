package helpercontext

import (
	"context"
	"database/sql"

	"ewallet-server-v1/constants"
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
