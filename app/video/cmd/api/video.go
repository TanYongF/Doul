package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"go_code/Doul/app/video/cmd/api/internal/config"
	"go_code/Doul/app/video/cmd/api/internal/handler"
	"go_code/Doul/app/video/cmd/api/internal/svc"
)

var configFile = flag.String("f", "etc/video-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//// todo 循环依赖问题
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
