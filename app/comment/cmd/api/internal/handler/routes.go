// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go_code/Doul/app/comment/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/douyin/comment/list",
					Handler: CommentListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/douyin/comment/action",
					Handler: CommentActionHandler(serverCtx),
				},
			}...,
		),
	)
}
