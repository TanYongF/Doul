package logic

import (
	"context"
	"go_code/Doul/app/usercenter/model"

	"go_code/Doul/app/usercenter/cmd/rpc/internal/svc"
	"go_code/Doul/app/usercenter/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowActionLogic {
	return &FollowActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowActionLogic) FollowAction(in *user.FollowActionReq) (*user.FollowActionResp, error) {
	var isDel int
	if in.Type == 1 {
		isDel = 0
	} else {
		isDel = 1
	}
	err := l.svcCtx.RelationModel.UpInsert(l.ctx, &model.DyRelation{
		FollowerId:  in.UserId,
		FollowingId: in.ToUserId,
		IsDel:       byte(isDel),
	})
	if err != nil {
		return nil, err
	}
	return &user.FollowActionResp{}, nil
}
