package rabbitq

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go_code/Doul/app/usercenter/cmd/mq/internal/svc"
	model2 "go_code/Doul/app/usercenter/model"
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

	var message rabbitmq.FollowRelationUpdateMessage
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

func (l *LikesUpdateStatus) execService(message rabbitmq.FollowRelationUpdateMessage) error {
	//入库
	var isDel byte
	if message.Type {
		isDel = 0
	} else {
		isDel = 1
	}
	err := l.svcCtx.DyRelationModel.UpInsert(l.ctx, &model2.DyRelation{
		FollowerId:  message.FollowerId,
		FollowingId: message.FollowingId,
		IsDel:       isDel,
	})

	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insert like relation db fail err : %v ,message:%+v", err, message)
	}
	return nil
}
