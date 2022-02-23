package moment_service

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/dao"
	"github.com/louis296/sunny-weather/dao/model"
	"github.com/louis296/sunny-weather/pkg/common"
)

type CreateMomentReq struct {
	Context string
}

type CreateMomentResp struct {
	Id      int
	Like    int
	Context string
	User    model.UserResp
}

func (r *CreateMomentReq) Handler(c *gin.Context) (interface{}, error) {
	tx := dao.DB.Begin()
	currentUser, _ := common.GetCurrentUser(c)
	moment := &model.Moment{
		Context: r.Context,
		UserId:  currentUser.Id,
	}
	if err := dao.CreateMoment(moment, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	moment = dao.GetMomentById(moment.Id)
	author := dao.GetUserById(moment.UserId)
	return CreateMomentResp{
		Id:      moment.Id,
		Like:    moment.Like,
		Context: moment.Context,
		User:    author.GenResp(),
	}, nil
}
