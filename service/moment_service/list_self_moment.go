package moment_service

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/dao"
	"github.com/louis296/sunny-weather/dao/model"
	"github.com/louis296/sunny-weather/pkg/common"
)

type ListSelfMomentReq struct {
	Limit  int
	Offset int
}

type ListSelfMomentResp struct {
	Total int
	List  []model.MomentResp
}

func (r *ListSelfMomentReq) Handler(c *gin.Context) (interface{}, error) {
	currentUser, _ := common.GetCurrentUser(c)
	result, total := dao.GetMomentsByUserId(currentUser.Id, r.Limit, r.Offset)
	var list []model.MomentResp
	for _, item := range result {
		momentResp := item.GenResp()
		momentResp.User = currentUser.GenResp()
		list = append(list, momentResp)
	}
	return ListSelfMomentResp{List: list, Total: total}, nil
}
