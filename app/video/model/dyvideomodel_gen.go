// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	dyVideoFieldNames          = builder.RawFieldNames(&DyVideo{})
	dyVideoRows                = strings.Join(dyVideoFieldNames, ",")
	dyVideoRowsExpectAutoSet   = strings.Join(stringx.Remove(dyVideoFieldNames, "`video_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dyVideoRowsWithPlaceHolder = strings.Join(stringx.Remove(dyVideoFieldNames, "`video_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheDouyinDyVideoVideoIdPrefix = "cache:douyin:dyVideo:videoId:"
)

type (
	dyVideoModel interface {
		Insert(ctx context.Context, data *DyVideo) (sql.Result, error)
		FindOne(ctx context.Context, videoId int64) (*DyVideo, error)
		Update(ctx context.Context, data *DyVideo) error
		Delete(ctx context.Context, videoId int64) error
	}

	defaultDyVideoModel struct {
		sqlc.CachedConn
		table string
	}

	DyVideo struct {
		VideoId       int64          `db:"video_id"`       // 视频ID
		UserId        sql.NullInt64  `db:"user_id"`        // 用户ID
		PlayUrl       string         `db:"play_url"`       // 播放地址
		CoverUrl      string         `db:"cover_url"`      // 视频封面地址
		FavoriteCount int64          `db:"favorite_count"` // 点赞量
		CommentCount  int64          `db:"comment_count"`  // 评论量
		Title         sql.NullString `db:"title"`          // 视频标题
		CreateDate    time.Time      `db:"create_date"`    // 创建时间
		UpdateDate    time.Time      `db:"update_date"`    // 更新时间
	}
)

func newDyVideoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultDyVideoModel {
	return &defaultDyVideoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`dy_video`",
	}
}

func (m *defaultDyVideoModel) Delete(ctx context.Context, videoId int64) error {
	douyinDyVideoVideoIdKey := fmt.Sprintf("%s%v", cacheDouyinDyVideoVideoIdPrefix, videoId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `video_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, videoId)
	}, douyinDyVideoVideoIdKey)
	return err
}

func (m *defaultDyVideoModel) FindOne(ctx context.Context, videoId int64) (*DyVideo, error) {
	douyinDyVideoVideoIdKey := fmt.Sprintf("%s%v", cacheDouyinDyVideoVideoIdPrefix, videoId)
	var resp DyVideo
	err := m.QueryRowCtx(ctx, &resp, douyinDyVideoVideoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `video_id` = ? limit 1", dyVideoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, videoId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDyVideoModel) Insert(ctx context.Context, data *DyVideo) (sql.Result, error) {
	douyinDyVideoVideoIdKey := fmt.Sprintf("%s%v", cacheDouyinDyVideoVideoIdPrefix, data.VideoId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, dyVideoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.PlayUrl, data.CoverUrl, data.FavoriteCount, data.CommentCount, data.Title, data.CreateDate, data.UpdateDate)
	}, douyinDyVideoVideoIdKey)
	return ret, err
}

func (m *defaultDyVideoModel) Update(ctx context.Context, data *DyVideo) error {
	douyinDyVideoVideoIdKey := fmt.Sprintf("%s%v", cacheDouyinDyVideoVideoIdPrefix, data.VideoId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `video_id` = ?", m.table, dyVideoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.PlayUrl, data.CoverUrl, data.FavoriteCount, data.CommentCount, data.Title, data.CreateDate, data.UpdateDate, data.VideoId)
	}, douyinDyVideoVideoIdKey)
	return err
}

func (m *defaultDyVideoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheDouyinDyVideoVideoIdPrefix, primary)
}

func (m *defaultDyVideoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `video_id` = ? limit 1", dyVideoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultDyVideoModel) tableName() string {
	return m.table
}