package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_code/Doul/common/globalkey"
)

var _ DyUserModel = (*customDyUserModel)(nil)

type (
	DyUserModelMasking struct {
		UserId        int64  `db:"user_id"`        // 用户ID
		Name          string `db:"name"`           // 用户名
		FollowerCount int64  `db:"follower_count"` // 粉丝总数
		IsFollow      int64  `db:"is_follow"`      // 是否已关注
		FollowCount   int64  `db:"follow_count"`   // 关注总数
	}
	// DyUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDyUserModel.
	DyUserModel interface {
		dyUserModel
		FindOneByUsername(ctx context.Context, username string) (*DyUser, error)
		FindOneByUserId(ctx context.Context, userId int64) (*DyUser, error)
	}

	customDyUserModel struct {
		*defaultDyUserModel
	}
)

func (c customDyUserModel) FindOneByUserId(ctx context.Context, userId int64) (*DyUser, error) {
	douyinDyUserUserIdKey := globalkey.GetUserById(userId)
	var resp DyUser
	err := c.QueryRowCtx(ctx, &resp, douyinDyUserUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", dyUserRows, c.table)
		return conn.QueryRowCtx(ctx, v, query, userId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindOneByUsername 通过用户名查找用户
func (c customDyUserModel) FindOneByUsername(ctx context.Context, username string) (*DyUser, error) {
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", dyUserRows, c.table)
	var resp DyUser
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		logx.Errorf("FindOneByUsername User Position Model Model err, err=%v", err)
		return nil, err
	}
}

// NewDyUserModel returns a model for the database table.
func NewDyUserModel(conn sqlx.SqlConn, c cache.CacheConf) DyUserModel {
	return &customDyUserModel{
		defaultDyUserModel: newDyUserModel(conn, c),
	}
}
