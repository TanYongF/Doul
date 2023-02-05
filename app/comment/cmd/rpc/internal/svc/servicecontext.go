package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go_code/Doul/app/comment/cmd/rpc/internal/config"
	"go_code/Doul/app/comment/model"
	"go_code/Doul/app/usercenter/cmd/rpc/userclient"
)

type ServiceContext struct {
	Config         config.Config
	UserRpc        userclient.User
	DyCommentModel model.DyCommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		DyCommentModel: model.NewDyCommentModel(conn, c.CacheRedis),
	}
}
