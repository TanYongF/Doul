// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package server

import (
	"context"

	"go_code/Doul/app/comment/cmd/rpc/comment"
	"go_code/Doul/app/comment/cmd/rpc/internal/logic"
	"go_code/Doul/app/comment/cmd/rpc/internal/svc"
)

type CommentServer struct {
	svcCtx *svc.ServiceContext
	comment.UnimplementedCommentServer
}

func NewCommentServer(svcCtx *svc.ServiceContext) *CommentServer {
	return &CommentServer{
		svcCtx: svcCtx,
	}
}

func (s *CommentServer) GetCommentList(ctx context.Context, in *comment.CommentListReq) (*comment.CommentListResp, error) {
	l := logic.NewGetCommentListLogic(ctx, s.svcCtx)
	return l.GetCommentList(in)
}

func (s *CommentServer) CreateComment(ctx context.Context, in *comment.PutCommentReq) (*comment.PutCommentResp, error) {
	l := logic.NewCreateCommentLogic(ctx, s.svcCtx)
	return l.CreateComment(in)
}

func (s *CommentServer) DeleteComment(ctx context.Context, in *comment.DeleteCommentReq) (*comment.DeleteCommentResp, error) {
	l := logic.NewDeleteCommentLogic(ctx, s.svcCtx)
	return l.DeleteComment(in)
}
