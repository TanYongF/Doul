package logic

import (
	"context"
	"github.com/pkg/errors"
	"go_code/Doul/common/xerr"

	"go_code/Doul/app/usercenter/cmd/rpc/internal/svc"
	"go_code/Doul/app/usercenter/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerCountLogic {
	return &GetFollowerCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerCountLogic) GetFollowerCount(in *user.GetFollowerCountReq) (*user.GetFollowerCountResp, error) {
	count, err := l.svcCtx.RelationModel.CountFollowersByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "DB error when count folloer numbers")
	}
	return &user.GetFollowerCountResp{
		Count: *count,
	}, nil
}
