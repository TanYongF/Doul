package logic

import (
	"context"
	"github.com/pkg/errors"
	"go_code/Doul/common/xerr"

	"go_code/Doul/app/comment/cmd/rpc/comment"
	"go_code/Doul/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentLogic) DeleteComment(in *comment.DeleteCommentReq) (*comment.DeleteCommentResp, error) {
	err := l.svcCtx.DyCommentModel.DeleteById(in.CommentId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "DB error when delete comments %d", in.CommentId)
	}

	return &comment.DeleteCommentResp{}, nil
}
