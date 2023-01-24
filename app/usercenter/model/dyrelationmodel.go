package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	_                         DyRelationModel = (*customDyRelationModel)(nil)
	cacheFollowerCountPrefix                  = "cache:dyRelation:followerCount:"
	cacheFollowingCountPrefix                 = "cache:dyRelation:followingCount:"
)

type (
	// DyRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDyRelationModel.
	DyRelationModel interface {
		dyRelationModel
		CheckFollowByFollowerAndFollowing(ctx context.Context, followerId int64, followingId int64) (*bool, error)
		CountFollowersByUserId(ctx context.Context, userId int64) (*int64, error)
		CountFollowingsByUserId(ctx context.Context, userId int64) (*int64, error)
	}

	customDyRelationModel struct {
		*defaultDyRelationModel
	}
)

// CountFollowersByUserId 获取关注列表
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

// CheckFollowByFollowerAndFollowing  follower_id : 博主ID following_id : 粉丝ID
func (c customDyRelationModel) CheckFollowByFollowerAndFollowing(ctx context.Context, followerId int64, followingId int64) (*bool, error) {
	var isFollow int
	query := fmt.Sprintf("select exists(select * from %s dr where `follower_id` = ? and `following_id` = ? and is_del = 0 limit 1)", c.table)
	err := c.QueryRowNoCacheCtx(ctx, &isFollow, query, followerId, followingId)
	resp := isFollow == 1
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// NewDyRelationModel returns a model for the database table.
func NewDyRelationModel(conn sqlx.SqlConn, c cache.CacheConf) *customDyRelationModel {
	return &customDyRelationModel{
		defaultDyRelationModel: newDyRelationModel(conn, c),
	}
}
