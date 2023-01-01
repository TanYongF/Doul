package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DyCommentModel = (*customDyCommentModel)(nil)

type (
	// DyCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDyCommentModel.
	DyCommentModel interface {
		dyCommentModel
	}

	customDyCommentModel struct {
		*defaultDyCommentModel
	}
)

// NewDyCommentModel returns a model for the database table.
func NewDyCommentModel(conn sqlx.SqlConn, c cache.CacheConf) DyCommentModel {
	return &customDyCommentModel{
		defaultDyCommentModel: newDyCommentModel(conn, c),
	}
}
