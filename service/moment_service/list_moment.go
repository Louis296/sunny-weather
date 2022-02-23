package moment_service

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/dao"
	"github.com/louis296/sunny-weather/dao/model"
)

type ListMomentReq struct {
	Offset int
	Limit  int
}

type ListMomentResp struct {
	List []model.MomentResp
}

func (r *ListMomentReq) Handler(c *gin.Context) (interface{}, error) {
	result := dao.GetAllMoments(r.Limit, r.Offset)
	var userIds []int
	var list []model.MomentResp
	for _, item := range result {
		userIds = append(userIds, item.UserId)
	}
	userLookup, err := dao.GetUserIdLookup(userIds...)
	if err != nil {
		return nil, err
	}
	for _, item := range result {
		momentResp := item.GenResp()
		if user, ok := userLookup[item.UserId]; ok {
			momentResp.User = user.GenResp()
		}
		list = append(list, momentResp)
	}
	return ListMomentResp{List: list}, nil
}
