package logic

import (
	"context"
	"github.com/jinzhu/copier"
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

	//Get the current user_id from context
	authId := l.ctx.Value("user_id").(int64)

	// 1. Get the user information body
	userReply, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.UserInfoReq{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	// 2. Get the follow-relation between A and B
	isFollow, err := l.svcCtx.UserRpc.CheckIsFollow(l.ctx, &user.CheckIsFollowReq{
		FollowerId:  userReply.Id,
		FollowingId: authId,
	})
	if err != nil {
		return nil, err
	}

	// 3. Build the resp body by copying information
	user := types.UserPO{}
	copier.Copy(&user, &userReply)
	user.IsFollow = isFollow.GetIsFollow()
	return &types.InfoRes{
		User: user,
	}, err
}
