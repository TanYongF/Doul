package logic

import (
	"context"
	"github.com/pkg/errors"
	"go_code/Doul/common/xerr"

	"go_code/Doul/app/usercenter/cmd/rpc/internal/svc"
	"go_code/Doul/app/usercenter/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowingCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingCountLogic {
	return &GetFollowingCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowingCountLogic) GetFollowingCount(in *user.GetFollowingCountReq) (*user.GetFollowingCountResp, error) {
	count, err := l.svcCtx.RelationModel.CountFollowingsByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "DB error when count followings numbers")
	}
	return &user.GetFollowingCountResp{
		Count: *count,
	}, nil
}
