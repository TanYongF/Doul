package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go_code/Doul/app/usercenter/cmd/api/internal/config"
	"go_code/Doul/app/usercenter/cmd/api/internal/middleware"
	"go_code/Doul/app/usercenter/cmd/rpc/userclient"
)

type ServiceContext struct {
	Config         config.Config
	UserRpc        userclient.User
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	userClient := userclient.NewUser(zrpc.MustNewClient(c.UserRpc))
	return &ServiceContext{
		Config:         c,
		UserRpc:        userClient,
		AuthMiddleware: middleware.NewAuthMiddleware(userClient).Handle,
	}
}
