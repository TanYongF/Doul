// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package commentclient

import (
	"context"

	"go_code/Doul/app/comment/cmd/rpc/comment"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommentBody     = comment.CommentBody
	CommentListReq  = comment.CommentListReq
	CommentListResp = comment.CommentListResp
	PutCommentReq   = comment.PutCommentReq
	PutCommentResp  = comment.PutCommentResp
	User            = comment.User

	Comment interface {
		GetCommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error)
		PutComment(ctx context.Context, in *PutCommentReq, opts ...grpc.CallOption) (*PutCommentResp, error)
	}

	defaultComment struct {
		cli zrpc.Client
	}
)

func NewComment(cli zrpc.Client) Comment {
	return &defaultComment{
		cli: cli,
	}
}

func (m *defaultComment) GetCommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error) {
	client := comment.NewCommentClient(m.cli.Conn())
	return client.GetCommentList(ctx, in, opts...)
}

func (m *defaultComment) PutComment(ctx context.Context, in *PutCommentReq, opts ...grpc.CallOption) (*PutCommentResp, error) {
	client := comment.NewCommentClient(m.cli.Conn())
	return client.PutComment(ctx, in, opts...)
}
