package logic

import (
	"context"
	"go_code/Doul/app/comment/model"
	"go_code/Doul/common/tool"

	"go_code/Doul/app/comment/cmd/rpc/comment"
	"go_code/Doul/app/comment/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommentLogic) CreateComment(in *comment.PutCommentReq) (*comment.PutCommentResp, error) {
	var commentToInsert = model.DyComment{
		UserId:  tool.GetUidFromCtx(l.ctx),
		Content: in.GetCommentText(),
		IsDel:   0,
		VideoId: in.GetVideoId(),
	}
	_, err := l.svcCtx.DyCommentModel.Insert(l.ctx, &commentToInsert)
	if err != nil {
		return nil, err
	}

	// Todo : 2023/4/11 To solve the time format
	return &comment.PutCommentResp{
		CommentId: commentToInsert.CommentId,
		Content:   commentToInsert.Content,
		CreateAt:  commentToInsert.CreatedAt.String(),
	}, nil
}
