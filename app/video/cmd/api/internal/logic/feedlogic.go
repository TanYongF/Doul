package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go_code/Doul/app/video/cmd/rpc/videoclient"
	"time"

	"go_code/Doul/app/video/cmd/api/internal/svc"
	"go_code/Doul/app/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedReq) (resp *types.FeedResp, err error) {
	feeds, err := l.svcCtx.VideoRpc.Feed(l.ctx, &videoclient.FeedReq{
		LatestTime: req.LastestTime,
	})
	if err != nil {
		return nil, err
	}
	videos := make([]types.Video, len(feeds.VideoList))

	for i := 0; i < len(feeds.VideoList); i++ {
		copier.Copy(&videos[i], &feeds.VideoList[i])
		copier.Copy(&videos[i].Author, &feeds.VideoList[i].Author)
	}

	var nextTime = int64(time.Now().Second())
	return &types.FeedResp{
		NextTime:  &nextTime,
		VideoList: videos,
	}, nil
}
