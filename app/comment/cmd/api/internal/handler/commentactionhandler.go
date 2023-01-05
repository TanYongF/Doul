package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_code/Doul/app/comment/cmd/api/internal/logic"
	"go_code/Doul/app/comment/cmd/api/internal/svc"
	"go_code/Doul/app/comment/cmd/api/internal/types"
	"go_code/Doul/common/response"
)

func CommentActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentActionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCommentActionLogic(r.Context(), svcCtx)
		resp, err := l.CommentAction(&req)
		response.HttpResult(r, w, resp, err)
	}
}
