package logic

import (
	"context"

	"go_code/Doul/app/comment/cmd/api/internal/svc"
	"go_code/Doul/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentListLogic) GetCommentList(req *types.CommentListReq) (resp *types.CommentListResp, err error) {
	// todo: add your logic here and delete this line

	return &types.CommentListResp{
		CommentList: nil,
		Response: types.Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
	}, nil
}
