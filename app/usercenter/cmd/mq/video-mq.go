package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"go_code/Doul/app/usercenter/cmd/mq/internal/config"
	"go_code/Doul/app/usercenter/cmd/mq/internal/listen"
)

var configFile = flag.String("f", "etc/video-mq.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range listen.Mqs(c) {
		serviceGroup.Add(mq)
	}
	logx.Info("Start successful")
	serviceGroup.Start()
}
