package logic

import (
	"context"

	"go_code/Doul/app/usercenter/cmd/rpc/internal/svc"
	"go_code/Doul/app/usercenter/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.UserInfoReq) (*user.UserInfoReply, error) {
	userFind, err := l.svcCtx.UserModel.FindOne(l.ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	return &user.UserInfoReply{
		Id:            userFind.UserId,
		Name:          userFind.Name,
		FollowCount:   userFind.FollowCount,
		FollowerCount: userFind.FollowerCount,
	}, nil
}
