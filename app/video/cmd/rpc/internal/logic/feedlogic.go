package logic

import (
	"context"
	"go_code/Doul/app/video/cmd/rpc/internal/svc"
	"go_code/Doul/app/video/cmd/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedLogic) Feed(in *video.FeedReq) (*video.FeedResp, error) {
	// todo: 还有限制返回时间戳未完成
	videos, err := l.svcCtx.DyVideoModel.GetVideoList(l.ctx)
	if err != nil {
		return nil, err
	}
	var videoList = make([]*video.VideoPO, 0)
	for _, v := range videos {
		videoList = append(
			videoList,
			&video.VideoPO{
				Author: &video.UserPO{
					FollowCount:   v.FollowCount,
					FollowerCount: v.FollowerCount,
					UserId:        v.UserId.Int64,
					IsFollow:      v.IsFollow == 1,
					Name:          v.Name,
				},
				CommentCount:  v.CommentCount,
				CoverURL:      v.CoverUrl,
				FavoriteCount: v.FavoriteCount,
				VideoId:       v.VideoId,
				IsFavorite:    false, //todo
				PlayURL:       v.PlayUrl,
				Title:         v.Title.String,
			},
		)
	}
	//todo nextTime 待完善
	return &video.FeedResp{
		NextTime:  "",
		VideoList: videoList,
	}, nil
}
