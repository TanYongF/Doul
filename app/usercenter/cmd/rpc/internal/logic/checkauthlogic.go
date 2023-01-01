package logic

import (
	"context"
	"github.com/pkg/errors"
	"go_code/Doul/app/usercenter/cmd/rpc/internal/svc"
	"go_code/Doul/app/usercenter/cmd/rpc/user"
	"go_code/Doul/common/globalkey"
	"go_code/Doul/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAuthLogic {
	return &CheckAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CheckAuth check the user has authed
func (l *CheckAuthLogic) CheckAuth(in *user.CheckAuthReq) (*user.CheckAuthReply, error) {
	exists, err := l.svcCtx.RedisClient.Exists(globalkey.TokenPrefix + in.GetToken())
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CACHE_ERROR), "failed to access redis when check user's auth")
	}

	//update the token's expire time (24 hours).
	l.svcCtx.RedisClient.Expire(globalkey.TokenPrefix+in.GetToken(), int(globalkey.TokenExpireTime.Seconds()))
	return &user.CheckAuthReply{
		Authed: exists,
	}, nil
}
