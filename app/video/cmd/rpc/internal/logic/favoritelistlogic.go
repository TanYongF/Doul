package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"
	"go_code/Doul/app/usercenter/cmd/rpc/userclient"
	"go_code/Doul/common/globalkey"
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

	// todo
	var resp = make([]*video.VideoPO, 0)
	for _, v := range videos {
		if v.UserId.Int64 == in.UserId {
			continue
		}
		var publishUser *userclient.UserInfoReply
		var isFavorite bool
		var favoriteCount int64
		err := mr.Finish(func() error {
			// Get User From UerRpc
			publishUser, err = l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.UserInfoReq{
				QueryId: v.UserId.Int64,
				UserId:  in.UserId,
			})
			return nil
		}, func() error {
			// Get Favorite Relation Between User and Video
			isFavorite, err = l.svcCtx.RedisClient.Sismember(globalkey.GetVideoLikesUsersRedisKey(v.VideoId), in.UserId)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "database error when check favorite relation between video and user")
			}
			if !isFavorite {
				isFavorite, err = l.svcCtx.DyFavoriteModel.CheckIsFavorite(l.ctx, in.UserId, v.VideoId)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "database error when check favorite relation between video and user")
				}
				if isFavorite {
					l.svcCtx.RedisClient.Sadd(globalkey.GetVideoLikesUsersRedisKey(v.VideoId), in.UserId)
				}
			}
			return nil
		}, func() error {
			// Get Favorite Count of this video
			favoriteCount, err = l.svcCtx.DyFavoriteModel.CountLikesByVideoId(l.ctx, v.VideoId)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "database error when count favorite nusmbers")
			}
			return nil
		})

		if err != nil {
			return nil, err
		}
		videoPo := &video.VideoPO{
			Author: &video.UserPO{
				FollowCount:   publishUser.FollowCount,
				FollowerCount: publishUser.FollowerCount,
				UserId:        publishUser.Id,
				IsFollow:      publishUser.IsFollow,
				Name:          publishUser.Name,
			},
			CommentCount:  v.CommentCount,
			CoverURL:      v.CoverUrl,
			FavoriteCount: favoriteCount,
			VideoId:       v.VideoId,
			IsFavorite:    isFavorite,
			PlayURL:       v.PlayUrl,
			Title:         v.Title.String,
		}
		resp = append(resp, videoPo)
	}

	return &video.FavoriteListResp{
		VideoList: resp,
	}, nil
}
