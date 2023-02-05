package logic

import (
	"context"
	"go_code/Doul/app/comment/model"

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
	// todo 这里应该从ctx取得 user_id
	l.svcCtx.DyCommentModel.Insert(l.ctx, &model.DyComment{
		UserId:  0,
		Content: in.GetCommentText(),
		IsDel:   0,
		VideoId: in.GetVideoId(),
	})

	return &comment.PutCommentResp{}, nil
}
