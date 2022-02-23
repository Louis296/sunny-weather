package moment_service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/dao"
	"github.com/louis296/sunny-weather/dao/model"
	"github.com/louis296/sunny-weather/pkg/common"
	"strconv"
)

type UpdateMomentReq struct {
	Context  string
	MomentId int
}

type UpdateMomentResp struct {
	model.MomentResp
}

func (r *UpdateMomentReq) Handler(c *gin.Context) (interface{}, error) {
	tx := dao.DB.Begin()
	defer func() {
		if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	currentUser, _ := common.GetCurrentUser(c)
	moment := dao.GetMomentById(r.MomentId)
	if moment == nil {
		return nil, errors.New("no moment with id " + strconv.Itoa(r.MomentId))
	}
	if moment.UserId != currentUser.Id {
		return nil, errors.New("can't edit moment created by others")
	}
	moment.Context = r.Context
	if err := dao.SaveMoment(moment, tx); err != nil {
		return nil, err
	}
	momentResp := moment.GenResp()
	author := dao.GetUserById(moment.UserId)
	momentResp.User = author.GenResp()
	return UpdateMomentResp{momentResp}, nil
}
