package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"go_code/Doul/common/globalkey"
	"strconv"
)

var (
	_ DyFavoriteModel = (*customDyFavoriteModel)(nil)
)

type (
	// DyFavoriteModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDyFavoriteModel.
	DyFavoriteModel interface {
		dyFavoriteModel
		CheckIsFavorite(ctx context.Context, userId int64, videoId int64) (bool, error)
		UpInsert(ctx context.Context, data *DyFavorite) error
		CountLikesByVideoId(ctx context.Context, videoId int64) (int64, error)
	}

	customDyFavoriteModel struct {
		*defaultDyFavoriteModel
		redisCli *redis.Redis
	}
)

func (c customDyFavoriteModel) CountLikesByVideoId(ctx context.Context, videoId int64) (int64, error) {
	var count int64
	var err error

	ret, err := c.redisCli.Hget(
		globalkey.GetVideoLikesCounterRedisKey(videoId),
		globalkey.GetVideoLikesCounterFieldKey(videoId))
	if stringx.NotEmpty(ret) {
		if err != nil {
			return 0, err
		}
		if count, err = strconv.ParseInt(ret, 10, 64); err != nil {
			return 0, err
		}
		return count, err
	}
	//查询加入缓存
	query := fmt.Sprintf("select count(*) from %s where video_id = ? and is_del = 0", c.table)
	err = c.QueryRowNoCache(&count, query, videoId)
	if err != nil {
		return 0, err
	}
	err = c.redisCli.Hset(
		globalkey.GetVideoLikesCounterRedisKey(videoId),
		globalkey.GetVideoLikesCounterFieldKey(videoId),
		strconv.FormatInt(count, 10))
	if err != nil {
		return 0, err
	}
	return count, nil
}

// UpInsert update if existed or insert
func (c customDyFavoriteModel) UpInsert(ctx context.Context, data *DyFavorite) error {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?) on duplicate key update `is_del` = ?", c.table, dyFavoriteRowsExpectAutoSet)
	_, err := c.ExecNoCacheCtx(ctx, query, data.UserId, data.VideoId, data.IsDel, data.IsDel)
	if err != nil {
		return err
	}
	return nil
}

func (c customDyFavoriteModel) CheckIsFavorite(ctx context.Context, userId int64, videoId int64) (bool, error) {
	// todo : 待完善具体存入逻辑
	//query := fmt.Sprintf("select 1 from %s where user_id = ? and video_id = ? and is_del = 0 limit 1", c.tableName())
	//var flag int
	//err := c.QueryRowNoCacheCtx(ctx, &flag, query, userId, videoId)
	//switch err {
	//case ErrNotFound:
	//	return false, nil
	//case nil:
	//	return flag == 1, nil
	//default:
	//	return false, err
	//}
	if userId == -1 {
		return false, nil
	}

	dyFavoriteUserIdVideoIdKey := fmt.Sprintf("%s%v:%v", cacheDyFavoriteUserIdVideoIdPrefix, userId, videoId)
	var resp int
	err := c.QueryRowIndexCtx(ctx, &resp, dyFavoriteUserIdVideoIdKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select 1 from %s where `user_id` = ? and `video_id` = ?  and `is_del` = 0 limit 1", c.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, videoId); err != nil {
			return 0, err
		}
		return 1, nil
	}, c.queryPrimary)
	switch err {
	case nil:
		return resp == 1, nil
	case sqlc.ErrNotFound:
		return false, nil
	default:
		return false, err
	}
}

// NewDyFavoriteModel returns a model for the database table.
func NewDyFavoriteModel(conn sqlx.SqlConn, c cache.CacheConf, cli *redis.Redis) DyFavoriteModel {
	return &customDyFavoriteModel{
		defaultDyFavoriteModel: newDyFavoriteModel(conn, c),
		redisCli:               cli,
	}
}
