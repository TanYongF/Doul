// Code generated by goctl. DO NOT EDIT.
// Source: security.proto

package server

import (
	"context"

	"go_code/Doul/app/security/internal/logic"
	"go_code/Doul/app/security/internal/svc"
	"go_code/Doul/app/security/security"
)

type SecurityServer struct {
	svcCtx *svc.ServiceContext
	security.UnimplementedSecurityServer
}

func NewSecurityServer(svcCtx *svc.ServiceContext) *SecurityServer {
	return &SecurityServer{
		svcCtx: svcCtx,
	}
}

func (s *SecurityServer) Check(ctx context.Context, in *security.CheckLegaContentReq) (*security.CheckLegalContentResp, error) {
	l := logic.NewCheckLogic(ctx, s.svcCtx)
	return l.Check(in)
}
