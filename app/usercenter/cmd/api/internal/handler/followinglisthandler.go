package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_code/Doul/app/usercenter/cmd/api/internal/logic"
	"go_code/Doul/app/usercenter/cmd/api/internal/svc"
	"go_code/Doul/app/usercenter/cmd/api/internal/types"
	"go_code/Doul/common/response"
)

func followingListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowingListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFollowingListLogic(r.Context(), svcCtx)
		resp, err := l.FollowingList(&req)
		response.HttpResult(r, w, resp, err)
	}
}
