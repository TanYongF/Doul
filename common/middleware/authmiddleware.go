package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stringx"
	"go_code/Doul/app/usercenter/cmd/rpc/userclient"
	"go_code/Doul/common/response"
	"go_code/Doul/common/xerr"
	"net/http"
)

type AuthMiddleware struct {
	UserRpc userclient.User
}

func NewAuthMiddleware(userClient userclient.User) *AuthMiddleware {
	return &AuthMiddleware{
		UserRpc: userClient,
	}
}

//
// Handle
//  @Description: 鉴权中间件方法
//  @receiver m 鉴权中间件
//  @param next
//  @return http.HandlerFunc
//
func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Get token from query or form-data
		token := r.FormValue("token")
		if stringx.HasEmpty(token) {
			token = r.PostFormValue("token")
		}

		//check is the token empty
		if stringx.HasEmpty(token) {
			response.HttpResult(r, w, false, xerr.NewErrCode(xerr.NO_AUTH))
			return
		}

		//invoke the rpc service to check users' auth
		auth, err := m.UserRpc.CheckAuth(context.Background(), &userclient.CheckAuthReq{
			Token: token,
		})

		//if has err or no authed, return message
		if err != nil {
			response.HttpResult(r, w, nil, err)
			return
		} else if !auth.Authed {
			response.HttpResult(r, w, false, xerr.NewErrCode(xerr.NO_AUTH))
			logx.Infof("user %s has not authed", r.FormValue("token"))
			return
		}
		next(w, r.WithContext(context.WithValue(r.Context(), "auth_id", auth.AuthedId)))
	}
}
