package logic

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"go_code/Doul/app/usercenter/cmd/rpc/user"
	"go_code/Doul/common/globalkey"

	"github.com/zeromicro/go-zero/core/stringx"
	"go_code/Doul/app/usercenter/cmd/rpc/internal/svc"

	"go_code/Doul/app/usercenter/model"

	"go_code/Doul/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAuthLogic {
	return &CheckAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CheckAuth check the user has authed
func (l *CheckAuthLogic) CheckAuth(in *user.CheckAuthReq) (*user.CheckAuthReply, error) {
	userJSON, err := l.svcCtx.RedisClient.Get(globalkey.TokenPrefix + in.GetToken())
	var dyUser model.DyUser
	isAuthed := false
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CACHE_ERROR), "failed to access redis when get user's information")
	}
	if stringx.NotEmpty(userJSON) {
		err := json.Unmarshal([]byte(userJSON), &dyUser)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.CACHE_ERROR), "failed to unmarshal json when get user's information")
		}
		//update the token's expire time (24 hours).
		l.svcCtx.RedisClient.Expire(globalkey.TokenPrefix+in.GetToken(), int(globalkey.TokenExpireTime.Seconds()))
		isAuthed = true
	}

	return &user.CheckAuthReply{
		Authed:   isAuthed,
		AuthedId: dyUser.UserId,
	}, nil
}
