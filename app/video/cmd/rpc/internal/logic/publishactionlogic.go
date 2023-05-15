package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"go_code/Doul/app/video/model"
	"go_code/Doul/common/xerr"
	"time"

	"go_code/Doul/app/video/cmd/rpc/internal/svc"
	"go_code/Doul/app/video/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishActionLogic) PublishAction(in *pb.PublishReq) (*pb.PublishResp, error) {
	_, err := l.svcCtx.DyVideoModel.Insert(l.ctx, &model.DyVideo{
		UserId: sql.NullInt64{
			Int64: in.UserId,
			Valid: true,
		},
		PlayUrl:       in.VideoUrl,
		CoverUrl:      in.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title: sql.NullString{
			String: in.Title,
			Valid:  true,
		},
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "database error when insert video record")
	}
	return &pb.PublishResp{}, nil
}
