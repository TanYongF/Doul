package logic

import (
	"context"

	"go_code/Doul/app/comment/cmd/rpc/comment"
	"go_code/Doul/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutCommentLogic {
	return &PutCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutCommentLogic) PutComment(in *comment.PutCommentReq) (*comment.PutCommentResp, error) {
	// todo: add your logic here and delete this line

	return &comment.PutCommentResp{}, nil
}
