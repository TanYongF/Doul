package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_code/Doul/app/usercenter/cmd/rpc/internal/config"
	"go_code/Doul/app/usercenter/model"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.DyUserModel
	RelationModel model.DyRelationModel
	RedisClient   *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	redisClient := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Pass = c.Redis.Pass
	})
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewDyUserModel(conn, c.CacheRedis),
		RelationModel: model.NewDyRelationModel(conn, c.CacheRedis, redisClient),
		RedisClient:   redisClient,
	}
}
