package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DyRelationModel = (*customDyRelationModel)(nil)

type (
	// DyRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDyRelationModel.
	DyRelationModel interface {
		dyRelationModel
	}

	customDyRelationModel struct {
		*defaultDyRelationModel
	}
)

// NewDyRelationModel returns a model for the database table.
func NewDyRelationModel(conn sqlx.SqlConn, c cache.CacheConf) DyRelationModel {
	return &customDyRelationModel{
		defaultDyRelationModel: newDyRelationModel(conn, c),
	}
}
