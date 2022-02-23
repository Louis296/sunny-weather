package comment_service

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/dao"
	"github.com/louis296/sunny-weather/dao/model"
	"github.com/louis296/sunny-weather/pkg/common"
)

type CreateCommentReq struct {
	Context  string
	MomentId int
}

type CreateCommentResp struct {
	model.CommentResp
}

func (r *CreateCommentReq) Handler(c *gin.Context) (interface{}, error) {
	tx := dao.DB.Begin()
	currentUser, _ := common.GetCurrentUser(c)
	comment := &model.Comment{
		Context:  r.Context,
		UserId:   currentUser.Id,
		MomentId: r.MomentId,
	}
	if err := dao.CreateComment(comment, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	comment = dao.GetCommentById(comment.Id)
	commentResp := comment.GenResp()
	commentResp.User = currentUser.GenResp()
	return CreateCommentResp{commentResp}, nil
}
