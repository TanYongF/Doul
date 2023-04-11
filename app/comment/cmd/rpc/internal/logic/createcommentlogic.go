package logic

import (
	"context"
	"go_code/Doul/app/comment/cmd/rpc/comment"
	"go_code/Doul/app/comment/cmd/rpc/internal/svc"
	"go_code/Doul/app/comment/model"

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
		UserId:  in.UserId,
		Content: in.GetCommentText(),
		IsDel:   0,
		VideoId: in.GetVideoId(),
	}
	if _, err := l.svcCtx.DyCommentModel.Insert(l.ctx, &commentToInsert); err != nil {
		return nil, err
	}

	// TODO : 2023/4/11 To solve the time format
	return &comment.PutCommentResp{
		CommentId: commentToInsert.CommentId,
		Content:   commentToInsert.Content,
		CreateAt:  commentToInsert.CreatedAt.String(),
	}, nil
}
