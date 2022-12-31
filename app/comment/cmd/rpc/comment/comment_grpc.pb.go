// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: comment.proto

package comment

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CommentClient is the client API for Comment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentClient interface {
	Greet(ctx context.Context, in *CommentReq, opts ...grpc.CallOption) (*CommentResp, error)
}

type commentClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentClient(cc grpc.ClientConnInterface) CommentClient {
	return &commentClient{cc}
}

func (c *commentClient) Greet(ctx context.Context, in *CommentReq, opts ...grpc.CallOption) (*CommentResp, error) {
	out := new(CommentResp)
	err := c.cc.Invoke(ctx, "/comment.Comment/greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServer is the server API for Comment service.
// All implementations must embed UnimplementedCommentServer
// for forward compatibility
type CommentServer interface {
	Greet(context.Context, *CommentReq) (*CommentResp, error)
	mustEmbedUnimplementedCommentServer()
}

// UnimplementedCommentServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServer struct {
}

func (UnimplementedCommentServer) Greet(context.Context, *CommentReq) (*CommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedCommentServer) mustEmbedUnimplementedCommentServer() {}

// UnsafeCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServer will
// result in compilation errors.
type UnsafeCommentServer interface {
	mustEmbedUnimplementedCommentServer()
}

func RegisterCommentServer(s grpc.ServiceRegistrar, srv CommentServer) {
	s.RegisterService(&Comment_ServiceDesc, srv)
}

func _Comment_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.Comment/greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).Greet(ctx, req.(*CommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Comment_ServiceDesc is the grpc.ServiceDesc for Comment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.Comment",
	HandlerType: (*CommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "greet",
			Handler:    _Comment_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
