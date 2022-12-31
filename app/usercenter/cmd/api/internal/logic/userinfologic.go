package logic

import (
	"context"
	"go_code/Doul/app/usercenter/cmd/rpc/user"

	"go_code/Doul/app/usercenter/cmd/api/internal/svc"
	"go_code/Doul/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo(req *types.InfoReq) (resp *types.InfoRes, err error) {
	userReply, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.UserInfoReq{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &types.InfoRes{
		FollowCount:   userReply.FollowCount,
		FollowerCount: userReply.FollowerCount,
		ID:            userReply.Id,
		IsFollow:      false,
		Name:          userReply.Name,
	}, err
}
