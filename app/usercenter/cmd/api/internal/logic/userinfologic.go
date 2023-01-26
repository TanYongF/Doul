package logic

import (
	"context"
	"go_code/Doul/app/usercenter/cmd/rpc/user"
	"go_code/Doul/common"

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

	//Get the current user_id from context
	authId := common.GetUidFromCtx(l.ctx)

	// 1. Get the user information body
	userReply, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.UserInfoReq{
		QueryId: req.UserId,
		UserId:  authId,
	})
	if err != nil {
		return nil, err
	}

	// 2. Build the resp body by copying information
	userPo := types.UserPO{}
	userPo.Name = userReply.Name
	userPo.IsFollow = userReply.IsFollow
	userPo.FollowCount = userReply.FollowCount
	userPo.FollowerCount = userReply.FollowerCount
	userPo.ID = userReply.Id

	return &types.InfoRes{
		User: userPo,
	}, err
}
