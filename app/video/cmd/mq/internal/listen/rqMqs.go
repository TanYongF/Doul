package listen

import (
	"context"
	"github.com/zeromicro/go-zero/core/service"
	"go_code/Doul/app/video/cmd/mq/internal/config"
	"go_code/Doul/app/video/cmd/mq/internal/mqs/rabbitq"
	"go_code/Doul/app/video/cmd/mq/internal/svc"
	"go_code/Doul/common/rabbitmq"
)

// pub sub use rq (rabbitmq)
func RabbitMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		rabbitmq.MustNewListener(c.RabbitMQConfig, rabbitq.NewPaymentUpdateStatusMq(ctx, svcContext)),

		//.....
	}

}
