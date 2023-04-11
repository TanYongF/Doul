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

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {

	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CommentAction Todo : 2023/4/11 Test
func (l *CommentActionLogic) CommentAction(req *types.CommentActionReq) (resp *types.CommentActionResp, err error) {

	// if action_type equal 1, insert this comment
	if req.ActionType == 1 {
		rpcResp, err := l.svcCtx.CommentRpc.CreateComment(l.ctx, &comment.PutCommentReq{
			VideoId:     req.VideoId,
			CommentText: req.CommentText,
			UserId:      tool.GetUidFromCtx(l.ctx),
		})
		if err != nil {
			return nil, err
		}

		videoUser, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.UserInfoReq{
			QueryId: tool.GetUidFromCtx(l.ctx),
			UserId:  tool.GetUidFromCtx(l.ctx),
		})
		if err != nil {
			return nil, err
		}

		var userPo types.User
		copier.Copy(&userPo, videoUser)

		return &types.CommentActionResp{
			Comment: types.Comment{
				Content:    rpcResp.Content,
				CreateDate: rpcResp.CreateAt,
				ID:         rpcResp.CommentId,
				User:       userPo,
			},
		}, nil
	} else {
		//else delete this comment
		_, err := l.svcCtx.CommentRpc.DeleteComment(l.ctx, &comment.DeleteCommentReq{
			VideoId:   req.VideoId,
			CommentId: req.CommentId,
		})
		if err != nil {
			return nil, err
		}
		return &types.CommentActionResp{}, nil
	}

}
