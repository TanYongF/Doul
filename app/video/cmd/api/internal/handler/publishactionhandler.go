package handler

import (
	"go_code/Doul/app/video/cmd/api/internal/logic"
	"go_code/Doul/app/video/cmd/api/internal/svc"
	"go_code/Doul/common/response"
	"net/http"
)

func publishActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPublishActionLogic(r.Context(), svcCtx)
		resp, err := l.PublishAction(r)
		response.HttpResult(r, w, resp, err)
	}
}
