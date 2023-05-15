package logic

import (
	"context"
	"go_code/Doul/app/usercenter/cmd/rpc/userclient"
	"go_code/Doul/common/tool"
	"strconv"

	"go_code/Doul/app/usercenter/cmd/api/internal/svc"
	"go_code/Doul/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowActionReq) (resp *types.FollowActionResp, err error) {
	userId := tool.GetUidFromCtx(l.ctx)
	isDel, err := strconv.Atoi(req.Type)
	if err != nil {
		return nil, err
	}
	toUserId, err := strconv.ParseInt(req.ToUserId, 10, 64)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.UserRpc.FollowAction(l.ctx, &userclient.FollowActionReq{
		UserId:   userId,
		ToUserId: toUserId,
		Type:     int32(isDel),
	})
	if err != nil {
		return nil, err
	}
	return &types.FollowActionResp{}, nil
}
