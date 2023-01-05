package logic

import (
	"context"

	"go_code/Doul/app/comment/cmd/rpc/comment"
	"go_code/Doul/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentListLogic) GetCommentList(in *comment.CommentListReq) (*comment.CommentListResp, error) {
	// todo: add your logic here and delete this line

	return &comment.CommentListResp{}, nil
}
