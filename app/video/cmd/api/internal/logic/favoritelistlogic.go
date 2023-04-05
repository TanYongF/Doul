package logic

import (
	"context"
	"go_code/Doul/app/video/cmd/rpc/video"
	"go_code/Doul/common/tool"

	"strconv"

	"go_code/Doul/app/video/cmd/api/internal/svc"
	"go_code/Doul/app/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.FavoriteVideofListReq) (resp *types.FavoriteVideoListResp, err error) {
	queryId, _ := strconv.ParseInt(req.UserId, 10, 64)
	favoriteList, err := l.svcCtx.VideoRpc.FavoriteList(l.ctx, &video.FavoriteListReq{
		QueryId: queryId,
		UserId:  tool.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}
	videos := make([]types.Video, 0)
	for _, v := range favoriteList.VideoList {
		vo := types.Video{
			Author: types.User{
				FollowCount:   v.Author.FollowCount,
				FollowerCount: v.Author.FollowerCount,
				UserId:        v.Author.UserId,
				IsFollow:      v.Author.IsFollow,
				Name:          v.Author.Name,
			},
			CommentCount:  v.CommentCount,
			CoverURL:      v.CoverURL,
			FavoriteCount: v.FavoriteCount,
			VideoId:       v.VideoId,
			IsFavorite:    v.IsFavorite,
			PlayURL:       v.PlayURL,
			Title:         v.Title,
		}
		videos = append(videos, vo)
	}
	return &types.FavoriteVideoListResp{
		VideoList: videos,
	}, nil
}
