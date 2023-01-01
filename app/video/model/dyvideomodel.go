package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DyVideoModel = (*customDyVideoModel)(nil)

type (
	// DyVideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDyVideoModel.
	DyVideoModel interface {
		dyVideoModel
	}

	customDyVideoModel struct {
		*defaultDyVideoModel
	}
)

// NewDyVideoModel returns a model for the database table.
func NewDyVideoModel(conn sqlx.SqlConn, c cache.CacheConf) DyVideoModel {
	return &customDyVideoModel{
		defaultDyVideoModel: newDyVideoModel(conn, c),
	}
}
