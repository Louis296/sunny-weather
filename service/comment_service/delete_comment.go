package comment_service

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/dao"
)

type DeleteCommentReq struct {
	CommentId int
}

func (r *DeleteCommentReq) Handler(c *gin.Context) (interface{}, error) {
	tx := dao.DB.Begin()
	defer func() {
		if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err := dao.DeleteCommentById(r.CommentId, tx); err != nil {
		return nil, err
	}
	return nil, nil
}
