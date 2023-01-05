package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go_code/Doul/app/comment/cmd/api/internal/config"
	"go_code/Doul/app/comment/cmd/rpc/comment"
	"go_code/Doul/app/comment/cmd/rpc/commentclient"
	"go_code/Doul/app/usercenter/cmd/rpc/user"
	"go_code/Doul/app/usercenter/cmd/rpc/userclient"
	"go_code/Doul/common/middleware"
)

type ServiceContext struct {
	Config         config.Config
	UserRpc        user.UserClient
	CommentRpc     comment.CommentClient
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	userClient := userclient.NewUser(zrpc.MustNewClient(c.UserRpc))
	return &ServiceContext{
		Config:         c,
		UserRpc:        userClient,
		CommentRpc:     commentclient.NewComment(zrpc.MustNewClient(c.CommentRpc)),
		AuthMiddleware: middleware.NewAuthMiddleware(userClient).Handle,
	}
}
