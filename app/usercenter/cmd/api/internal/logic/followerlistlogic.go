package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go_code/Doul/app/usercenter/cmd/api/internal/svc"
	"go_code/Doul/app/usercenter/cmd/api/internal/types"
	"go_code/Doul/app/usercenter/cmd/rpc/userclient"
	"go_code/Doul/common"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowerListLogic) FollowerList(req *types.FollowerListReq) (resp *types.FollowerListResp, err error) {
	FollowerList, err := l.svcCtx.UserRpc.GetFollowerList(l.ctx, &userclient.GetFollowerListReq{
		UserId:  common.GetUidFromCtx(l.ctx),
		QueryId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	userList := make([]types.UserPO, len(FollowerList.Users))
	for i := 0; i < len(FollowerList.Users); i++ {
		copier.Copy(&userList[i], FollowerList.Users[i])
		userList[i].ID = FollowerList.Users[i].UserId
	}

	return &types.FollowerListResp{
		UserList: userList,
	}, nil
}
