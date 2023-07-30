package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_code/Doul/app/usercenter/cmd/rpc/internal/svc"
	"go_code/Doul/app/usercenter/cmd/rpc/user"
	"go_code/Doul/app/usercenter/model"
	"go_code/Doul/common/globalkey"
	"go_code/Doul/common/tool"
	"go_code/Doul/common/xerr"
	"strconv"
	"strings"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")
var ErrUserNoExistsError = xerr.NewErrMsg("无此用户")

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginReply, error) {
	dyuser, err := l.checkUser(in.Username, in.Password)
	if err != nil {
		return nil, err
	}

	//Generate token and user detail json
	token := tool.TokenGenerator()

	//将token加入redis中，过期时间是24小时, 键是token, 值是用户对象
	val := strconv.FormatInt(dyuser.UserId, 10) + ":" + dyuser.Name
	err = l.svcCtx.RedisClient.SetexCtx(l.ctx, globalkey.GetTokenKeyByToken(token), val, int(globalkey.TokenExpireTime.Seconds()))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CACHE_ERROR), "cache wrong when insert token")
	}

	return &user.LoginReply{
		UserId: dyuser.UserId,
		Token:  token,
	}, nil
}

func (l *LoginLogic) checkUser(username string, password string) (dbUser *model.DyUser, err error) {
	dbUser, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, username)
	//if has error

	if err != nil {
		switch err {
		case sqlx.ErrNotFound:
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.NO_SUCH_USER), "无此用户")
		default:
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据用户名称查询用户信息失败，mobile:%s,err:%v", username, err)
		}
	}
	//验证密码
	formPass := tool.Md5(password)
	if strings.Compare(dbUser.Password, formPass) != 0 {
		return nil, errors.Wrap(ErrUsernamePwdError, "密码匹配出错")
	}

	return dbUser, nil
}
