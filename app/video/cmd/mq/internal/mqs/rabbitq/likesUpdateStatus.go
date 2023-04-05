package rabbitq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go_code/Doul/app/video/cmd/mq/internal/svc"
	"go_code/Doul/app/video/model"
	"go_code/Doul/common/rabbitmq"
	"go_code/Doul/common/xerr"
)

type LikesUpdateStatus struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentUpdateStatusMq(ctx context.Context, svcCtx *svc.ServiceContext) *LikesUpdateStatus {
	return &LikesUpdateStatus{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikesUpdateStatus) Consume(val string) error {

	var message rabbitmq.LikesRelationUpdateStockMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("LikesUpdateStatus->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("LikesUpdateStatus->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

func (l *LikesUpdateStatus) execService(message rabbitmq.LikesRelationUpdateStockMessage) error {
	fmt.Printf("%d %d %d \n", message.Type, message.VideoId, message.UserId)
	//入库
	var isDel int
	if message.Type {
		isDel = 0
	} else {
		isDel = 1
	}
	err := l.svcCtx.DyFavoriteModel.UpInsert(l.ctx, &model.DyFavorite{
		UserId:  message.UserId,
		VideoId: message.VideoId,
		IsDel:   byte(isDel),
	})

	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insert like relation db fail err : %v ,message:%+v", err, message)
	}
	return nil
}
