package logic

import (
	"context"
	"github.com/pkg/errors"
	"go_code/Doul/common/xerr"

	"go_code/Doul/app/usercenter/cmd/rpc/internal/svc"
	"go_code/Doul/app/usercenter/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIsFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIsFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIsFollowLogic {
	return &CheckIsFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckIsFollowLogic) CheckIsFollow(in *user.CheckIsFollowReq) (*user.CheckIsFollowResp, error) {
	isFollow, err := l.svcCtx.RelationModel.CheckFollowByFollowerAndFollowing(l.ctx, in.FollowerId, in.FollowingId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "database error when check follow relation")
	}
	return &user.CheckIsFollowResp{
		IsFollow: isFollow,
	}, nil
}
