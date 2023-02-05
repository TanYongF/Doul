package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_code/Doul/app/video/cmd/api/internal/logic"
	"go_code/Doul/app/video/cmd/api/internal/svc"
	"go_code/Doul/app/video/cmd/api/internal/types"
	"go_code/Doul/common/response"
)

func favoriteListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteVideofListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFavoriteListLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteList(&req)
		response.HttpResult(r, w, resp, err)
	}
}
