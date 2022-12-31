package middleware

import (
	"context"
	"fmt"
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

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware before")
		auth, err := m.UserRpc.CheckAuth(context.Background(), &userclient.CheckAuthReq{
			Token: r.FormValue("token"),
		})
		if err != nil {
			response.Response(w, false, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR))
			return
		} else if !auth.Authed {
			response.Response(w, false, xerr.NewErrCode(xerr.NO_AUTH))
			return
		}
		next(w, r)
		fmt.Println("middleware after")
	}
}
