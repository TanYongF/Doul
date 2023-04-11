package model

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DyCommentModel = (*customDyCommentModel)(nil)

type (
	// DyCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDyCommentModel.
	DyCommentModel interface {
		dyCommentModel
		FindAllByVideoId(videoId int64) ([]*DyComment, error)
		DeleteById(commentId int64) error
		FindByVideoId(videoId int64, pageSize int64, pageNumber int64) ([]*DyComment, error)
	}

	customDyCommentModel struct {
		*defaultDyCommentModel
	}
)

// FindByVideoId find comments by videoId , not delete, pageable
// @pageSize:   the numbers of comments of a page
// @pageNumber: the offset of all result pages
func (c customDyCommentModel) FindByVideoId(videoId int64, pageSize int64, pageNumber int64) ([]*DyComment, error) {
	query := fmt.Sprintf("select * from %s where comment_id in (select t.id from (select comment_id as id from %s where video_id = ? and is_del = 0 LIMIT ?, ?) as t) order by created_at desc", c.tableName(), c.tableName())
	var comments []*DyComment
	err := c.QueryRowsNoCache(&comments, query, videoId, pageSize*(pageNumber-1), pageSize)
	switch err {
	case nil:
		return comments, nil
	default:
		logx.Errorf("GetComments Model err, err=%v", err)
		return nil, err
	}
}

func (c customDyCommentModel) DeleteById(commentId int64) error {
	query := fmt.Sprintf("update %s set is_del = 1 where comment_id = ?", c.tableName())
	_, err := c.ExecNoCache(query, commentId)
	switch err {
	case nil:
		return nil
	default:
		logx.Errorf("Delete Comments Model err, err=%v", err)
		return nil
	}

}

// FindAllByVideoId return all comments of video, not delete
func (c customDyCommentModel) FindAllByVideoId(videoId int64) ([]*DyComment, error) {
	query := fmt.Sprintf("select * from %s where video_id = ? and is_del = 0 order by created_at desc", c.tableName())
	var comments []*DyComment
	err := c.QueryRowsNoCache(&comments, query, videoId)
	switch err {
	case nil:
		return comments, nil
	default:
		logx.Errorf("GetComments Model err, err=%v", err)
		return nil, err
	}

}

// NewDyCommentModel returns a model for the database table.
func NewDyCommentModel(conn sqlx.SqlConn, c cache.CacheConf) DyCommentModel {
	return &customDyCommentModel{
		defaultDyCommentModel: newDyCommentModel(conn, c),
	}
}
