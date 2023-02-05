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
		FindAllByVideoId(videoId int64) ([]DyComment, error)
		DeleteById(commentId int64) error
	}

	customDyCommentModel struct {
		*defaultDyCommentModel
	}
)

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

func (c customDyCommentModel) FindAllByVideoId(videoId int64) ([]DyComment, error) {
	query := fmt.Sprintf("select * from %s where video_id = ? and is_del = 0", c.tableName())
	var comments []DyComment
	err := c.QueryRowsNoCache(comments, query, videoId)
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
