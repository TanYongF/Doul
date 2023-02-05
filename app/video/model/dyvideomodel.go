package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DyVideoModel = (*customDyVideoModel)(nil)

type (
	// DyVideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDyVideoModel.
	DyVideoModel interface {
		dyVideoModel
		GetVideoList(ctx context.Context) ([]*DyVideoWithUser, error)
		GetPublishListByUserId(ctx context.Context, userId int64) ([]*DyVideoWithUser, error)
		GetFavoriteListByUserId(ctx context.Context, userId int64) ([]*DyVideoWithUser, error)
	}
	//todo 该如何简化
	DyVideoWithUser struct {
		DyVideo
		Name          string `db:"name"`           // 用户名
		FollowerCount int64  `db:"follower_count"` // 粉丝总数
		IsFollow      int64  `db:"is_follow"`      // 是否已关注
		FollowCount   int64  `db:"follow_count"`   // 关注总数
	}
	customDyVideoModel struct {
		*defaultDyVideoModel
	}
)

// GetVideoList 搭配feed接口，获取视频集合
func (c customDyVideoModel) GetVideoList(ctx context.Context) ([]*DyVideoWithUser, error) {
	//TODO 搭配推荐算法
	query := "select dv.* ,  du.name , du.follower_count , du.follow_count, du.is_follow  from dy_video dv join dy_user du using (user_id) order by create_date desc limit 30"

	var resp []*DyVideoWithUser
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query)

	switch err {
	case nil:
		return resp, nil
	default:
		logx.Errorf("GetVideoList Model db err, err=%v", err)
		return nil, err
	}
}

// GetPublishListByUserId 获取特定用户的投稿视频集合
func (c customDyVideoModel) GetPublishListByUserId(ctx context.Context, userId int64) ([]*DyVideoWithUser, error) {
	query := fmt.Sprintf("select dv.* ,  du.name , du.follower_count , du.follow_count, du.is_follow from dy_video dv  left join dy_user du on dv.user_id = du.user_id where dv.user_id = ?")
	var resp []*DyVideoWithUser
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	default:
		logx.Errorf("GetPulishListByUserID Model err, err=%v", err)
		return nil, err
	}
}

func (c customDyVideoModel) GetFavoriteListByUserId(ctx context.Context, userId int64) ([]*DyVideoWithUser, error) {
	query := fmt.Sprintf("select dv.* ,  du.name , du.follower_count , du.follow_count, du.is_follow from dy_video dv  left join dy_user du on dv.user_id = du.user_id where (select ifnull((select 1 from dy_favorite df where df.user_id = ? and df.video_id = dv.video_id  and df.is_del = 0), 0)) = 1")
	var resp []*DyVideoWithUser
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	default:
		logx.Errorf("GetFavoriteListByUserId Model err, err=%v", err)
		return nil, err
	}
}

// NewDyVideoModel returns a model for the database table.
func NewDyVideoModel(conn sqlx.SqlConn, c cache.CacheConf) DyVideoModel {
	return &customDyVideoModel{
		defaultDyVideoModel: newDyVideoModel(conn, c),
	}
}
