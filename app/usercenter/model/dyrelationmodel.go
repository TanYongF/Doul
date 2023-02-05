package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
)

var (
	_                            DyRelationModel = (*customDyRelationModel)(nil)
	cacheFollowerCountPrefix                     = "cache:douyin:dyRelation:followerCount:"
	cacheFollowingCountPrefix                    = "cache:douyin:dyRelation:followingCount:"
	cacheRelationPrefix                          = "cache:douyin:dyRelation:"
	cacheRelationFollowingSuffix                 = ":subscribe"
	cacheRelationFansSuffix                      = ":fans"
)

type (
	// DyRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDyRelationModel.
	DyRelationModel interface {
		dyRelationModel
		CheckFollowByFollowerAndFollowing(ctx context.Context, followerId int64, followingId int64) (bool, error)
		CountFollowersByUserId(ctx context.Context, userId int64) (*int64, error)
		CountFollowingsByUserId(ctx context.Context, userId int64) (*int64, error)
		GetFollowerList(ctx context.Context, userId int64) ([]DyUser, error)
		GetFollowingList(ctx context.Context, userId int64) ([]DyUser, error)
	}

	customDyRelationModel struct {
		*defaultDyRelationModel
		rdc *redis.Redis
	}
)

// GetFollowerList 获取粉丝列表
func (c customDyRelationModel) GetFollowerList(ctx context.Context, userId int64) ([]DyUser, error) {
	query := fmt.Sprintf("select * from dy_user du where du.user_id  in (select dr.follower_id from dy_relation dr where dr.following_id = ?)")
	var followers []DyUser
	err := c.QueryRowsNoCacheCtx(ctx, &followers, query, userId)
	switch err {
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	case nil:
		return followers, nil
	default:
		return nil, err
	}
}

// GetFollowingList 获取关注列表
func (c customDyRelationModel) GetFollowingList(ctx context.Context, userId int64) ([]DyUser, error) {
	query := fmt.Sprintf("select * from dy_user du where du.user_id  in (select dr.following_id from dy_relation dr where dr.follower_id = ?)")
	var followings []DyUser
	err := c.QueryRowsNoCacheCtx(ctx, &followings, query, userId)
	logx.Info(followings)
	switch err {
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	case nil:
		return followings, nil
	default:
		return nil, err
	}
}

// CountFollowersByUserId 获取关注人数
func (c customDyRelationModel) CountFollowersByUserId(ctx context.Context, userId int64) (*int64, error) {

	var followerCount int64
	err := c.QueryRowCtx(ctx, &followerCount, cacheFollowerCountPrefix+string(userId), func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select count(*) from %s dr where follower_id  = ?", c.table)
		return conn.QueryRowCtx(ctx, v, query, userId)
	})
	if err != nil {
		return nil, err
	}
	return &followerCount, err
}

// CountFollowingsByUserId  获取粉丝列表
func (c customDyRelationModel) CountFollowingsByUserId(ctx context.Context, userId int64) (*int64, error) {
	var followingCount int64
	err := c.QueryRowCtx(ctx, &followingCount, cacheFollowerCountPrefix+string(userId), func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select count(*) from %s dr where following_id  = ?", c.table)
		return conn.QueryRowCtx(ctx, v, query, userId)
	})
	if err != nil {
		return nil, err
	}
	return &followingCount, err
}

// CheckFollowByFollowerAndFollowing 获取两者关注关系 follower_id : 博主ID following_id : 粉丝ID
func (c customDyRelationModel) CheckFollowByFollowerAndFollowing(ctx context.Context, followerId int64, followingId int64) (bool, error) {
	var isFollow bool

	key := cacheRelationPrefix + strconv.FormatInt(followerId, 10) + cacheRelationFollowingSuffix
	exists, _ := c.rdc.Exists(key)
	logx.Error(key)
	if exists {
		isFollow, _ = c.rdc.Sismember(key, followingId)
	} else {
		subscribes, _ := c.GetFollowingList(ctx, followerId)
		for _, v := range subscribes {
			c.rdc.Sadd(key, v.UserId)
		}
	}
	isFollow, _ = c.rdc.Sismember(key, followingId)
	return isFollow, nil

	//query := fmt.Sprintf("select exists(select * from %s dr where `follower_id` = ? and `following_id` = ? and is_del = 0 limit 1)", c.table)
	//err := c.QueryRowNoCacheCtx(ctx, &isFollow, query, followerId, followingId)
	//resp := isFollow == 1
	//switch err {
	//case nil:
	//	return resp, nil
	//default:
	//	return false, err
	//}
}

// NewDyRelationModel returns a model for the database table.
func NewDyRelationModel(conn sqlx.SqlConn, c cache.CacheConf, redisClient *redis.Redis) *customDyRelationModel {
	return &customDyRelationModel{
		defaultDyRelationModel: newDyRelationModel(conn, c),
		rdc:                    redisClient,
	}
}
