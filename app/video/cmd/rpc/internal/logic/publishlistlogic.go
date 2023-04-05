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

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *video.PublishListReq) (*video.PublishListResp, error) {
	videos, err := l.svcCtx.DyVideoModel.GetPublishListByUserId(l.ctx, in.QueryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Get Publish List error")
	}

	var resp = make([]*video.VideoPO, 0)
	for _, v := range videos {
		// Get user info
		var publishUser *userclient.UserInfoReply
		var isFavorite bool
		var favoriteCount int64
		err := mr.Finish(func() error {
			// Get User From UerRpc
			if publishUser, err = l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.UserInfoReq{
				QueryId: v.UserId.Int64,
				UserId:  in.UserId,
			}); err != nil {
				return err
			}
			return nil
		}, func() error {
			// Get Favorite Relation Between User and Video
			if isFavorite, err = l.svcCtx.RedisClient.Sismember(globalkey.GetVideoLikesUsersRedisKey(v.VideoId), in.UserId); err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "database error when check favorite relation between video and user")
			}
			if !isFavorite {
				if isFavorite, err = l.svcCtx.DyFavoriteModel.CheckIsFavorite(l.ctx, in.UserId, v.VideoId); err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "database error when check favorite relation between video and user")
				}
				if isFavorite {
					l.svcCtx.RedisClient.Sadd(globalkey.GetVideoLikesUsersRedisKey(v.VideoId), in.UserId)
				}
			}
			return nil
		}, func() error {
			// Get Favorite Count of this video
			if favoriteCount, err = l.svcCtx.DyFavoriteModel.CountLikesByVideoId(l.ctx, v.VideoId); err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "database error when count favorite nusmbers")
			}
			return nil
		})

		if err != nil {
			return nil, err
		}
		// type invert
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

	return &video.PublishListResp{
		VideoList: resp,
	}, nil
}
