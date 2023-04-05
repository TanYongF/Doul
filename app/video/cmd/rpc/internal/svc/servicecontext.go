package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go_code/Doul/app/usercenter/cmd/rpc/userclient"
	"go_code/Doul/app/video/cmd/rpc/internal/config"
	"go_code/Doul/app/video/model"
	"go_code/Doul/common/rabbitmq"
)

type ServiceContext struct {
	Config          config.Config
	DyVideoModel    model.DyVideoModel
	DyFavoriteModel model.DyFavoriteModel
	RedisClient     *redis.Redis
	UserRpc         userclient.User
	MqSender        rabbitmq.Sender
}

func NewServiceContext(c config.Config) *ServiceContext {
	//加载Mysql
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	//加载Redis Cli
	redisCli := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Pass = c.Redis.Pass
	})
	return &ServiceContext{
		Config:          c,
		DyVideoModel:    model.NewDyVideoModel(sqlConn, c.CacheRedis),
		DyFavoriteModel: model.NewDyFavoriteModel(sqlConn, c.CacheRedis, redisCli),
		RedisClient:     redisCli,
		UserRpc:         userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		MqSender:        rabbitmq.MustNewSender(c.RabbitMQ),
	}

}
