package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_code/Doul/app/video/cmd/mq/internal/config"
	"go_code/Doul/app/video/model"
)

type ServiceContext struct {
	Config          config.Config
	DyVideoModel    model.DyVideoModel
	DyFavoriteModel model.DyFavoriteModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//加载Mysql
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:          c,
		DyVideoModel:    model.NewDyVideoModel(sqlConn, c.CacheRedis),
		DyFavoriteModel: model.NewDyFavoriteModel(sqlConn, c.CacheRedis, nil),
	}

}
