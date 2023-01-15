package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go_code/Doul/app/usercenter/cmd/rpc/userclient"
	"go_code/Doul/app/video/cmd/rpc/internal/config"
	"go_code/Doul/app/video/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config       config.Config
	DyVideoModel model.DyVideoModel
	RedisClient  *redis.Redis
	DbEngin      *gorm.DB
	UserRpc      userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	//启动Gorm支持
	db, err := gorm.Open(mysql.Open(c.DataSourceName), &gorm.Config{})
	//如果出错就GameOver了
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:       c,
		DyVideoModel: model.NewDyVideoModel(conn, c.CacheRedis),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Pass = c.Redis.Pass
		}),
		DbEngin: db,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}

}
