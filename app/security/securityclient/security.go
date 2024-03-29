// Code generated by goctl. DO NOT EDIT.
// Source: security.proto

package securityclient

import (
	"context"

	"go_code/Doul/app/security/security"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CheckLegaContentReq   = security.CheckLegaContentReq
	CheckLegalContentResp = security.CheckLegalContentResp

	Security interface {
		Check(ctx context.Context, in *CheckLegaContentReq, opts ...grpc.CallOption) (*CheckLegalContentResp, error)
	}

	defaultSecurity struct {
		cli zrpc.Client
	}
)

func NewSecurity(cli zrpc.Client) Security {
	return &defaultSecurity{
		cli: cli,
	}
}

func (m *defaultSecurity) Check(ctx context.Context, in *CheckLegaContentReq, opts ...grpc.CallOption) (*CheckLegalContentResp, error) {
	client := security.NewSecurityClient(m.cli.Conn())
	return client.Check(ctx, in, opts...)
}
