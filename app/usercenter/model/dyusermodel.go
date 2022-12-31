package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DyUserModel = (*customDyUserModel)(nil)

type (
	// DyUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDyUserModel.
	DyUserModel interface {
		dyUserModel
		FindOneByUsername(ctx context.Context, username string) (*DyUser, error)
	}

	customDyUserModel struct {
		*defaultDyUserModel
	}
)

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
