package logic

import (
	"context"
	"github.com/pkg/errors"
	"go_code/Doul/app/usercenter/cmd/rpc/userclient"
	"go_code/Doul/common/tool"
	"go_code/Doul/common/xerr"

	"go_code/Doul/app/video/cmd/rpc/internal/svc"
	"go_code/Doul/app/video/cmd/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *video.FavoriteListReq) (*video.FavoriteListResp, error) {
	videos, err := l.svcCtx.DyVideoModel.GetFavoriteListByUserId(l.ctx, in.QueryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Get Favorite List error")
	}

	// todo Author 和 isFavorite 待填充
	var resp = make([]*video.VideoPO, 0)
	for _, v := range videos {
		user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.UserInfoReq{
			QueryId: in.QueryId,
			UserId:  tool.GetUidFromCtx(l.ctx),
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get user %d info error", v.UserId.Int64)
		}
		videoPo := &video.VideoPO{
			Author: &video.UserPO{
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				UserId:        user.Id,
				IsFollow:      user.IsFollow,
				Name:          user.Name,
			},
			CommentCount:  v.CommentCount,
			CoverURL:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			VideoId:       v.VideoId,
			IsFavorite:    false, //todo 待填充
			PlayURL:       v.PlayUrl,
			Title:         v.Title.String,
		}
		resp = append(resp, videoPo)
	}

	return &video.FavoriteListResp{
		VideoList: resp,
	}, nil
}
