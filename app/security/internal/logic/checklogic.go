package logic

import (
	"context"

	"go_code/Doul/app/security/internal/svc"
	"go_code/Doul/app/security/security"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *security.CheckLegaContentReq) (*security.CheckLegalContentResp, error) {
	_, _, found := l.svcCtx.Filter.Filter(in.Content)
	return &security.CheckLegalContentResp{
		Legal: !found,
	}, nil
}
