package logic

import (
	"context"
	"github.com/pkg/errors"
	"go_code/Doul/common/xerr"

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
	comments, err := l.svcCtx.DyCommentModel.FindByVideoId(in.VideoId, 3, 1)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "DB error when get comments")
	}
	var resp = make([]*comment.CommentBody, 0)
	for _, c := range comments {
		commentBody := comment.CommentBody{
			VideoId:  c.VideoId,
			Content:  c.Content,
			UserId:   c.UserId,
			CreateAt: c.CreatedAt.Format("01-02"),
		}
		resp = append(resp, &commentBody)
	}

	return &comment.CommentListResp{
		Comments: resp,
	}, nil
}
