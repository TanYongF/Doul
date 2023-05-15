package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	Redis          redis.RedisConf
	DataSourceName string
}
