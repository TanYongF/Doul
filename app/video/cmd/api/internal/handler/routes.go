// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go_code/Doul/app/video/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/douyin/feed",
				Handler: feedHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/douyin/publish/action",
					Handler: publishActionHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/douyin/publish/list",
					Handler: publishListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/douyin/favorite/list",
					Handler: favoriteListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/douyin/favorite/action",
					Handler: favoriteHandler(serverCtx),
				},
			}...,
		),
	)
}
