package common

import (
	"context"
)

var (
	CtxKeyUserid = "auth_id"
)

func GetUidFromCtx(ctx context.Context) int64 {
	var int64Uid int64
	if uid, ok := ctx.Value(CtxKeyUserid).(int64); ok {
		int64Uid = uid
	}
	return int64Uid
}
