package logic

import (
	"context"
	"go_code/Doul/common/globalkey"

	"go_code/Doul/app/usercenter/cmd/rpc/internal/svc"
	"go_code/Doul/app/usercenter/cmd/rpc/user"

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

// CheckAuth 检查用户是否登陆
func (l *CheckAuthLogic) CheckAuth(in *user.CheckAuthReq) (*user.CheckAuthReply, error) {
	exists, err := l.svcCtx.RedisClient.Exists(globalkey.TokenPrefix + in.GetToken())
	if err != nil {
		return nil, err
	}
	return &user.CheckAuthReply{
		Authed: exists,
	}, nil
}