package main

import (
	"flag"
	"fmt"
	"go_code/Doul/common/interceptor"

	"go_code/Doul/app/comment/cmd/rpc/comment"
	"go_code/Doul/app/comment/cmd/rpc/internal/config"
	"go_code/Doul/app/comment/cmd/rpc/internal/server"
	"go_code/Doul/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/comment.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		comment.RegisterCommentServer(grpcServer, server.NewCommentServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//add log interceptor
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
