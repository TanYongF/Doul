package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go_code/Doul/app/comment/cmd/api/internal/svc"
	"go_code/Doul/app/comment/cmd/api/internal/types"
	"go_code/Doul/app/comment/cmd/rpc/comment"
	"go_code/Doul/app/usercenter/cmd/rpc/user"
	"go_code/Doul/common/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListReq) (resp *types.CommentListResp, err error) {
	rpcResult, err := l.svcCtx.CommentRpc.GetCommentList(l.ctx, &comment.CommentListReq{VideoId: req.VideoId})
	if err != nil {
		return nil, err
	}

	commentList := make([]types.Comment, len(rpcResult.Comments))
	for i := 0; i < len(rpcResult.Comments); i++ {
		userFromDB, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.UserInfoReq{
			QueryId: rpcResult.Comments[i].UserId,
			UserId:  tool.GetUidFromCtx(l.ctx),
		})
		if err != nil {
			return nil, err
		}
		err = copier.Copy(&commentList[i].User, userFromDB)
		if err != nil {
			return nil, err
		}
		commentList[i].Content = rpcResult.Comments[i].Content
		commentList[i].CreateDate = rpcResult.Comments[i].CreateAt
	}
	return &types.CommentListResp{
		CommentList: commentList,
	}, nil
}
