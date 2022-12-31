package logic

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"go_code/Doul/app/usercenter/cmd/rpc/internal/svc"
	"go_code/Doul/app/usercenter/cmd/rpc/user"
	"go_code/Doul/common/globalkey"
	"go_code/Doul/common/tool"
	"go_code/Doul/common/xerr"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginReply, error) {
	//查找用户是否存在
	dyuser, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.GetUsername())
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindOneBySn db err : %v , in:%+v", err, in)
	}
	if dyuser == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.NO_SUCH_USER), "such user not exsist")
	}
	//验证密码
	md5Pass := tool.Md5(in.GetPassword())
	if strings.Compare(dyuser.Password, md5Pass) != 0 {
		l.Logger.Error("验证错误")
		return nil, errors.Wrap(xerr.NewErrCode(xerr.WRONG_PASSWORD), "password wrong in")
	}

	//生成Token
	token := in.Username + in.Password
	//将token加入redis中，过期时间是24小时, 键是token, 值是用户对象
	userJson, _ := json.Marshal(*dyuser)
	expireDuration := time.Hour * 24
	err = l.svcCtx.RedisClient.SetexCtx(l.ctx, globalkey.TokenPrefix+token, string(userJson), int(expireDuration))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CACHE_ERROR), "cache wrong")
	}

	return &user.LoginReply{
		UserId: dyuser.UserId,
		Token:  token,
	}, nil
}
