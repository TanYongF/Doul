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

type FollowingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowingListLogic {
	return &FollowingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowingListLogic) FollowingList(req *types.FollowingListReq) (resp *types.FollowingListResp, err error) {
	followingListByDB, err := l.svcCtx.UserRpc.GetFollowingList(l.ctx, &userclient.GetFollowingListReq{
		QueryId: req.UserId,
		UserId:  common.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}
	userList := make([]types.UserPO, len(followingListByDB.Users))

	for i := 0; i < len(followingListByDB.Users); i++ {
		copier.Copy(&userList[i], followingListByDB.Users[i])
		userList[i].ID = followingListByDB.Users[i].UserId
	}
	return &types.FollowingListResp{
		UserList: userList,
	}, nil
}
