package listen

import (
	"context"
	"github.com/zeromicro/go-zero/core/service"
	"go_code/Doul/app/usercenter/cmd/mq/internal/config"
	"go_code/Doul/app/usercenter/cmd/mq/internal/svc"
)

// back to all consumers
func Mqs(c config.Config) []service.Service {

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	var services []service.Service

	//kq ：pub sub
	services = append(services, RabbitMqs(c, ctx, svcContext)...)

	return services
}
