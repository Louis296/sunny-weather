package comment_service

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/dao"
	"github.com/louis296/sunny-weather/dao/model"
)

type ListCommentReq struct {
	MomentId int
	Offset   int
	Limit    int
}

type ListCommentResp struct {
	Total int
	List  []model.CommentResp
}

func (r *ListCommentReq) Handler(c *gin.Context) (interface{}, error) {
	result, total := dao.GetCommentsByMomentId(r.MomentId, r.Limit, r.Offset)
	var userIds []int
	var list []model.CommentResp
	for _, item := range result {
		userIds = append(userIds, item.UserId)
	}
	userLookup, err := dao.GetUserIdLookup(userIds...)
	if err != nil {
		return nil, err
	}
	for _, item := range result {
		commentResp := item.GenResp()
		if user, ok := userLookup[item.UserId]; ok {
			commentResp.User = user.GenResp()
		}
		list = append(list, commentResp)
	}
	return ListCommentResp{List: list, Total: total}, nil
}
