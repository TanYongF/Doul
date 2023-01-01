package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DyFavoriteModel = (*customDyFavoriteModel)(nil)

type (
	// DyFavoriteModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDyFavoriteModel.
	DyFavoriteModel interface {
		dyFavoriteModel
	}

	customDyFavoriteModel struct {
		*defaultDyFavoriteModel
	}
)

// NewDyFavoriteModel returns a model for the database table.
func NewDyFavoriteModel(conn sqlx.SqlConn, c cache.CacheConf) DyFavoriteModel {
	return &customDyFavoriteModel{
		defaultDyFavoriteModel: newDyFavoriteModel(conn, c),
	}
}
