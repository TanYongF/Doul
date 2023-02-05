package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_code/Doul/app/usercenter/cmd/rpc/internal/svc"
	"go_code/Doul/app/usercenter/cmd/rpc/user"
	"go_code/Doul/common/tool"
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
	userFind, err := l.svcCtx.UserModel.FindOne(l.ctx, in.QueryId)

	fmt.Println(tool.GetUidFromCtx(l.ctx))

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errors.Wrapf(ErrUserNoExistsError, "用户名id=%d未找到", in.QueryId)
		}
		return nil, err
	}

	if isFollow, err := l.svcCtx.RelationModel.CheckFollowByFollowerAndFollowing(l.ctx, in.UserId, in.QueryId); err != nil {
		return nil, errors.Wrapf(err, "DB error when check the relation beteween %d and %d", in.UserId, in.QueryId)
	} else {
		return &user.UserInfoReply{
			Id:            userFind.UserId,
			Name:          userFind.Name,
			FollowCount:   userFind.FollowCount,
			FollowerCount: userFind.FollowerCount,
			IsFollow:      isFollow,
		}, nil
	}

}
