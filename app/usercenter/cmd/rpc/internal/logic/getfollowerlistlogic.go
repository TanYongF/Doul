package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go_code/Doul/app/usercenter/cmd/rpc/internal/svc"
	"go_code/Doul/app/usercenter/cmd/rpc/user"
	"go_code/Doul/app/usercenter/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerListLogic) GetFollowerList(in *user.GetFollowerListReq) (*user.GetFollowerListResp, error) {
	followerList, err := l.svcCtx.RelationModel.GetFollowerList(l.ctx, in.QueryId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(err, "DB error when GetFollowerListLogic")
	}
	var resp = make([]*user.UserPO, 0)
	for _, v := range followerList {
		var userPo user.UserPO
		copier.Copy(&userPo, v)

		// check the relation between A and B
		if isFollow, err := l.svcCtx.RelationModel.CheckFollowByFollowerAndFollowing(l.ctx, in.UserId, userPo.UserId); err != nil {
			return nil, errors.Wrapf(err, "DB error when check follow realtion")
		} else {
			userPo.IsFollow = isFollow
		}
		resp = append(resp, &userPo)
	}
	return &user.GetFollowerListResp{
		Users: resp,
	}, nil
}
