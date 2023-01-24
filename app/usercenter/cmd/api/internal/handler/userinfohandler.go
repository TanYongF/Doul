package handler

import (
	"go_code/Doul/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_code/Doul/app/usercenter/cmd/api/internal/logic"
	"go_code/Doul/app/usercenter/cmd/api/internal/svc"
	"go_code/Doul/app/usercenter/cmd/api/internal/types"
)

func userinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserinfoLogic(r.Context(), svcCtx)

		resp, err := l.Userinfo(&req)
		response.HttpResult(r, w, resp, err)
	}
}
