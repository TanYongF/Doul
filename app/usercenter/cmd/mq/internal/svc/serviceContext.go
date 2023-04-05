package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_code/Doul/app/usercenter/cmd/mq/internal/config"
	"go_code/Doul/app/usercenter/model"
)

type ServiceContext struct {
	Config          config.Config
	DyRelationModel model.DyRelationModel
	RedisClient     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	//加载Mysql
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	redisClient := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Pass = c.Redis.Pass
	})
	return &ServiceContext{
		Config:          c,
		DyRelationModel: model.NewDyRelationModel(sqlConn, c.CacheRedis, redisClient),
	}

}
