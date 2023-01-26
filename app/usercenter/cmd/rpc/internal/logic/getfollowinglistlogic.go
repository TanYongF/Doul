package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go_code/Doul/app/usercenter/model"

	"go_code/Doul/app/usercenter/cmd/rpc/internal/svc"
	"go_code/Doul/app/usercenter/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingListLogic {
	return &GetFollowingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowingListLogic) GetFollowingList(in *user.GetFollowingListReq) (*user.GetFollowingListResp, error) {
	followingList, err := l.svcCtx.RelationModel.GetFollowingList(l.ctx, in.GetQueryId())
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(err, "DB error when get following list")
	}

	var resp = make([]*user.UserPO, 0)
	for _, v := range followingList {
		var userPo user.UserPO
		copier.Copy(&userPo, v)
		//check the relation between A and B
		if isFollow, err := l.svcCtx.RelationModel.CheckFollowByFollowerAndFollowing(l.ctx, in.UserId, userPo.UserId); err != nil {
			return nil, errors.Wrapf(err, "DB error when check follow realtion")
		} else {
			userPo.IsFollow = isFollow
		}
		resp = append(resp, &userPo)
	}
	return &user.GetFollowingListResp{
		Users: resp,
	}, nil
}
