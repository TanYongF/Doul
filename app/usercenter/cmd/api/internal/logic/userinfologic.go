package logic

import (
	"context"
	"fmt"
	"go_code/Doul/app/usercenter/cmd/api/internal/svc"
	"go_code/Doul/app/usercenter/cmd/api/internal/types"
	"go_code/Doul/app/usercenter/cmd/rpc/user"
	"go_code/Doul/common/tool"
	"google.golang.org/grpc/metadata"

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
	authId := tool.GetUidFromCtx(l.ctx)
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		get := md.Get("username")
		fmt.Println(get)
	}
	// 1. Get the user information body
	userReply, err := l.svcCtx.UserRpc.GetUser(context.WithValue(l.ctx, "auth_id", authId), &user.UserInfoReq{
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
