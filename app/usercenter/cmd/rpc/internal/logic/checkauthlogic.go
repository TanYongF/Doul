package logic

import (
	"context"
	"github.com/pkg/errors"
	"go_code/Doul/app/usercenter/cmd/rpc/user"
	"go_code/Doul/common/globalkey"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/stringx"
	"go_code/Doul/app/usercenter/cmd/rpc/internal/svc"

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
	value, err := l.svcCtx.RedisClient.Get(globalkey.GetTokenKeyByToken(in.Token))
	isAuthed := false
	var userId int64

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CACHE_ERROR), "failed to access redis when get user's information")
	}
	if stringx.NotEmpty(value) {
		//update the token's expire time (24 hours).
		err := l.svcCtx.RedisClient.Expire(globalkey.GetTokenKeyByToken(in.Token), int(globalkey.TokenExpireTime.Seconds()))
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.CACHE_ERROR), "Error while updating token validity")
		}
		isAuthed = true
		userId, err = GetUserIdFromTokenKey(value)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "take info error when explain info")
		}
	} else {
		isAuthed = false
	}

	return &user.CheckAuthReply{
		Authed:   isAuthed,
		AuthedId: userId,
	}, nil
}

// GetUserIdFromTokenKey 通过value获取 token
func GetUserIdFromTokenKey(info string) (userId int64, err error) {
	arr := strings.Split(info, ":")
	userId, err = strconv.ParseInt(arr[0], 10, 64)
	if err != nil {
		return 0, err
	}
	return userId, err
}
