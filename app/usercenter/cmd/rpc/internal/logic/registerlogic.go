package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"go_code/Doul/app/usercenter/cmd/rpc/internal/svc"
	"go_code/Doul/app/usercenter/cmd/rpc/user"
	"go_code/Doul/app/usercenter/model"
	"go_code/Doul/common/tool"
	"go_code/Doul/common/xerr"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterReply, error) {
	passAfterMd5 := tool.Md5(in.GetPassword())
	dyuser, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.GetUsername())
	if err != nil && !errors.Is(err, sqlc.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "can't found")
	}
	//如果username已被使用
	if dyuser != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USERNAME_HAS_REGISTER), "用户名已经被注册")
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.DyUser{
		Name:          in.Username,
		FollowerCount: 0,
		Password:      passAfterMd5,
		Salt:          "fa00900fafa",
		FollowCount:   0,
	})
	dyuser, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.GetUsername())
	if err != nil {
		return nil, err
	}
	return &user.RegisterReply{
		UserId: dyuser.UserId,
		Token:  "fafafaf8989",
	}, nil
}
