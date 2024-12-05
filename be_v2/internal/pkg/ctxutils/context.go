package ctxutils

import (
	"context"

	"ewallet-server-v2/internal/constant"
)

func GetUserId(ctx context.Context) int64 {
	val := ctx.Value(constant.ContextUserId)

	switch val := val.(type) {
	case int64:
		return int64(val)
	default:
		return 0
	}
}
