package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"go_code/Doul/common/rabbitmq"
)

type Config struct {
	zrpc.RpcServerConf
	UserRpcConf zrpc.RpcClientConf
	Mysql       struct { // 数据库配置，除mysql外，可能还有mongo等其他数据库
		DataSource string // mysql链接地址，满足 $user:$password@tcp($ip:$port)/$db?$queries 格式即可
	}
	CacheRedis     cache.CacheConf // redis缓存
	Redis          redis.RedisConf
	DataSourceName string
	RabbitMQ       rabbitmq.RabbitSenderConf
}
