package tool

import (
	"context"
)

var (
	CtxKeyUserid = "auth_id"
)

// GetUidFromCtx 通过 context 获取用户ID
func GetUidFromCtx(ctx context.Context) int64 {
	var int64Uid int64
	if uid, ok := ctx.Value(CtxKeyUserid).(int64); ok {
		int64Uid = uid
	} else if uid, ok := ctx.Value(CtxKeyUserid).(int); ok {
		int64Uid = int64(uid)
	}
	return int64Uid
}
