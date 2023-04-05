package logic

import (
	"context"
	"go_code/Doul/app/video/cmd/rpc/video"
	"go_code/Doul/common/tool"

	"go_code/Doul/app/video/cmd/api/internal/svc"
	"go_code/Doul/app/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteLogic {
	return &FavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteLogic) Favorite(req *types.FavoriteReq) (resp *types.FavoriteResp, err error) {
	_, err = l.svcCtx.VideoRpc.Like(l.ctx, &video.LikeReq{
		UserId:  tool.GetUidFromCtx(l.ctx),
		VideoId: req.VideoId,
		Type:    req.ActionType == 1,
	})
	if err != nil {
		return nil, err
	}
	return &types.FavoriteResp{}, nil
}
